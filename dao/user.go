package dao

import (
	"Tstore/models"

	"github.com/jinzhu/gorm"
)

var UserDao = &userDao{}

type userDao struct {
}

func (*userDao) FondByName(db *gorm.DB, userName string) (*models.User, error) {
	user := &models.User{}

	if db.Table("users").Select("*").Where("user_name = ?", userName).Scan(user).RecordNotFound() {
		return nil, ErrorNoData
	}

	return user, nil
}
