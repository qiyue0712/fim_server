// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"
	"fim_server/utils/jwts"
	"fmt"

	"fim_server/fim_auth/auth_api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(token string) (resp string, err error) {
	if token == "" {
		l.Error("认证失败：token为空")
		err = errors.New("认证失败")
		return
	}

	l.Infof("收到token: %s", token)
	l.Infof("使用AccessSecret: %s", l.svcCtx.Config.Auth.AccessSecret)

	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		l.Errorf("token解析失败: %v", err)
		err = errors.New("认证失败")
		return
	}

	l.Infof("token解析成功，用户ID: %d", payload.UserID)

	logoutKey := fmt.Sprintf("logout_%d", payload.UserID)
	l.Infof("检查Redis key: %s", logoutKey)
	_, err = l.svcCtx.Redis.Get(logoutKey).Result()
	if err == nil {
		l.Errorf("用户已注销，Redis中存在key: %s", logoutKey)
		err = errors.New("认证失败")
		return
	}

	l.Info("认证成功")
	resp = "ok"
	err = nil
	return
}
