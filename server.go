package main

import (
	"github.com/gin-gonic/gin"
	"goPro4/db"
	"goPro4/routes/allRoute"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func main() {
	DB, err = db.GetMysqlConnection()
	if err != nil {
		panic(err)
	}
	ginRouter := gin.Default()
	allRoute.InitRouter(ginRouter)
	ginRouter.Run(":8080")
}
