package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	Category    string  `gorm:"type:varchar(20);default:nothing;index:category"`
	Storage     uint    `gorm:"default:0;not null"`
	Price       float64 `gorm:"Type:decimal;not null;index:price;index:category"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
}

type ProductProfile struct {
	ID       uint    `gorm:"primary_key"`
	Name     string  `gorm:"not null"`
	Category string  `gorm:"type:varchar(20);default:nothing;index:category"`
	Price    float64 `gorm:"Type:decimal;not null;index:price;index:category"` // 价格非空
	Title    string  `gorm:"not null"`                                         // 设置 Num字段自增
}
