package v1

import (
	"blog-backend/app/common/error_code"
	"blog-backend/app/common/request"
	"blog-backend/app/common/response"
	"blog-backend/app/service"
	"github.com/gin-gonic/gin"
)

func GetUserInfoByUUID(ctx *gin.Context) {
	var requestData request.GetByUUIDRequest
	if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if responseDate, code := service.GetUserInfoByUUID(requestData); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.CommonSuccess(code, responseDate, error_code.ErrMsg(code), ctx)
	}
}

func CreateTokenByUsernamePassword(ctx *gin.Context) {
	var requestData request.LoginByUsernameAndPasswordRequest
	if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if responseData, code := service.LoginByUsernameAndPassword(requestData); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.Created(responseData, error_code.ErrMsg(code), ctx)
	}
}

func SendVerifyCodeWithEmail(ctx *gin.Context) {
	var requestData request.SendVerifyCodeWithEmailRequest
	if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if responseData, code := service.SendVerifyCodeWithEmail(requestData, ctx.GetHeader("Trace-Id")); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.CommonSuccess(code, responseData, error_code.ErrMsg(code), ctx)
	}
}
