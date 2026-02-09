package logic

import (
	"context"

	"fengfengstudy/user/rpc/internal/svc"
	"fengfengstudy/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {

	return &user.UserResponse{
		Id:     in.Id,
		Name:   "湖南省，长沙市",
		Gender: true,
	}, nil
}
