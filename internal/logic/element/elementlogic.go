package logic

import (
	"context"
	"encoding/json"
	"github.com/acger/pair-api/internal/svc"
	"github.com/acger/pair-api/internal/types"
	"github.com/acger/pair-svc/pairclient"
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"
)

type ElementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) ElementLogic {
	return ElementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ElementLogic) Element() (*types.EleRsp, error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	r, err := l.svcCtx.PairSvc.ElementView(l.ctx, &pairclient.EleViewReq{Uid: uint64(uid)})

	if err != nil {
		return &types.EleRsp{Code: 1}, nil
	}

	if r.Code != 0 {
		return &types.EleRsp{Code: r.Code, Message: r.Message}, nil
	}

	ele := make([]*types.Element, len(r.Element))

	for i, e := range r.Element {
		et := types.Element{}
		copier.Copy(&et, e)
		ele[i] = &et
	}

	return &types.EleRsp{Code: 0, Element: ele}, nil
}
