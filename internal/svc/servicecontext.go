package svc

import (
	"github.com/acger/pair-api/internal/config"
	"github.com/acger/pair-svc/pairclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	PairSvc pairclient.Pair
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		PairSvc: pairclient.NewPair(zrpc.MustNewClient(c.PairSvc)),
	}
}
