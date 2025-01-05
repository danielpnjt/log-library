package logger

import (
	"os"
	"time"

	"github.com/danielpnjt/log-library/constant"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var logFile *os.File

// InitLogger initializes the global logger
func InitLogger(config *Config) error {
	log = logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)

	// Create logs directory if it doesn't exist
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			return err
		}
	}

	// Create log file with current date
	logFileName := "logs/" + time.Now().Format(constant.LogDateFormat) + ".log"
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// Set output to file
	log.SetOutput(file)

	// Set log format to JSON
	log.SetFormatter(&logrus.JSONFormatter{})

	return nil
}

// CloseLogger closes the log file
func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}

// Info logs an info message
func Info(message string, fields logrus.Fields) {
	log.WithFields(fields).Info(message)
}

// Warn logs a warning message
func Warn(message string, fields logrus.Fields) {
	log.WithFields(fields).Warn(message)
}

// Error logs an error message
func Error(message string, err error, fields logrus.Fields) {
	log.WithFields(fields).WithError(err).Error(message)
}
