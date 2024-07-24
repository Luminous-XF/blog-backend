package middleware

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/response"
    "blog-backend/app/service"
    "blog-backend/pkg/jwt"
    "github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        tokenStr := ctx.Request.Header.Get("Authorization")
        if len(tokenStr) == 0 {
            code := error_code.AuthTokenNULL
            response.CommonFailed(code, code.String(), ctx)
            ctx.Abort()
            return
        }

        j := jwt.NewJWT()
        claims, code := j.ParseToken(tokenStr)
        if !code.IsSuccess() {
            response.CommonFailed(code, code.String(), ctx)
            ctx.Abort()
            return
        }

        if _, ok := service.IsUsernameExist(claims.Username); !ok {
            code = error_code.AuthFailed
            response.CommonFailed(code, code.String(), ctx)
            ctx.Abort()
            return
        }

        ctx.Set("claims", claims)
        ctx.Next()
    }
}
