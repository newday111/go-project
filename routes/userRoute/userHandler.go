package userRoute

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goPro4/utils"
	"time"
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
		utils.Response(c, utils.RespParamsError, utils.RespMsg[utils.RespParamsError], nil)
		return
	}

	info := &loginUser{
		UserName: user.UserName,
	}

	userInfo := info.getUserInfoToName(user.UserName)

	jwtKey := []byte(viper.GetString("login_auth.secret_key"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userInfo.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为24小时后
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		utils.Response(c, utils.RespCreateTokenFail, utils.RespMsg[utils.RespCreateTokenFail], nil)
		return
	}
	c.Header("Authorization", tokenString)
	c.Set("username", userInfo.UserName)

	cookPath := viper.GetString("cookies.cook_path")
	cookIp := viper.GetString("cookies.cook_ip")
	c.SetCookie("username", "小明", 3600, cookPath, cookIp, false, false)

	utils.Response(c, utils.RespSuccess, utils.RespMsg[utils.RespSuccess], nil)
}

func userRegister(c *gin.Context) {
	utils.Response(c, 200, "success", nil)
}
