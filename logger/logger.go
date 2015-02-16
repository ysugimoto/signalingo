package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	NONE = iota
	INFO
	WARN
	FATAL
	ALL
)

var logLevel = NONE
var logLevelStr = "none"

func SetLevel(level string) {
	switch level {
	case "none":
		logLevel = NONE
	case "info":
		logLevel = INFO
	case "warn":
		logLevel = WARN
	case "fatal":
		logLevel = FATAL
	case "all":
		logLevel = ALL
	}
	logLevelStr = level
}

func SetLogFile(path string) {
	if fs, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		panic(fmt.Sprintf("%v", err))
	} else {
		log.SetOutput(fs)
	}
}

func Info(message string) {
	if logLevel < 1 {
		return
	}

	log.Printf("[%s] %s", logLevelStr, message)
}

func Infof(format string, params ...interface{}) {
	if logLevel < 1 {
		return
	}

	log.Printf("[Info] "+format, params...)
}

func Warn(message string) {
	if logLevel < 2 {
		return
	}

	log.Printf("[%s] %s", logLevelStr, message)
}

func Warnf(format string, params ...interface{}) {
	if logLevel < 1 {
		return
	}

	log.Printf("[Warning] "+format, params...)
}

func Fatal(message string) {
	if logLevel < 3 {
		return
	}

	log.Printf("[%s] %s", logLevelStr, message)
}

func Fatalf(format string, params ...interface{}) {
	if logLevel < 1 {
		return
	}

	log.Printf("[Fatal] "+format, params...)
}

func Write(message string) {
	log.Printf("[%s] %s", logLevelStr, message)
}
