package userRoute

import (
	"github.com/gin-gonic/gin"
	"goPro4/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRoute := Router.Group("/user")
	{
		userRoute.POST("/login", userLogin)
		userRoute.POST("/register", userRegister, middlewares.AuthMiddleware())
	}
}
