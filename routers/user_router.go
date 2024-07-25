package routers

import (
    v1 "blog-backend/app/controller/api/v1"
    "github.com/gin-gonic/gin"
)

func InitUserPublicRouter(Router *gin.RouterGroup) {
    UserRouter := Router.Group("users")
    {
        UserRouter.POST("/verify-code/using-email", v1.GetRegisterVerifyCodeWithEmail)
        UserRouter.POST("", v1.CreateUserByEmailVerifyCode)
        UserRouter.GET("/:uuid", v1.GetUserInfoByUUID)
    }
}

func InitUserPrivateRouter(Router *gin.RouterGroup) {

}
