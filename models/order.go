package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	ProductID   uint    `gorm:"not null"`
	ProductName string  `gorm:"not null"`
	Num         uint    `gorm:"not null"`
	Price       float64 `gorm:"Type:decimal;not null"`
}
