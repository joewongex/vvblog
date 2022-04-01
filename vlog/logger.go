package vlog

import (
	"io"
	"log"
	"os"
	"vvblog/config"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

var (
	levels = map[string]logrus.Level{
		"panic": logrus.PanicLevel,
		"fatal": logrus.FatalLevel,
		"error": logrus.ErrorLevel,
		"warn":  logrus.WarnLevel,
		"info":  logrus.InfoLevel,
		"debug": logrus.DebugLevel,
		"trace": logrus.TraceLevel,
	}
	appLogger *logrus.Logger
)

func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.Bytes()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func init() {
	infoFile, err := os.OpenFile(config.Log.Filename+"_info.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("打开%s日志文件失败：%v", config.Log.Filename+"_info.log", err)
	}
	errorFile, err := os.OpenFile(config.Log.Filename+"_error.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("打开%s日志文件失败：%v", config.Log.Filename+"_error.log", err)
	}

	infoWriter := []io.Writer{infoFile}
	errorWriter := []io.Writer{errorFile}
	// if config.App.Env == "development" {
	// 	infoWriter = append(infoWriter, os.Stdout)
	// 	errorWriter = append(errorWriter, os.Stderr)
	// }

	appLogger = logrus.New()
	// appLogger.Formatter.(*logrus.TextFormatter).DisableColors = true
	appLogger.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		HideKeys:        true,
		NoColors:        true,
		NoFieldsColors:  true,
		CallerFirst:     true,
	})
	appLogger.ReportCaller = true
	if level, ok := levels[config.Log.Level]; ok {
		appLogger.SetLevel(level)
	} else {
		log.Printf("[warn]配置的日志等级%s未能识别，将使用默认日志等级\n", config.Log.Level)
	}

	appLogger.AddHook(&WriterHook{
		Writer: io.MultiWriter(infoWriter...),
		LogLevels: []logrus.Level{
			logrus.WarnLevel,
			logrus.InfoLevel,
			logrus.DebugLevel,
			logrus.TraceLevel,
		},
	})

	appLogger.AddHook(&WriterHook{
		Writer: io.MultiWriter(errorWriter...),
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		},
	})
}

func Fatalf(format string, args ...interface{}) {
	appLogger.Fatalf(format, args...)
}

func Error(args ...interface{}) {
	appLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	appLogger.Errorf(format, args...)
}

func Infof(format string, args ...interface{}) {
	appLogger.Infof(format, args...)
}
