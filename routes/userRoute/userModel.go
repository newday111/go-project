package userRoute

import (
	"fmt"
	"goPro4/db"
)

type loginUser struct {
	UserName string `gorm:"username"`
	Token    string `gorm:"token"`
}

func (userInfo *loginUser) getUserInfoToName(username string) {
	//	指定查询的表名
	db.DB.Table("login_users").Select([]string{"name", "token"}).Where("name = ?", username).Find(&userInfo)
	fmt.Println(userInfo)
}
