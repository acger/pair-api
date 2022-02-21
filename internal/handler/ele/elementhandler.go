package ele

import (
	"net/http"

	"github.com/acger/pair-api/internal/logic/ele"
	"github.com/acger/pair-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ElementHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := ele.NewElementLogic(r.Context(), svcCtx)
		resp, err := l.Element()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
