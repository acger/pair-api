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

type SaveElementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) SaveElementLogic {
	return SaveElementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveElementLogic) SaveElement(req types.EleSaveReq) (*types.Rsp, error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	ele := make([]*pairclient.Element, len(req.Element))

	for i, e := range req.Element {

		if e.Name == ""{
			continue
		}

		et := pairclient.Element{}
		copier.Copy(&et, e)
		ele[i] = &et
	}

	r, err := l.svcCtx.PairSvc.ElementSave(l.ctx, &pairclient.EleSaveReq{
		Uid:     uint64(uid),
		Element: ele,
	})

	if err != nil {
		return &types.Rsp{Code: 1}, nil
	}

	if r.Code != 0 {
		return &types.Rsp{Code: r.Code, Message: r.Message}, nil
	}

	return &types.Rsp{Code: 0}, nil
}
