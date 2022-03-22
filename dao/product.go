package dao

import (
	"Tstore/models"

	"github.com/jinzhu/gorm"
)

var ProductDao = &productDao{}

type productDao struct {
}

func (*productDao) GetList(db *gorm.DB, limit, size int) ([]*models.ProductProfile, error) {
	productList := make([]*models.ProductProfile, 0)

	println(limit, size)
	rows, err := db.Table("product").Select("id,name,category,price,title").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		productprofile := &models.ProductProfile{}
		err = rows.Scan(&productprofile.ID, &productprofile.Name, &productprofile.Category, &productprofile.Price, &productprofile.Title)
		if err != nil {
			println(err.Error())
		}
		println(productprofile.Name)
		productList = append(productList, productprofile)
	}
	return productList, nil
}
