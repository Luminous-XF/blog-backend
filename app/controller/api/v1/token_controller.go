package v1

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/request"
    "blog-backend/app/common/response"
    "blog-backend/app/service"
    "github.com/gin-gonic/gin"
)

// CreateTokenByUsernamePassword 通过用户名密码登录获取 token
func CreateTokenByUsernamePassword(ctx *gin.Context) {
    var req request.LoginByUsernameAndPasswordRequest
    if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
        code := error_code.ParamBindError
        response.CommonFailed(code, code.String(), ctx)
        return
    }

    if rsp, code := service.LoginByUsernameAndPassword(&req); !code.IsSuccess() {
        response.CommonFailed(code, code.String(), ctx)
    } else {
        response.CommonSuccess(code, rsp, code.String(), ctx)
    }
}

// CreateTokenByRefreshToken 通过 Refresh Token 刷新 Access Token
func CreateTokenByRefreshToken(ctx *gin.Context) {
    var req request.CreateTokenByRefreshTokenRequest
    req.RefreshToken = ctx.GetHeader("refreshToken")
    if len(req.RefreshToken) == 0 {
        code := error_code.RefreshTokenInvalid
        response.CommonFailed(code, code.String(), ctx)
        return
    }

    if rsp, code := service.CreateTokenByRefreshToken(&req); !code.IsSuccess() {
        response.CommonFailed(code, code.String(), ctx)
    } else {
        response.CommonSuccess(code, rsp, code.String(), ctx)
    }
}
