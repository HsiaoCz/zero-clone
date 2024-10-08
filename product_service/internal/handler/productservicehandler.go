package handler

import (
	"net/http"

	"github.com/HsiaoCz/zero-clone/product_service/internal/logic"
	"github.com/HsiaoCz/zero-clone/product_service/internal/svc"
	"github.com/HsiaoCz/zero-clone/product_service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Product_serviceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewProduct_serviceLogic(r.Context(), svcCtx)
		resp, err := l.Product_service(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
