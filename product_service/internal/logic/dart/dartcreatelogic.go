package dart

import (
	"context"

	"github.com/HsiaoCz/zero-clone/product_service/internal/svc"
	"github.com/HsiaoCz/zero-clone/product_service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DartCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDartCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DartCreateLogic {
	return &DartCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DartCreateLogic) DartCreate(req *types.DartCreateRequest) (resp *types.DartCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.DartCreateResponse{ProductName: req.ProductName + req.ProductCode}, nil
}
