package svc

import (
	"Tstore/backend/product/rpc/seckill/internal/config"
	"Tstore/models"
	"github.com/zeromicro/go-zero/core/stores/redis"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
	Cache   *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	//启动Gorm支持
	db, err := gorm.Open("mysql", c.DataSourceName)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	//自动同步更新表结构
	if db.HasTable(&models.Order{}) { //判断表是否存在
		db.AutoMigrate(&models.Order{}) //存在就自动适配表，也就说原先没字段的就增加字段
	} else {
		db.CreateTable(&models.Order{}) //不存在就创建新表
	}

	return &ServiceContext{
		Config:  c,
		DbEngin: db,
		Cache:   redis.New(c.CacheRedis),
	}
}
