package userRoute

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	fmt.Println(user)
	utils.Response(c, 200, "success", nil)
}

func userRegister(c *gin.Context) {
	utils.Response(c, 200, "success", nil)
}
