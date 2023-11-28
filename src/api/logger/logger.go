package logger

import (
	"os"

	"github.com/flexuxs/clubHubApp/src/api/config"
	"github.com/sirupsen/logrus"
)

func NewLogrus(l *logrus.Logger, config *config.Configuration) *logrus.Logger {
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.InfoLevel)

	return l
}
