package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"p8/internal/logic"
	"p8/internal/svc"
)

func GetAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetAllLogic(r.Context(), svcCtx)
		resp, err := l.GetAll()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
