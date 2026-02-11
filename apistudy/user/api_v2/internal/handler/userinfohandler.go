// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fengfengstudy/common/response"
	"net/http"

	"fengfengstudy/apistudy/user/api_v2/internal/logic"
	"fengfengstudy/apistudy/user/api_v2/internal/svc"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo()
		response.Response(r, w, resp, err)

	}
}
