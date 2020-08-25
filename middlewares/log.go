package middlewares

import (
	"gin-base/conf"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
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
func varInit() {
	logConf  = conf.LogConf
	logLevel = parseLevel(logConf.Level)
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

func InitGinLogger(r *gin.Engine) {
	varInit()
	if logLevel == zap.DebugLevel {
		logPath := logConf.Path
		if strings.Contains(logPath, ".log") {
			logPath = strings.Replace(logPath, ".log", ".debug.log", -1)
		}
		logfile, _ := os.Create(logPath)
		if logConf.ToStd {
			gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
		} else {
			gin.DefaultWriter = io.MultiWriter(logfile)
		}
		gin.ForceConsoleColor()
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		zapLogger, zapRecovery := Log()
		r.Use(zapLogger)
		r.Use(zapRecovery)
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
		EncodeTime:     zapcore.RFC3339TimeEncoder, // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
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
	//logger = zap.New(core, zap.AddCaller())
	logger = zap.New(core)
	// 将自定义的logger替换为全局的logger
	zap.ReplaceGlobals(logger)
	return Ginzap(zap.L(), false), ginzap.RecoveryWithZap(zap.L(), true)
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

func Ginzap(logger *zap.Logger, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			logger.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
			)
		}
	}
}