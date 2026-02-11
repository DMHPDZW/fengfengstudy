// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fengfengstudy/common/response"
	"net/http"

	"fengfengstudy/apistudy/user/api_v2/internal/logic"
	"fengfengstudy/apistudy/user/api_v2/internal/svc"
	"fengfengstudy/apistudy/user/api_v2/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
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
		response.Response(r, w, resp, err)
	}
}
