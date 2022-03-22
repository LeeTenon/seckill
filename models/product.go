package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	Category    string  `gorm:"type:varchar(20);index:category"`
	Storage     uint    `gorm:"default:0;not null"`
	Price       float64 `gorm:"Type:decimal;not null;index:price;index:category"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
}

type ProductProfile struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Name     string  `gorm:"not null" json:"name"`
	Category string  `gorm:"type:varchar(20);index:category" json:"category"`
	Price    float64 `gorm:"Type:decimal;not null;index:price;index:category" json:"price"` // 价格非空
	Title    string  `gorm:"not null" json:"title"`                                         // 设置 Num字段自增
}
