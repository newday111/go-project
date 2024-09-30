package albumRoute

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goPro4/utils"
)

func albumQuery(c *gin.Context) {
	fmt.Println("albumQuery")
	utils.Response(c, utils.RespGetDataSuccess, utils.RespMsg[utils.RespGetDataSuccess], nil)
}

func albumUpdate(c *gin.Context) {
	fmt.Println("albumQuery")
	utils.Response(c, utils.RespGetDataSuccess, utils.RespMsg[utils.RespGetDataSuccess], nil)
}
