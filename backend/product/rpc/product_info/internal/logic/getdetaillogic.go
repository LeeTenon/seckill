package logic

import (
	"context"

	"Tstore/backend/product/rpc/product_info/internal/svc"
	"Tstore/backend/product/rpc/product_info/product_info"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取商品详情
func (l *GetDetailLogic) GetDetail(in *product_info.DetailReq) (*product_info.DetailResp, error) {
	// todo: add your logic here and delete this line

	return &product_info.DetailResp{}, nil
}
