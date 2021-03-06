package svc

import (
	"Tstore/backend/product/api/internal/config"
	"Tstore/backend/product/rpc/product_info/productinfoclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	ProductInfoHandle productinfoclient.ProductInfo
	Cache             *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		ProductInfoHandle: productinfoclient.NewProductInfo(zrpc.MustNewClient(c.ProductInfo)),
		Cache:             redis.New(c.CacheRedis),
	}
}
