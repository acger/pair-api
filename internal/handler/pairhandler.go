package handler

import (
	"net/http"

	"github.com/acger/pair-api/internal/logic"
	"github.com/acger/pair-api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func pairHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPairLogic(r.Context(), ctx)
		resp, err := l.Pair()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
