package middlewares

import (
	"gin-base/conf"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	// EncoderConfig
	TIME_KEY       = "time"
	LEVLE_KEY      = "level"
	NAME_KEY       = "logger"
	CALLER_KEY     = "caller"
	MESSAGE_KEY    = "msg"
	STACKTRACE_KEY = "stacktrace"

	// 每个日志文件保存的最大尺寸 单位：M
	MAX_SIZE = 10
)

// 设置日志级别、输出格式和日志文件的路径
func SetLogs() {
	logConf := conf.LogConf
	logLevel := parseLevel(logConf.Level)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        TIME_KEY,
		LevelKey:       LEVLE_KEY,
		NameKey:        NAME_KEY,
		CallerKey:      CALLER_KEY,
		MessageKey:     MESSAGE_KEY,
		StacktraceKey:  STACKTRACE_KEY,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 大写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器(相对路径+行号)
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志输出格式 (json / console)
	var encoder zapcore.Encoder
	switch logConf.LogFormat {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 添加日志切割归档功能
	hook := lumberjack.Logger{
		Filename:   logConf.Path,       // 日志文件路径
		MaxSize:    MAX_SIZE,           // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: logConf.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     logConf.MaxAge,     // 文件最多保存多少天
		Compress:   true,               // 是否压缩
	}

	var outPut zapcore.WriteSyncer

	if logConf.ToStd {
		outPut = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook))
	} else {
		outPut = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
	}

	core := zapcore.NewCore(
		encoder,                        // 编码器配置
		outPut,                         // 打印到控制台和文件
		zap.NewAtomicLevelAt(logLevel), // 日志级别
	)

	var logger *zap.Logger
	if logLevel == zap.DebugLevel {
		// 开启文件及行号
		caller := zap.AddCaller()
		// 开启开发模式，堆栈跟踪
		development := zap.Development()
		logger = zap.New(core, caller, development)
	} else {
		logger = zap.New(core)
	}

	// 将自定义的logger替换为全局的logger
	zap.ReplaceGlobals(logger)
}

func parseLevel(val string) zapcore.Level {
	switch val {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.DebugLevel
	}
}
