package models

import (
	"Tstore/utils"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `db:"user_name"` // 用户名
	Password string `db:"password"`  // 密码
	Avatar   string `db:"avagotar"`  // 头像
	DelFlag  int64  `db:"del_flag"`  // 删除标志
}

// BeforeCreate 在创建前检验验证一下密码的有效性
func (u *User) BeforeCreate(db *gorm.DB) error {
	if len(u.Password) < 6 {
		return errors.New("密码太简单了")
	}
	//对密码进行加密存储
	u.Password = utils.Password(u.Password)
	return nil
}
