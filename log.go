package signaling

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

var logger *log.Logger

type Logger struct{}

func (l Logger) SetUp(logFilePath, level string) {
	logger = log.New()
	logger.Formatter = new(log.JSONFormatter)
	if lv, err := log.ParseLevel(level); err == nil {
		logger.Level = lv
	}
	if fs, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err == nil {
		logger.Out = fs
	}
}

func (l Logger) Debug(message string) {
	logger.Debug(message)
}
func (l Logger) Fatal(message string) {
	logger.Fatal(message)
}
func (l Logger) Info(message string) {
	logger.Info(message)
}
func (l Logger) Warn(message string) {
	logger.Warn(message)
}
func (l Logger) Error(message string) {
	logger.Error(message)
}
func (l Logger) Panic(message string) {
	logger.Panic(message)
}
