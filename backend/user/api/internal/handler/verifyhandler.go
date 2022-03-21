package handler

import (
	"net/http"

	"Tstore/backend/user/api/internal/logic"
	"Tstore/backend/user/api/internal/svc"
	"Tstore/backend/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVerifyLogic(r.Context(), svcCtx)
		resp, err := l.Verify(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
