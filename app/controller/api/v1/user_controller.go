package v1

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/request"
    "blog-backend/app/common/response"
    "blog-backend/app/service"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

//
// GetUserInfoByUUID
//  @Description: 通过 UUID 获取用户信息
//  @param ctx
//
func GetUserInfoByUUID(ctx *gin.Context) {
    var req request.GetByUUIDRequest
    var err error
    req.UUID, err = uuid.Parse(ctx.Param("uuid"))
    if err != nil {
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

//
// GetRegisterVerifyCodeWithEmail
//  @Description: 通过邮箱发送验证码
//  @param ctx
//
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

//
// CreateUserByEmailVerifyCode
//  @Description: 通过邮箱验证码注册账号
//  @param ctx
//
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
