package dao

import (
	"Tstore/models"
	"github.com/jinzhu/gorm"
)

var ProductDao = &productDao{}

type productDao struct {
}

func (*productDao) GetList(db *gorm.DB, limit, size int) ([]*models.ProductProfile, error) {
	// DB查询
	productList := make([]*models.ProductProfile, 5)
	rows, err := db.Table("product").Offset(limit).Limit(size).Select("id,name,category,price,title").Find(&productList).Rows()
	if err != nil {
		return nil, err
	} else if !rows.Next() {
		return nil, ErrorNoData
	}

	return productList, nil
}
