package logic

import (
	"Tstore/backend/product/rpc/product_info/productinfoclient"
	"context"

	"Tstore/backend/product/api/internal/svc"
	"Tstore/backend/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListLogic) GetList(req *types.GetListReq) (*types.GetListResp, error) {
	size := 5
	result, err := l.svcCtx.ProductInfoHandle.GetList(l.ctx, &productinfoclient.ListReq{
		Page:       int32(req.Page),
		Size:       int32(size),
		Category:   req.Category,
		PriceOrder: 0,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	resp := &types.GetListResp{
		List: result.List,
	}

	return resp, nil
}