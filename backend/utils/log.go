package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		ForceColors:      true,
		DisableTimestamp: false,
		TimestampFormat:  "2006-01-02 15:04:05",
	})

	Logger.SetLevel(logrus.InfoLevel)
}
