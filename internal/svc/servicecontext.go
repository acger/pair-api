package svc

import (
	"github.com/acger/pair-api/internal/config"
	"github.com/acger/pair-svc/pair"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	PairSvc pair.Pair
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		PairSvc: pair.NewPair(zrpc.MustNewClient(c.PairSvc)),
	}
}
