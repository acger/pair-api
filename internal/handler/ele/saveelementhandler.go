package ele

import (
	"net/http"

	"github.com/acger/pair-api/internal/logic/ele"
	"github.com/acger/pair-api/internal/svc"
	"github.com/acger/pair-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveElementHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EleSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := ele.NewSaveElementLogic(r.Context(), svcCtx)
		resp, err := l.SaveElement(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
