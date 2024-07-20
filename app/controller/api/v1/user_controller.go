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
    var requestData request.GetByUUIDRequest
    if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
        response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
        return
    }

    if responseDate, code := service.GetUserInfoByUUID(&requestData); !error_code.IsSuccess(code) {
        response.CommonFailed(code, error_code.ErrMsg(code), ctx)
    } else {
        response.CommonSuccess(code, responseDate, error_code.ErrMsg(code), ctx)
    }
}

// CreateTokenByUsernamePassword 通过用户名密码登录获取 token
func CreateTokenByUsernamePassword(ctx *gin.Context) {
    var requestData request.LoginByUsernameAndPasswordRequest
    if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
        response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
        return
    }

    if responseData, code := service.LoginByUsernameAndPassword(&requestData); !error_code.IsSuccess(code) {
        response.CommonFailed(code, error_code.ErrMsg(code), ctx)
    } else {
        response.Created(responseData, error_code.ErrMsg(code), ctx)
    }
}

// GetRegisterVerifyCodeWithEmail 通过邮箱发送验证码
func GetRegisterVerifyCodeWithEmail(ctx *gin.Context) {
    var requestData request.GetRegisterVerifyCodeWithEmailRequest
    if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
        response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
        return
    }

    if responseData, code := service.SendVerifyCodeWithEmail(&requestData, ctx.GetHeader("Trace-Id")); !error_code.IsSuccess(code) {
        response.CommonFailed(code, error_code.ErrMsg(code), ctx)
    } else {
        response.CommonSuccess(code, responseData, error_code.ErrMsg(code), ctx)
    }
}

func CreateUserByEmailVerifyCode(ctx *gin.Context) {
    var requestData request.CreateUserByEmailVerifyCodeRequest
    if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
        response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
        return
    }

    if code := service.CreateUserWithEmailVerifyCode(&requestData); !error_code.IsSuccess(code) {
        response.CommonFailed(code, error_code.ErrMsg(code), ctx)
    } else {
        response.Created(nil, error_code.ErrMsg(code), ctx)
    }
}
