package middleware

import (
	"blog-backend/global"
	"blog-backend/pkg/helper"
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"strconv"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取 response 内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
		ctx.Writer = w

		// 获取请求数据
		var requestBody []byte
		if ctx.Request.Body != nil {
			// ctx.Request.Body 是一个 buffer 对象, 只能读取一次
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			// 读取后, 重新赋值 ctx.Request.Body, 以供后续的其它操作
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		start := time.Now()
		ctx.Next()

		// 开始记录日志
		cost := time.Since(start)
		responseStatus := ctx.Writer.Status()

		logFields := []zap.Field{
			zap.String("traceId", ctx.GetHeader("Trace-Id")),
			zap.Int("status", responseStatus),
			zap.String("request", ctx.Request.Method+" "+ctx.Request.URL.String()),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helper.MicrosecondsStr(cost)),
		}

		if ctx.Request.Method == "POST" ||
			ctx.Request.Method == "PUT" ||
			ctx.Request.Method == "PATCH" {
			// 请求内容
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))

			// 响应数据
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if responseStatus >= 400 && responseStatus < 500 {
			global.Logger.Warn("HTTP Warning received with "+" status: "+strconv.Itoa(responseStatus)+" ", logFields...)
		} else if responseStatus >= 500 && responseStatus < 600 {
			global.Logger.Error("HTTP Error received with "+" status: "+strconv.Itoa(responseStatus)+" ", logFields...)
		} else {
			global.Logger.Info("HTTP Access log", logFields...)
		}
	}
}
