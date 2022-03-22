package logic

import (
	"Tstore/backend/product/rpc/product_info/internal/svc"
	"Tstore/backend/product/rpc/product_info/product_info"
	"Tstore/dao"
	"context"
	"encoding/json"
	"errors"

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
	limit := (in.Page - 1) * in.Size
	result, err := dao.ProductDao.GetList(l.svcCtx.DbEngin, int(limit), int(in.Size))
	if err != nil {
		if err == dao.ErrorNoData {
			return nil, errors.New("无商品数据")
		}
		return nil, err
	}

	// 序列化
	date, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	// 返回消息体
	resp := &product_info.ListResp{
		List: date,
	}

	return resp, nil
}
