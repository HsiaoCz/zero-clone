package producer

import (
	"context"

	"github.com/HsiaoCz/zero-clone/product_service/internal/svc"
	"github.com/HsiaoCz/zero-clone/product_service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProducerCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProducerCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProducerCreateLogic {
	return &ProducerCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProducerCreateLogic) ProducerCreate(req *types.ProducerCreateRequest) (resp *types.ProducerCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.ProducerCreateResponse{ProducerName: req.ProducerName + req.ProducerCode}, nil
}
