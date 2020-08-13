package middlewares

import (
	"gin-base/conf"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	logConf  = conf.LogConf
	logLevel = parseLevel(logConf.Level)
)

/**
 * @description: 给日志配置设置默认值
 * @author: Lorin
 * @time: 2020/8/13 上午10:35
 */
func init() {
	if logConf.Path == "" {
		logConf.Path = "./logs/gin-base.log"
	}
	if logConf.MaxSize == 0 {
		logConf.MaxSize = 20
	}
	if logConf.LogFormat == "" {
		logConf.LogFormat = "console"
	}
}

/**
 * @description: 配置zap日志、lumberjack日志切割归档，并将设置后的zap日志全局置入 后续的程序中如果需要写日志 直接zap.L()即可
 * @return: 日志中间件, recovery中间件
 * @author: Lorin
 * @time: 2020/8/13 上午10:35
 */
func Log() (gin.HandlerFunc, gin.HandlerFunc) {
	var logger *zap.Logger
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 如果是debug模式，采用CapitalColorLevelEncoder输出彩色的比较好看
	if logLevel == zap.DebugLevel {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
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
		MaxSize:    logConf.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: logConf.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     logConf.MaxAge,     // 文件最多保存多少天
		Compress:   true,               // 是否压缩
	}
	var outPut zapcore.WriteSyncer
	// 是否输出到控制台
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
	/**
	 * 是否跟踪堆栈：
	 * 	zap.AddStacktrace：跟踪对应级别的堆栈，打印时的输出形式较为友好，适用于开发环境下
	 * 	ginzap.RecoveryWithZap(zap.L(), true)：第二位参数表示是否跟踪堆栈，不过这个堆栈信息打印出来十分不友好，换行都没有，适用于生产环境下的记录
	 */
	stack := false
	if logLevel == zap.DebugLevel {
		// ErrorLevel 开发中报错了再打印堆栈
		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	} else {
		logger = zap.New(core, zap.AddCaller())
		stack = true
	}
	// 将自定义的logger替换为全局的logger
	zap.ReplaceGlobals(logger)
	return ginzap.Ginzap(zap.L(), time.RFC3339, false), ginzap.RecoveryWithZap(zap.L(), stack)
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
