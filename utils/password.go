package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Password 密码加密
func Password(plainpwd string) string {
	//谷歌的加密包
	hash, err := bcrypt.GenerateFromPassword([]byte(plainpwd), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return encodePWD
}

// CheckPassword 密码校验
func CheckPassword(plainpwd, cryptedpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cryptedpwd), []byte(plainpwd)) //验证（对比）
	return err == nil
}
