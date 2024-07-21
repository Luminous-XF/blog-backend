package response

import (
    "blog-backend/app/common/error_code"
    "github.com/gin-gonic/gin"
)

// Response 请求响应结构体
type Response struct {
    Code    error_code.ErrorCode `json:"code,omitempty"`    // 错误码(状态码)
    Data    interface{}          `json:"data,omitempty"`    // 响应数据
    Msg     string               `json:"msg,omitempty"`     // 响应信息
    TraceId string               `json:"traceId,omitempty"` // 请求 ID
}

// CommonSuccess 请求处理成功
func CommonSuccess(code error_code.ErrorCode, data interface{}, msg string, ctx *gin.Context) {
    ctx.JSON(code.Status(), Response{
        Code:    code,
        Data:    data,
        Msg:     msg,
        TraceId: ctx.GetHeader("Trace-Id"),
    })
}

// SuccessWithMessage 请求处理成功, 仅返回 Msg
func SuccessWithMessage(code error_code.ErrorCode, msg string, ctx *gin.Context) {
    ctx.JSON(code.Status(), Response{
        Code:    code,
        Msg:     msg,
        TraceId: ctx.GetHeader("Trace-Id"),
    })
}

// CommonFailed 请求处理失败
func CommonFailed(code error_code.ErrorCode, msg string, ctx *gin.Context) {
    ctx.JSON(code.Status(), Response{
        Code:    code,
        Msg:     msg,
        TraceId: ctx.GetHeader("Trace-Id"),
    })
}

// SuccessNoContent 请求成功, 不返回任何内容
func SuccessNoContent(code error_code.ErrorCode, ctx *gin.Context) {
    ctx.JSON(code.Status(), Response{
        Code:    code,
        TraceId: ctx.GetHeader("Trace-Id"),
    })
}
