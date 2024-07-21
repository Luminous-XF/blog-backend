package v1

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/request"
    "blog-backend/app/common/response"
    "blog-backend/app/service"
    "github.com/gin-gonic/gin"
)

// GetUserInfoByUUID 通过 UUID 获取用户信息
func GetUserInfoByUUID(ctx *gin.Context) {
    var req request.GetByUUIDRequest
    if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
        code := error_code.ParamBindError
        response.CommonFailed(code, code.String(), ctx)
        return
    }

    if rsp, code := service.GetUserInfoByUUID(&req); !code.IsSuccess() {
        response.CommonFailed(code, code.String(), ctx)
    } else {
        response.CommonSuccess(code, rsp, code.String(), ctx)
    }
}

// CreateTokenByUsernamePassword 通过用户名密码登录获取 token
func CreateTokenByUsernamePassword(ctx *gin.Context) {
    var req request.LoginByUsernameAndPasswordRequest
    if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
        code := error_code.ParamBindError
        response.CommonFailed(code, code.String(), ctx)
        return
    }

    if responseData, code := service.LoginByUsernameAndPassword(&req); !code.IsSuccess() {
        response.CommonFailed(code, code.String(), ctx)
    } else {
        response.CommonSuccess(code, responseData, code.String(), ctx)
    }
}

// GetRegisterVerifyCodeWithEmail 通过邮箱发送验证码
func GetRegisterVerifyCodeWithEmail(ctx *gin.Context) {
    var req request.GetRegisterVerifyCodeWithEmailRequest
    if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
        code := error_code.ParamBindError
        response.CommonFailed(code, code.String(), ctx)
        return
    }

    if responseData, code := service.GetRegisterVerifyCodeWithEmail(&req, ctx.GetHeader("Trace-Id")); !code.IsSuccess() {
        response.CommonFailed(code, code.String(), ctx)
    } else {
        response.CommonSuccess(code, responseData, code.String(), ctx)
    }
}

// CreateUserByEmailVerifyCode 通过邮箱验证码注册账号
func CreateUserByEmailVerifyCode(ctx *gin.Context) {
    var req request.CreateUserByEmailVerifyCodeRequest
    if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
        code := error_code.ParamBindError
        response.CommonFailed(code, code.String(), ctx)
        return
    }

    if rsp, code := service.CreateUserWithEmailVerifyCode(&req); !code.IsSuccess() {
        response.CommonFailed(code, code.String(), ctx)
    } else {
        response.CommonSuccess(code, rsp, code.String(), ctx)
    }
}
