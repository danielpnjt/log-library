package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Config berisi konfigurasi untuk logger
type Config struct {
	LogLevel string
	LogFile  string
}

const logDateFormat = "2006-01-02"

// SetupLogger inisialisasi logger dengan konfigurasi
func SetupLogger(config *Config) (*logrus.Logger, *os.File, error) {
	logger := logrus.New()

	// Set level log
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, nil, err
	}
	logger.SetLevel(level)

	// Create logs directory if it doesn't exist
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			return nil, nil, err
		}
	}

	// Create log file with current date
	logFileName := "logs/" + time.Now().Format(logDateFormat) + ".log"
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, nil, err
	}
	logger.SetOutput(file)

	// Set formatter ke JSON
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger, file, nil
}
