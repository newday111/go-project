package middlewares

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(ginRoute *gin.Engine) {
	ginRoute.Use(RecordLogs())
}
