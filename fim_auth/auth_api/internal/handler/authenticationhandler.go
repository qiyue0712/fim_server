// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fim_server/common/response"
	"fim_server/fim_auth/auth_api/internal/logic"
	"fim_server/fim_auth/auth_api/internal/svc"
	"net/http"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		resp, err := l.Authentication(token)
		response.Response(r, w, resp, err)
	}
}
