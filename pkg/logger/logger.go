package logger

import (
	"blog-backend/config"
	"encoding/json"
	"go.uber.org/zap"
	"sync"
)

type Logger struct {
}

var (
	logger    *Logger
	zapLogger *zap.Logger
	once      sync.Once
)

func InitLogger(cfg *config.Config) *Logger {
	once.Do(func() {
		logger = &Logger{}
		initLogger(cfg)
	})
	return logger
}

// Dump 调试专用，不会中断程序，会在终端打印出 warning 消息。
// 第一个参数会使用 json.Marshal 进行渲染，第二个参数消息（可选）
// zapLogger.Dump(user.User{Name:"test"})
// zapLogger.Dump(user.User{Name:"test"}, "用户信息")
func (log Logger) Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	// 判断第二个参数是否传参 msg
	if len(msg) > 0 {
		zapLogger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		zapLogger.Warn("Dump", zap.String("data", valueString))
	}
}

// LogIf 当 err != nil 时记录 error 等级的日志
func (log Logger) LogIf(err error) {
	if err != nil {
		zapLogger.Error("Error Occurred:", zap.Error(err))
	}
}

// LogWarnIf 当 err != nil 时记录 warning 等级的日志
func (log Logger) LogWarnIf(err error) {
	if err != nil {
		zapLogger.Warn("Error Occurred:", zap.Error(err))
	}
}

// LogInfoIf 当 err != nil 时记录 info 等级的日志
func (log Logger) LogInfoIf(err error) {
	if err != nil {
		zapLogger.Info("Error Occurred:", zap.Error(err))
	}
}

// Debug 调试日志，详尽的程序日志
// 调用示例：
// zapLogger.Debug("Database", zap.String("sql", sql))
func (log Logger) Debug(moduleName string, fields ...zap.Field) {
	zapLogger.Debug(moduleName, fields...)
}

// Info 告知类日志
func (log Logger) Info(moduleName string, fields ...zap.Field) {
	zapLogger.Info(moduleName, fields...)
}

// Warn 警告类
func (log Logger) Warn(moduleName string, fields ...zap.Field) {
	zapLogger.Warn(moduleName, fields...)
}

// Error 错误时记录，不应该中断程序，查看日志时重点关注
func (log Logger) Error(moduleName string, fields ...zap.Field) {
	zapLogger.Error(moduleName, fields...)
}

// Fatal 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func (log Logger) Fatal(moduleName string, fields ...zap.Field) {
	zapLogger.Fatal(moduleName, fields...)
}

// DebugString 记录一条字符串类型的 debug 日志，调用示例：
// zapLogger.DebugString("SMS", "短信内容", string(result.RawResponse))
func (log Logger) DebugString(moduleName, name, msg string) {
	zapLogger.Debug(moduleName, zap.String(name, msg))
}

func (log Logger) InfoString(moduleName, name, msg string) {
	zapLogger.Info(moduleName, zap.String(name, msg))
}

func (log Logger) WarnString(moduleName, name, msg string) {
	zapLogger.Warn(moduleName, zap.String(name, msg))
}

func (log Logger) ErrorString(moduleName, name, msg string) {
	zapLogger.Error(moduleName, zap.String(name, msg))
}

func (log Logger) FatalString(moduleName, name, msg string) {
	zapLogger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码。调用示例：
// zapLogger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func (log Logger) DebugJSON(moduleName, name string, value interface{}) {
	zapLogger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func (log Logger) InfoJSON(moduleName, name string, value interface{}) {
	zapLogger.Info(moduleName, zap.String(name, jsonString(value)))
}

func (log Logger) WarnJSON(moduleName, name string, value interface{}) {
	zapLogger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func (log Logger) ErrorJSON(moduleName, name string, value interface{}) {
	zapLogger.Error(moduleName, zap.String(name, jsonString(value)))
}

func (log Logger) FatalJSON(moduleName, name string, value interface{}) {
	zapLogger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		zapLogger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
