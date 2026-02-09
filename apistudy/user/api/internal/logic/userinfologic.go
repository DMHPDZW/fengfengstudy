// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"fengfengstudy/apistudy/user/api/internal/svc"
	"fengfengstudy/apistudy/user/api/internal/types"

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
	// todo: add your logic here and delete this line

	return &types.UserInfoResponse{Code: 1, Data: types.UserInfo{UserId: 1, Username: "fengfeng"}, Msg: `ok`}, nil
}
