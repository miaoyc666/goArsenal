package logger

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"time"
)

/*
File name    : logger.go
Author       : miaoyc
Create time  : 2021/12/21 12:45 上午
Update time  : 2023/07/27 12:45 上午
Description  : 日志文件相关
*/

/*
logrus内置日志级别如下：
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}
*/

var (
	logger *logrus.Logger
)

func getLog(logFilePath, logLevel string, rotationTime, maxAge int) *logrus.Logger {
	writer, err := rotateLogs.New(
		logFilePath+".%Y%m%d%H%M",
		rotateLogs.WithLinkName(logFilePath),                               // 生成软链，指向最新日志文件
		rotateLogs.WithMaxAge(time.Duration(maxAge)*time.Hour),             // 文件最大保存时间
		rotateLogs.WithRotationTime(time.Duration(rotationTime)*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		panic(err.Error())
	}
	// 设置日志级别
	var level logrus.Level
	level.UnmarshalText([]byte(logLevel))
	if err != nil {
		panic(err.Error())
	}
	logger = &logrus.Logger{
		Out:   writer,
		Level: level,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05+800",
			LogFormat:       "[%time%][%lvl%]%msg%\n",
		},
	}
	return logger
}

// Setup initialize the log instance
// log level: panic, fatal, error, warn/warning, info, debug, trace
func Setup(logPath, logLevel string, rotationTime, maxAge int) {
	logger = getLog(logPath, logLevel, rotationTime, maxAge)
}

// Panic output logs at panic level
func Panic(v ...interface{}) {
	logger.Panic(v...)
}

// Panicf output logs at panic level
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Fatalf output logs at fatal level
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Errorf output logs at error level
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Warnf output logs at warn level
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Infof output logs at info level
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Debugf output logs at debug level
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Trace output logs at trace level
func Trace(v ...interface{}) {
	logger.Trace(v...)
}

// Tracef output logs at trace level
func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}
