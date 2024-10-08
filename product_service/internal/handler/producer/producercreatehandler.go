package producer

import (
	"net/http"

	"github.com/HsiaoCz/zero-clone/product_service/internal/logic/producer"
	"github.com/HsiaoCz/zero-clone/product_service/internal/svc"
	"github.com/HsiaoCz/zero-clone/product_service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProducerCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProducerCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := producer.NewProducerCreateLogic(r.Context(), svcCtx)
		resp, err := l.ProducerCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
