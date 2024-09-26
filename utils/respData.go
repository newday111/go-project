package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, code int, msg string, data interface{}) {
	resp := response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}
