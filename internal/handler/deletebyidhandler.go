package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"p8/internal/logic"
	"p8/internal/svc"
	"p8/internal/types"
)

func DeleteByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.User
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteByIdLogic(r.Context(), svcCtx)
		resp, err := l.DeleteById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
