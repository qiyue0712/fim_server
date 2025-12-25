// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fim_server/common/response"
	"fim_server/fim_auth/auth_api/internal/logic"
	"fim_server/fim_auth/auth_api/internal/svc"
	"net/http"
)

func open_login_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewOpen_login_infoLogic(r.Context(), svcCtx)
		resp, err := l.Open_login_info()
		response.Response(r, w, resp, err)
	}
}
