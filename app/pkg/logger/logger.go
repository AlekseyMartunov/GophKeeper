package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	log *logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		log: logger,
	}
}

func (l *Logger) Info(text string) {
	l.log.Infoln(text)
}

func (l *Logger) Error(err error) {
	l.log.Errorln(err.Error())
}
