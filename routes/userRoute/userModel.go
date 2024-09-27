package userRoute

import (
	"goPro4/db"
)

const loginUserTableName = "login_users"

type loginUser struct {
	UserName string `gorm:"username"`
	Token    string `gorm:"token"`
	IsValid  int    `gorm: "is_valid"`
}

func (userInfo *loginUser) getUserInfoToName(username string) *loginUser {
	//	指定查询的表名
	db.DB.Table(loginUserTableName).Select([]string{"name", "token", "is_valid"}).Where("name = ?", username).Find(userInfo)
	return userInfo
}
