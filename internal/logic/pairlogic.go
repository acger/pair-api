package logic

import (
	"context"
	"encoding/json"
	"github.com/acger/pair-api/internal/svc"
	"github.com/acger/pair-api/internal/types"
	"github.com/acger/pair-svc/pair"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type PairLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPairLogic(ctx context.Context, svcCtx *svc.ServiceContext) PairLogic {
	return PairLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PairLogic) Pair() (*types.EleListRsp, error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	r, err := l.svcCtx.PairSvc.ElementPair(l.ctx, &pair.ElePairReq{Uid: uint64(uid)})

	if err != nil {
		return &types.EleListRsp{Code: 1}, nil
	}

	if r.Code != 0 {
		return &types.EleListRsp{Code: r.Code, Message: r.Message}, nil
	}

	userEleRsp := make([]*types.UserElement, len(r.UserElement))

	for i,ue := range r.UserElement{
		userEleRsp[i] = &types.UserElement{}
		copier.Copy(userEleRsp[i], ue)
	}

	return &types.EleListRsp{Code: 0, UserElement: userEleRsp}, nil
}






















