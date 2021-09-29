package handler

import (
	"net/http"

	"github.com/acger/pair-api/internal/logic/element"
	"github.com/acger/pair-api/internal/svc"
	"github.com/acger/pair-api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func SaveElementHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EleSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSaveElementLogic(r.Context(), ctx)
		resp, err := l.SaveElement(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
