package glog

import (
	"github.com/leonscriptcc/goddd/infrastructure/gconfig"
	"github.com/leonscriptcc/goddd/infrastructure/tools"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var logger *zap.Logger

// Init 日志初始化
func Init() (err error) {
	if gconfig.Parameters.Mode == tools.ENV_DEV {
		// 开发环境日志输出到终端展示
		logger, err = zap.NewDevelopment()
		if err != nil {
			return err
		}
	} else {
		// 自定义encoder
		cfg := zap.NewProductionConfig()

		// 编写hook
		// 修改输出时间的格式
		cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
		}

		// 自定义 info zap core
		// hook 确定输出的日志级别
		infoLV := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zap.ErrorLevel
		})
		// 日志rotate
		infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   gconfig.Parameters.ZapLogConfig.InfoLogPath, //日志文件存放目录，如果文件夹不存在会自动创建
			MaxSize:    gconfig.Parameters.ZapLogConfig.MaxSize,     //文件大小限制,单位MB
			MaxBackups: gconfig.Parameters.ZapLogConfig.MaxBackups,  //最大保留日志文件数量
			MaxAge:     gconfig.Parameters.ZapLogConfig.MaxAge,      //日志文件保留天数
			Compress:   gconfig.Parameters.ZapLogConfig.Compress,    //是否压缩处理
		})
		// 创建zap
		infoCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(infoFileWriteSyncer),
			infoLV,
		)

		// 自定义 err zap core
		// 日志rotate
		errFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   gconfig.Parameters.ZapLogConfig.ErrLogPath, //日志文件存放目录，如果文件夹不存在会自动创建
			MaxSize:    gconfig.Parameters.ZapLogConfig.MaxSize,    //文件大小限制,单位MB
			MaxBackups: gconfig.Parameters.ZapLogConfig.MaxBackups, //最大保留日志文件数量
			MaxAge:     gconfig.Parameters.ZapLogConfig.MaxAge,     //日志文件保留天数
			Compress:   gconfig.Parameters.ZapLogConfig.Compress,   //是否压缩处理
		})
		// hook 确定输出的日志级别
		errLV := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zap.ErrorLevel
		})
		// 创建zap
		errCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(errFileWriteSyncer),
			errLV,
		)

		// 创建自定义logger
		logger = zap.New(zapcore.NewTee(infoCore, errCore), zap.AddCaller())
	}
	return nil
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
