package handler

import (
	"net/http"

	"github.com/acger/pair-api/internal/logic/element"
	"github.com/acger/pair-api/internal/svc"
	"github.com/acger/pair-api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ListElementHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListElementLogic(r.Context(), ctx)
		resp, err := l.ListElement(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
