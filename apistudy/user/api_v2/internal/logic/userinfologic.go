// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"fengfengstudy/apistudy/user/api_v2/internal/svc"
	"fengfengstudy/apistudy/user/api_v2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo() (resp *types.UserInfoResponse, err error) {

	return &types.UserInfoResponse{
		UserId:   1,
		Username: "fengfeng",
	}, nil
}
