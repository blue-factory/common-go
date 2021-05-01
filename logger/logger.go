package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger ...
type Logger struct {
	logger *logrus.Logger
}

// NewWithPath ...
func NewWithPath(path string) (*Logger, error) {
	// define logrus intance
	log := logrus.New()

	// load log output file
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	// configure log
	mw := io.MultiWriter(os.Stdout, file)

	log.SetOutput(mw)
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	return &Logger{
		logger: log,
	}, nil
}

// New ...
func New() *Logger {
	// define logrus intance
	log := logrus.New()

	// configure log
	mw := io.MultiWriter(os.Stdout)

	log.SetOutput(mw)
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	return &Logger{
		logger: log,
	}
}

// Warnf ...
func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args...)
}

// InfoWithValues ...
func (l *Logger) InfoWithValues(msg string, meta map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(meta)).Info(msg)
}

// ErrWithValues ..
func (l *Logger) ErrWithValues(err error) {
	l.logger.WithFields(logrus.Fields{
		"error_cause": err,
	}).Error(err)
}

// Info ...
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args)
}

// Infof ...
func (l *Logger) Infof(msg string, args ...interface{}) {
	l.logger.Infof(msg, args)
}

// GetLogger ...
func (l *Logger) GetLogger() *logrus.Logger {
	return l.logger
}
