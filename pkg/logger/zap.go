// Package logger zapLogger工具配置
package logger

import (
    "blog-backend/config"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
    "strings"
    "time"
)

// initLogger 初始化日志工具
func initLogger() {
    cfg := config.CONFIG

    // 获取日志写入介质
    writeSyncer := getLogWriter(cfg)

    // 设置日志等级
    var logLevel zapcore.Level
    switch cfg.LogConfig.Level {
    case "debug":
        logLevel = zapcore.DebugLevel
    case "info":
        logLevel = zapcore.InfoLevel
    case "warn":
        logLevel = zapcore.WarnLevel
    case "error":
        logLevel = zapcore.ErrorLevel
    case "panic":
        logLevel = zapcore.PanicLevel
    case "fatal":
        logLevel = zapcore.FatalLevel
    default:
        panic("Config-LogConfig  log level error")
    }

    // 初始化core
    core := zapcore.NewCore(getEncoder(cfg), writeSyncer, logLevel)

    // 初始化 zapLogger
    logger = zap.New(core,
        zap.AddCaller(),
        zap.AddCallerSkip(1),
        zap.AddStacktrace(zap.ErrorLevel),
    )

    // 将自定义的 zapLogger 替换为全局的logger
    zap.ReplaceGlobals(logger)
}

// getLogWriter 设置日志写入介质
func getLogWriter(cfg *config.Config) zapcore.WriteSyncer {
    // 如果配置了按日期记录日志文件, 则重新定义文件名
    logCfg := &cfg.LogConfig
    var fileName = logCfg.FilePath + "/" + logCfg.FileName
    if logCfg.Type == "daily" {
        logName := time.Now().Format("2006-01-02") + ".log"
        fileName = strings.ReplaceAll(fileName, ".log", "_"+logName)
    }

    // 滚动日志
    lumberJackLogger := &lumberjack.Logger{
        Filename:   fileName,
        MaxSize:    logCfg.MaxSize,
        MaxBackups: logCfg.MaxBackup,
        MaxAge:     logCfg.MaxAge,
        Compress:   logCfg.Compress,
    }

    // 本地环境配置
    if cfg.AppConfig.Env == "local" {
        // 终端、日志文件同时记录
        return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
    }
    // 其它环境在日志文件记录
    return zapcore.AddSync(lumberJackLogger)
}

// getEncoder
func getEncoder(cfg *config.Config) zapcore.Encoder {
    // 日志格式规则
    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "zapLogger",
        CallerKey:      "caller",
        FunctionKey:    zapcore.OmitKey,
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding, // 每行日志结尾加换行 '\n'
        EncodeLevel:    zapcore.CapitalLevelEncoder,
        EncodeTime:     customTimeEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
    }

    // 本地环境配置
    if cfg.AppConfig.Env == "local" {
        // 终端输出关键词高亮显示
        // encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
        // 本地使用内置的 Console 解码器
        return zapcore.NewConsoleEncoder(encoderConfig)
    }
    // 其它环境使用 JSON 编码器
    return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
