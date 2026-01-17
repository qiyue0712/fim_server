// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"fim_server/core"
	"fim_server/fim_auth/auth_api/internal/config"
	"fim_server/fim_user/user_rpc/users"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc users.Users // 添加 UserRpc 字段
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	client := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	//mysqlDb.AutoMigrate(&auth_models.UserModel{})
	return &ServiceContext{
		Config:  c,
		DB:      mysqlDb,
		Redis:   client,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
	}
}
