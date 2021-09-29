package handler

import (
	"net/http"

	"github.com/acger/pair-api/internal/logic/element"
	"github.com/acger/pair-api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func ElementHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewElementLogic(r.Context(), ctx)
		resp, err := l.Element()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
