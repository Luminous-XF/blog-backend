package routers

import (
	v1 "blog-backend/app/controller/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserPublicRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.POST("/uuid", v1.GetUserInfoByUUID)
		UserRouter.POST("/verify-code/using-email", v1.SendVerifyCodeWithEmail)
	}
}

func InitUserPrivateRouter(Router *gin.RouterGroup) {

}
