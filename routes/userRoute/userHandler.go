package userRoute

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goPro4/utils"
)

type User struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Email    string `json:"email"`
}

func getUserStruct(c *gin.Context) (*User, error) {
	user := &User{}
	if err := c.BindJSON(user); err != nil {
		//	记录日志
		return nil, err
	}
	return user, nil
}

func userLogin(c *gin.Context) {
	user, err := getUserStruct(c)
	if err != nil {
		utils.Response(c, 400, "参数错误", nil)
		return
	}

	info := &loginUser{
		UserName: user.UserName,
	}

	userInfo := info.getUserInfoToName(user.UserName)
	c.Set("username", userInfo.UserName)

	cookPath := viper.GetString("cookies.cook_path")
	cookIp := viper.GetString("cookies.cook_ip")
	c.SetCookie("username", "小明", 3600, cookPath, cookIp, false, false)

	utils.Response(c, 200, "success", nil)
}

func userRegister(c *gin.Context) {
	utils.Response(c, 200, "success", nil)
}
