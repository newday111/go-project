package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goPro4/utils"
)

func RecordLogs() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		//	这里需要记录请求路径 和 访问用户
		userResponseInfo := ""
		username, exists := c.Get("username")
		if exists {
			userResponseInfo = fmt.Sprintf("user_account: %s | request_url: %s | response_status: %d", username, c.Request.URL, c.Writer.Status())
		} else {
			userResponseInfo = "从上下文中获取不到用户账号"
		}
		utils.Info(userResponseInfo)
	}
}
