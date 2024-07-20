package v1

import (
	"blog-backend/app/common/error_code"
	"blog-backend/app/common/request"
	"blog-backend/app/common/response"
	"blog-backend/app/service"
	"github.com/gin-gonic/gin"
)

func GetPostList(ctx *gin.Context) {
	var requestData request.PageInfoRequest
	if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if responseData, code := service.GetPostList(requestData); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.CommonSuccess(code, responseData, error_code.ErrMsg(code), ctx)
	}
}

func GetPostInfoByUUID(ctx *gin.Context) {
	var requestData request.GetByUUIDRequest
	if err := ctx.ShouldBindBodyWithJSON(&requestData); err != nil {
		response.CommonFailed(error_code.ParamBindError, error_code.ErrMsg(error_code.ParamBindError), ctx)
		return
	}

	if responseData, code := service.GetPostByUUID(requestData); !error_code.IsSuccess(code) {
		response.CommonFailed(code, error_code.ErrMsg(code), ctx)
	} else {
		response.CommonSuccess(code, responseData, error_code.ErrMsg(code), ctx)
	}
}
