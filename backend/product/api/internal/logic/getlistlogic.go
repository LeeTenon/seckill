package logic

import (
	"Tstore/backend/product/api/internal/svc"
	"Tstore/backend/product/api/internal/types"
	"Tstore/backend/product/rpc/product_info/productinfoclient"
	"Tstore/models"
	"context"
	"fmt"
	"strconv"

	"encoding/json"
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
	// 查询缓存：1. 根据分类和页码缓存商品列表（热点数据）；2.价格排序页面进行缓存（快速排序）。
	fmt.Printf("list:%s:%s", req.Category, strconv.Itoa(req.Page))
	cc, err := l.svcCtx.Cache.Get(fmt.Sprintf("list:%s:%s", req.Category, strconv.Itoa(req.Page)))
	if err == nil && cc != "" {
		return getResp([]byte(cc))
	}

	// 查询DB列表
	size := 5
	result, err := l.svcCtx.ProductInfoHandle.GetList(l.ctx, &productinfoclient.ListReq{
		Page:       int32(req.Page),
		Size:       int32(size),
		Category:   req.Category,
		PriceOrder: 0, // 价格排序功能暂定
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	// 写入缓存
	l.svcCtx.Cache.Set(fmt.Sprintf("list:%s:%s", req.Category, strconv.Itoa(req.Page)), string(result.List))

	// 返回消息体
	return getResp(result.List)
}

func getResp(js []byte) (*types.GetListResp, error) {
	// 列表反序列化
	date := make([]*models.ProductProfile, 0)
	err := json.Unmarshal(js, &date)
	if err != nil {
		return nil, err
	}

	// 返回消息体
	resp := &types.GetListResp{
		List: date,
	}
	return resp, nil
}
