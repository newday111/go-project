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

	//	get the params passed in the front-end

	user, err := getUserStruct(c)
	if err != nil {
		utils.Error(utils.UserRespMsg[utils.UserParamsAnalysisFailed])
		utils.Response(c, utils.RecordLogGetFail, utils.UserParamsAnalysisFailed, utils.UserRespMsg[utils.UserParamsAnalysisFailed], nil)
		return
	}

	//	query user information and verify it

	userInfo := &loginUser{}
	userInfo.getUserInfoToName(user.UserName)

	if userInfo.Name == "" {
		utils.Error(utils.UserRespMsg[utils.UserQueryFailed] + user.UserName)
		utils.Response(c, utils.RecordLogGetFail, utils.UserQueryFailed, utils.UserRespMsg[utils.UserQueryFailed], nil)
		return
	}

	// generate tokens with user information

	jwtKey := []byte(viper.GetString("login_auth.secret_key"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userInfo.Name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为24小时后
	})

	//	convert the generate token into a string

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		utils.Error(utils.UserRespMsg[utils.UserTokenCreateFailed] + "=====>" + user.UserName)
		utils.Response(c, utils.RecordLogCreateFail, utils.UserTokenCreateFailed, utils.UserRespMsg[utils.UserTokenCreateFailed], nil)
		return
	}

	c.Header("Authorization", tokenString)
	c.Set("username", userInfo.Name)

	cookPath := viper.GetString("cookies.cook_path")
	cookIp := viper.GetString("cookies.cook_ip")
	c.SetCookie("username", "小明", 3600, cookPath, cookIp, false, false)
	utils.Response(c, utils.RecordLogSuccess, utils.UserLoginSuccess, utils.UserRespMsg[utils.UserLoginSuccess], nil)
}

func userRegister(c *gin.Context) {
	//utils.Response(c,200, "success", nil)
}
