package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Config berisi konfigurasi untuk logger
type Config struct {
	LogLevel string
	LogFile  string
}

// SetupLogger inisialisasi logger dengan konfigurasi
func SetupLogger(config *Config) (*logrus.Logger, *os.File, error) {
	logger := logrus.New()

	// Set level log
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, nil, err
	}
	logger.SetLevel(level)

	// Set output log ke file
	file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, nil, err
	}
	logger.SetOutput(file)

	// Set formatter ke JSON
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger, file, nil
}
