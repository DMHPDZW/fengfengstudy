// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"fengfengstudy/demo/user/rpc/userclient"
	"fengfengstudy/demo/video/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
