/*
 * @Descripttion:Log文件
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-16 16:13:36
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-24 01:47:27
 */

package dlog

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

const (
	OUT_TYPE_FILE   = "file"
	OUT_TYPE_STDOUT = "stdout"
	LOG_LEVEL_INFO  = "info"
	LOG_LEVEL_DEBUG = "debug"
	LOG_LEVEL_FATAL = "fatal"
	LOG_LEVEL_WARN  = "warn"
	LOG_LEVEL_ERROR = "error"
)

func Init(logLevel string, outType string) {
	logrus.SetFormatter(&logrus.TextFormatter{})
	switch outType {
	case OUT_TYPE_FILE:
		logFile := "./runtime-logs/runtime-log.log"
		writer, _ := rotatelogs.New(
			logFile+".%Y%m%d",
			rotatelogs.WithLinkName(logFile),
			rotatelogs.WithMaxAge(time.Duration(72)*time.Hour),
		)
		logrus.SetOutput(writer)
	case OUT_TYPE_STDOUT:
	default:
		logrus.SetOutput(os.Stdout)
	}

	switch logLevel {
	case LOG_LEVEL_INFO:
		logrus.SetLevel(logrus.InfoLevel)
	case LOG_LEVEL_DEBUG:
		logrus.SetLevel(logrus.DebugLevel)
	case LOG_LEVEL_WARN:
		logrus.SetLevel(logrus.WarnLevel)
	case LOG_LEVEL_ERROR:
		logrus.SetLevel(logrus.ErrorLevel)
	case LOG_LEVEL_FATAL:
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func Debug(args ...interface{}) {
	logrus.Debug(args)
}

func Info(args ...interface{}) {
	logrus.Info(args)
}

func Error(args ...interface{}) {
	logrus.Error(args)
}

func Warn(args ...interface{}) {
	logrus.Warn(args)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args)
}
