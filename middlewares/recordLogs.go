package middlewares

import (
	"github.com/gin-gonic/gin"
	"goPro4/utils"
)

func RecordLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		//	这里需要记录请求路径 和 访问用户
		utils.Info(c.Request.URL)
	}
}
