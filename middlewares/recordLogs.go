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
			cookieUserName, err := c.Cookie("username")
			if err != nil {
				utils.Error(err.Error())
			}
			//	从前端解密出来的token
			userResponseInfo = fmt.Sprintf("user_account: %s | request_url: %s | response_status: %d", cookieUserName, c.Request.URL, c.Writer.Status())
		}
		utils.Info(userResponseInfo)
	}
}
