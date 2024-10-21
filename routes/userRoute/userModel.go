package userRoute

import (
	"goPro4/db"
)

const loginUserTableName = "user_info"

type loginUser struct {
	Name    string `gorm:"name"`
	Token   string `gorm:"token"`
	Email   string `gorm:"email"`
	IsValid int    `gorm: "is_valid"`
}

func (userInfo *loginUser) getUserInfoToName(username, email string) {
	//	指定查询的表名
	db.DB.Table(loginUserTableName).Select([]string{"name", "email", "is_valid"}).Where("name = ?", username).Or("email = ?", email).Find(userInfo)
}

func (userInfo *loginUser) registerUser(user *User) (int64, error) {
	//	指定字段创建用户
	createResult := db.DB.Table(loginUserTableName).Select("name", "pass_word", "email").Create(user)
	if createResult.Error != nil {
		return 0, createResult.Error
	}
	return createResult.RowsAffected, nil

}
