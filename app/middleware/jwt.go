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
			response.Unauthorized(code, error_code.ErrMsg(code), ctx)
			ctx.Abort()
			return
		}

		j := jwt.NewJWT()
		claims, code := j.ParseToken(tokenStr)
		if !error_code.IsSuccess(code) {
			response.Unauthorized(code, error_code.ErrMsg(code), ctx)
			ctx.Abort()
			return
		}

		if _, ok := service.IsUsernameExist(claims.Username); !ok {
			code = error_code.AuthFailed
			response.Unauthorized(code, error_code.ErrMsg(code), ctx)
			ctx.Abort()
			return
		}

		// if claims.ExpiresAt.Unix()-time.Now().Unix() < global.CONFIG.JWTConfig.BufferTime {
		// 	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(global.CONFIG.JWTConfig.ExpiresTime)))
		//
		// }

		ctx.Set("claims", claims)
		ctx.Next()
	}
}