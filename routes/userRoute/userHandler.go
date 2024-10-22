package userRoute

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goPro4/db"
	"goPro4/utils"
	"time"
)

type User struct {
	Name             string `json:"name"`
	PassWord         string `json:"password"`
	Email            string `json:"email"`
	VerificationCode string `json:"code"`
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
	userInfo.getUserInfoToName(user.Name, user.Name)

	if userInfo.Name == "" {
		utils.Error(utils.UserRespMsg[utils.UserQueryFailed] + user.Name)
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
		utils.Error(utils.UserRespMsg[utils.UserTokenCreateFailed] + "=====>" + user.Name)
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
	user, err := getUserStruct(c)
	if err != nil {
		utils.Error(utils.UserRespMsg[utils.UserParamsAnalysisFailed])
		utils.Response(c, utils.RecordLogGetFail, utils.UserParamsAnalysisFailed, utils.UserRespMsg[utils.UserParamsAnalysisFailed], nil)
		return
	}
	utils.SetContextValue(c, "username", user.Name+"<===>"+user.Email)

	//	用户名和邮箱都必须是唯一的
	userInfo := &loginUser{}

	userInfo.getUserInfoToName(user.Name, user.Email)
	if userInfo.Name != "" {
		utils.Info("已经存在该用户" + "====" + user.Name + "====" + user.Email)
		utils.Response(c, utils.RecordLogCreateFail, utils.UserCreateFailed, utils.UserRespMsg[utils.UserCreateFailed], nil)
		return
	}

	//	获取密码进行加密
	user.PassWord, err = utils.EncryptionPassWord(user.PassWord, user.Name)
	if err != nil {
		utils.Error(utils.UserRespMsg[utils.UserParamsAnalysisFailed] + "====" + "")
		return
	}

	_, err = userInfo.registerUser(user)
	if err != nil {
		utils.Error(err.Error() + "====" + user.Name + "====" + "注册失败")
		utils.Response(c, utils.RecordLogCreateFail, utils.UserRegisterFailed, utils.UserRespMsg[utils.UserRegisterFailed], nil)
		return
	}
	utils.Response(c, utils.RecordLogSuccess, utils.UserRegisterSuccess, utils.UserRespMsg[utils.UserRegisterSuccess], nil)
}

func userLoginVerificationCode(c *gin.Context) {
	user, err := getUserStruct(c)
	if err != nil {
		utils.Error(utils.UserRespMsg[utils.UserParamsAnalysisFailed])
		utils.Response(c, utils.RecordLogGetFail, utils.UserParamsAnalysisFailed, utils.UserRespMsg[utils.UserParamsAnalysisFailed], nil)
		return
	}
	ctx := context.Background()
	emailToCode, err := db.Rdb.Get(ctx, user.Email).Result()
	if err != nil {
		utils.Error(err.Error() + "====>" + user.Email)
		return
	}

	if emailToCode == "" {
		code := utils.RandomGenerateVerificationCode(5)
		//	存储验证码到redis中,并且设置过期时间、
		db.Rdb.Set(c, user.Email, code, 5*time.Minute)
		go utils.SendEmail(user.Email, "机壳空间验证码", "验证码为:"+code+",验证码有效时间为5分钟")
		utils.SetContextValue(c, "username", user.Email)
		utils.Response(c, utils.RecordLogSuccess, utils.VerificationCodeSendSuccess, utils.VerificationCodeRespMsg[utils.VerificationCodeSendSuccess], nil)
		return
	}

	utils.Response(c, utils.RecordLogGetFail, utils.VerificationCodeExist, utils.VerificationCodeRespMsg[utils.VerificationCodeExist], nil)

}
