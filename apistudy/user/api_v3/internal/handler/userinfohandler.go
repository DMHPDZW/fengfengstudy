// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fengfengstudy/apistudy/user/api_v3/internal/logic"
	"fengfengstudy/apistudy/user/api_v3/internal/svc"
	"fengfengstudy/common/response"

	"net/http"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo()
		// if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r, w, resp, err)

	}
}
