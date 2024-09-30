package main

import (
	"github.com/gin-gonic/gin"
	"goPro4/db"
	"goPro4/middlewares"
	"goPro4/routes/allRoute"
	"goPro4/utils"
)

func main() {
	db.GetMysqlConnection()
	utils.InitLogger()
	ginRouter := gin.Default()
	middlewares.InitMiddleware(ginRouter)
	allRoute.InitRouter(ginRouter)
	ginRouter.Run(":8080")
}
