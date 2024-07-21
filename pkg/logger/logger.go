package logger

import (
	"encoding/json"
	"go.uber.org/zap"
)

// logger
var logger *zap.Logger

func InitLogger() {
    initLogger()
}

// Dump 调试专用，不会中断程序，会在终端打印出 warning 消息。
// 第一个参数会使用 json.Marshal 进行渲染，第二个参数消息（可选）
// zapLogger.Dump(user.User{Name:"test"})
// zapLogger.Dump(user.User{Name:"test"}, "用户信息")
func Dump(value interface{}, msg ...string) {
    valueString := jsonString(value)
    // 判断第二个参数是否传参 msg
    if len(msg) > 0 {
        logger.Warn("Dump", zap.String(msg[0], valueString))
    } else {
        logger.Warn("Dump", zap.String("data", valueString))
    }
}

// LogIf 当 err != nil 时记录 error 等级的日志
func LogIf(err error) {
    if err != nil {
        logger.Error("Error Occurred:", zap.Error(err))
    }
}

// LogWarnIf 当 err != nil 时记录 warning 等级的日志
func LogWarnIf(err error) {
    if err != nil {
        logger.Warn("Error Occurred:", zap.Error(err))
    }
}

// LogInfoIf 当 err != nil 时记录 info 等级的日志
func LogInfoIf(err error) {
    if err != nil {
        logger.Info("Error Occurred:", zap.Error(err))
    }
}

// Debug 调试日志，详尽的程序日志
// 调用示例：
// zapLogger.Debug("Database", zap.String("sql", sql))
func Debug(moduleName string, fields ...zap.Field) {
    logger.Debug(moduleName, fields...)
}

// Info 告知类日志
func Info(moduleName string, fields ...zap.Field) {
    logger.Info(moduleName, fields...)
}

// Warn 警告类
func Warn(moduleName string, fields ...zap.Field) {
    logger.Warn(moduleName, fields...)
}

// Error 错误时记录，不应该中断程序，查看日志时重点关注
func Error(moduleName string, fields ...zap.Field) {
    logger.Error(moduleName, fields...)
}

// Fatal 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func Fatal(moduleName string, fields ...zap.Field) {
    logger.Fatal(moduleName, fields...)
}

// DebugString 记录一条字符串类型的 debug 日志，调用示例：
// zapLogger.DebugString("SMS", "短信内容", string(result.RawResponse))
func DebugString(moduleName, name, msg string) {
    logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
    logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
    logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
    logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
    logger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码。调用示例：
// zapLogger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func DebugJSON(moduleName, name string, value interface{}) {
    logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
    logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
    logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
    logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
    logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

func jsonString(value interface{}) string {
    b, err := json.Marshal(value)
    if err != nil {
        logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
    }
    return string(b)
}
