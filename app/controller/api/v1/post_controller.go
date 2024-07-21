package v1

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/request"
    "blog-backend/app/common/response"
    "blog-backend/app/service"
    "fmt"
    "github.com/gin-gonic/gin"
)

func GetPostList(ctx *gin.Context) {
    var req request.PageInfoRequest
    if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
        code := error_code.ParamBindError
        response.CommonFailed(code, code.String(), ctx)
        return
    }

    if rsp, code := service.GetPostList(req); !code.IsSuccess() {
        response.CommonFailed(code, code.String(), ctx)
    } else {
        response.CommonSuccess(code, rsp, code.String(), ctx)
    }
}

func GetPostInfoByUUID(ctx *gin.Context) {
    var req request.GetByUUIDRequest
    if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
        code := error_code.ParamBindError
        response.CommonFailed(code, code.String(), ctx)
        return
    }

    if rsp, code := service.GetPostByUUID(req); !code.IsSuccess() {
        fmt.Printf("Failed: %#v\n", rsp)
        response.CommonFailed(code, code.String(), ctx)
    } else {
        fmt.Printf("Success: %#v\n", rsp)
        response.CommonSuccess(code, rsp, code.String(), ctx)
    }
}
