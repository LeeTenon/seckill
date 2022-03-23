package logic

import (
	"Tstore/backend/product/api/internal/svc"
	"Tstore/backend/product/api/internal/types"
	"Tstore/backend/product/rpc/product_info/productinfoclient"
	"Tstore/models"

	"context"
	"encoding/json"
	"fmt"
	"strconv"

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
	// 价格排序页处理
	if req.PriceOrder == 1 {
		priceOrderListIncr() // 价格增序
	} else if req.PriceOrder == 2 {
		priceOrderListDecr() // 价格降序
	}

	// 正常分类页处理
	// 查询缓存：1. 根据分类和页码缓存商品列表（热点数据）；2.价格排序页面进行缓存（快速排序）。
	cc, err := l.svcCtx.Cache.Get(fmt.Sprintf("list:%s:%s", req.Category, strconv.Itoa(req.Page)))
	if err == nil && cc != "" {
		return getResp([]byte(cc))
	} else if err != nil {
		return nil, err
	}

	// 查询DB列表
	size := 5 // 需移至前端请求中
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

func priceOrderListIncr() {

}

func priceOrderListDecr() {

}

// getResp 消息体返回函数
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
