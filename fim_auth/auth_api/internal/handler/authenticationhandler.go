// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fim_server/common/response"
	"net/http"

	"fim_server/fim_auth/auth_api/internal/logic"
	"fim_server/fim_auth/auth_api/internal/svc"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r, w, resp, err)
	}
}
