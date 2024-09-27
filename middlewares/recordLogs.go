package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goPro4/utils"
)

func RecordLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		//	这里需要记录请求路径 和 访问用户
		c.Next()
		userResponseInfo := fmt.Sprintf("user_account: %s | request_url: %s | response_status: %d", "xiaoming", c.Request.URL, c.Writer.Status())
		utils.Info(userResponseInfo)
	}
}
