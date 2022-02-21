package ele

import (
	"context"

	"github.com/acger/pair-api/internal/svc"
	"github.com/acger/pair-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListElementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListElementLogic {
	return ListElementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListElementLogic) ListElement(req types.EleListReq) (resp *types.EleListRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
