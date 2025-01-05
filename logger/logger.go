package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// InitLogger menginisialisasi logger global
func InitLogger(config *Config) error {
	var file *os.File
	var err error
	log, file, err = SetupLogger(config)
	if err != nil {
		return err
	}
	defer file.Close()
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
