package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Log struct {
	LoggerInfo *zap.Logger
	LoggerDebug *zap.Logger
	LoggerWarn *zap.Logger
	LoggerError *zap.Logger
}

var Logger Log

func init()  {
	root,err := os.Getwd()
	if err != nil{
		panic("获取目录失败")
	}
	Logger.LoggerInfo= InitLogs(root+"/storage/info.log","info")
	Logger.LoggerDebug= InitLogs(root+"/storage/debug.log","debug")
	Logger.LoggerWarn = InitLogs(root+"/storage/warn.log","warn")
	Logger.LoggerError=  InitLogs(root+"/storage/error.log","error")
}

func InitLogs(logPath string,logLevel string) *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    500,     // megabytes
		MaxBackups: 1,       // 最多保留300个备份
		MaxAge:     1,       // days
		Compress:   false,   // 是否压缩 disabled by default
	})
	var level zapcore.Level
	switch logLevel {
		case "info":
			level=zap.InfoLevel
			break
		case "debug":
			level=zap.DebugLevel
			break
		case "warn":
			level=zap.WarnLevel
			break
		case "error":
			level=zap.ErrorLevel
			break
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder :=zapcore.NewJSONEncoder(encoderConfig)
	core :=zapcore.NewCore(encoder,zapcore.NewMultiWriteSyncer(os.Stdout,w),level)
	return zap.New(core,zap.AddCaller())

}

func (log *Log) LogInfo(msg string,fields...zap.Field)  {
	log.LoggerInfo.Info(msg,fields...)
	defer log.LoggerInfo.Sync()
}

func (log *Log)LogDebug(msg string,fields...zap.Field)  {
	log.LoggerDebug.Debug(msg,fields...)
	defer log.LoggerDebug.Sync()
}

func (log *Log)LogWarn(msg string,fields...zap.Field)  {
	log.LoggerWarn.Warn(msg,fields...)
	defer log.LoggerWarn.Sync()
}

func (log *Log)LogError(msg string,fields...zap.Field)  {
	log.LoggerError.Error(msg,fields...)
	defer log.LoggerError.Sync()
}

