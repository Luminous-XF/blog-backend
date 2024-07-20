package routers

import (
	v1 "blog-backend/app/controller/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPostPublicRouter(Router *gin.RouterGroup) {
	PostRouter := Router.Group("posts")
	{
		PostRouter.POST("", v1.GetPostList)
		PostRouter.POST("/uuid", v1.GetPostInfoByUUID)
	}
}

func InitPostPrivateRouter(Router *gin.RouterGroup) {

}
