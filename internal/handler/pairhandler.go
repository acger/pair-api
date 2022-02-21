package handler

import (
	"net/http"

	"github.com/acger/pair-api/internal/logic"
	"github.com/acger/pair-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func pairHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPairLogic(r.Context(), svcCtx)
		resp, err := l.Pair()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
