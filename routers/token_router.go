package routers

import (
    v1 "blog-backend/app/controller/api/v1"
    "github.com/gin-gonic/gin"
)

func InitTokenPublicRouter(Router *gin.RouterGroup) {
    TokenRouter := Router.Group("tokens")
    {
        TokenRouter.POST("/username-password", v1.CreateTokenByUsernamePassword)
        TokenRouter.GET("/refresh-token", v1.CreateTokenByRefreshToken)
    }
}
