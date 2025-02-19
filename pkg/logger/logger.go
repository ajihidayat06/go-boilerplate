package logger

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

var log = logrus.New()

func InitLogger() {
	log.SetFormatter(&logrus.JSONFormatter{})
}

func Info(message string, fields map[string]interface{}) {
	log.WithFields(fields).Info(message)
}

func Error(message string, err error) {
	pc, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	log.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"error":     err,
		"file":      file,
		"line":      line,
		"function":  funcName,
	}).Error(message)
}
