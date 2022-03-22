package logic

import (
	"context"

	"Tstore/backend/product/rpc/product_info/internal/svc"
	"Tstore/backend/product/rpc/product_info/product_info"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取商品列表
func (l *GetListLogic) GetList(in *product_info.ListReq) (*product_info.ListResp, error) {
	// todo: add your logic here and delete this line

	return &product_info.ListResp{}, nil
}
