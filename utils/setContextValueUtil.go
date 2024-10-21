package utils

import "github.com/gin-gonic/gin"

func SetContextValue(c *gin.Context, key string, value interface{}) {
	c.Set(key, value)
}
