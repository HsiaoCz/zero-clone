package logic

import (
	"context"

	"github.com/HsiaoCz/zero-clone/product_service/internal/svc"
	"github.com/HsiaoCz/zero-clone/product_service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Product_serviceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProduct_serviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Product_serviceLogic {
	return &Product_serviceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Product_serviceLogic) Product_service(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	return &types.Response{Message: req.Name}, nil
}
