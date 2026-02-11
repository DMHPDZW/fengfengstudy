// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fengfengstudy/apistudy/user/api_v3/internal/logic"
	"fengfengstudy/apistudy/user/api_v3/internal/svc"
	"fengfengstudy/apistudy/user/api_v3/internal/types"
	"fengfengstudy/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		// if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r, w, resp, err)

	}
}
