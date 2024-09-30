package albumRoute

import (
	"github.com/gin-gonic/gin"
	"goPro4/middlewares"
)

func InitAlbumRouter(Router *gin.RouterGroup) {
	albumRoute := Router.Group("/alb", middlewares.AuthMiddleware())
	{
		albumRoute.POST("/query/album", albumQuery)
		albumRoute.POST("/update/album", albumUpdate)
	}
}
