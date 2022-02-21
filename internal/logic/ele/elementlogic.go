package ele

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"

	"github.com/acger/pair-api/internal/svc"
	"github.com/acger/pair-api/internal/types"
	"github.com/acger/pair-svc/pair"
	"github.com/zeromicro/go-zero/core/logx"
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

func (l *ElementLogic) Element() (resp *types.EleRsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	r, err := l.svcCtx.PairSvc.ElementView(l.ctx, &pair.EleViewReq{Uid: uint64(uid)})

	if err != nil {
		return &types.EleRsp{Code: 1}, nil
	}

	if r.Code != 0 {
		return &types.EleRsp{Code: r.Code, Message: r.Message}, nil
	}


	ele := &types.Element{}

	copier.Copy(ele, r.Element)

	return &types.EleRsp{Code: 0, Element: ele}, nil
}
