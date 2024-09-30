package allRoute

import (
	"github.com/gin-gonic/gin"
	"goPro4/routes/albumRoute"
	"goPro4/routes/userRoute"
)

func InitRouter(router *gin.Engine) {
	allRouter := router.Group("/v1")
	userRoute.InitUserRouter(allRouter)
	albumRoute.InitAlbumRouter(allRouter)
}
