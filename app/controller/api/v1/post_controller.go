package v1

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/request"
    "blog-backend/app/common/response"
    "blog-backend/app/service"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

//
// GetPostList
//  @Description: 获取帖子列表
//  @param ctx
//
func GetPostList(ctx *gin.Context) {
    var req request.PageInfoRequest
    if err := ctx.ShouldBindQuery(&req); err != nil {
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

//
// GetPostInfoByUUID
//  @Description: 通过帖子 UUID 查找帖子
//  @param ctx
//
func GetPostInfoByUUID(ctx *gin.Context) {
    var req request.GetByUUIDRequest
    var err error
    req.UUID, err = uuid.Parse(ctx.Param("uuid"))
    if err != nil {
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
