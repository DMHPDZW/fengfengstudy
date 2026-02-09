// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fengfengstudy/demo/video/api/internal/logic"
	"fengfengstudy/demo/video/api/internal/svc"
	"fengfengstudy/demo/video/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetVideoLogic(r.Context(), svcCtx)
		resp, err := l.GetVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
