package main

import (
	"github.com/gin-gonic/gin"
	"goPro4/db"
	"goPro4/routes/allRoute"
)

func main() {
	db.GetMysqlConnection()
	ginRouter := gin.Default()
	allRoute.InitRouter(ginRouter)
	ginRouter.Run(":8080")
}
