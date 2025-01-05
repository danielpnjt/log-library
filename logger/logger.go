package logger

import (
	"os"
	"time"

	"github.com/danielpnjt/log-library/constant"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// InitLogger menginisialisasi logger global
func InitLogger(config *Config) error {
	logFileName := "logs/" + time.Now().Format(constant.LogDateFormat) + ".log"

	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	log = logrus.New()
	log.Out = file
	log.SetFormatter(&logrus.JSONFormatter{})

	return nil
}

// Info untuk mencatat log level info
func Info(message string, fields logrus.Fields) {
	log.WithFields(fields).Info(message)
}

// Warn untuk mencatat log level warning
func Warn(message string, fields logrus.Fields) {
	log.WithFields(fields).Warn(message)
}

// Error untuk mencatat log level error
func Error(message string, err error, fields logrus.Fields) {
	log.WithFields(fields).WithError(err).Error(message)
}
