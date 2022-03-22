package handler

import (
	"net/http"

	"Tstore/backend/product/api/internal/logic"
	"Tstore/backend/product/api/internal/svc"
	"Tstore/backend/product/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetListLogic(r.Context(), svcCtx)
		resp, err := l.GetList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
