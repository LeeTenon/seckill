package svc

import (
	"Tstore/backend/user/api/internal/config"
	"Tstore/backend/user/rpc/login/loginclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Login  loginclient.Login
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Login:  loginclient.NewLogin(zrpc.MustNewClient(c.Login)),
	}
}
