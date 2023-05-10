package handler

import (
	"net/http"

	"p8/internal/logic"
	"p8/internal/svc"
	"p8/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestId
		ids, ok := r.URL.Query()["id"]
		if !ok || len(ids) < 1 {
			httpx.ErrorCtx(r.Context(), w, nil)
			return
		}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
