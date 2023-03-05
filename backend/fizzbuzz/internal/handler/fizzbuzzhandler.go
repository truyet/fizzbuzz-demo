package handler

import (
	"net/http"

	"fizzbuzz/internal/logic"
	"fizzbuzz/internal/svc"
	"fizzbuzz/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FizzbuzzHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FizzbuzzRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFizzbuzzLogic(r.Context(), svcCtx)
		resp, err := l.Fizzbuzz(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
