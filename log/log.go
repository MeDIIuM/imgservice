package log

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"nir/config"
	"nir/di"
	"os"
	"path/filepath"
	"time"
)

func New(cfg *config.Config) (*logrus.Entry, error) {
	logPath := filepath.Join(".", "logs")

	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	_, err = os.Create(filepath.Join(logPath, fmt.Sprintf("%s.log", time.Now().String())))
	if err != nil {
		return nil, err
	}

	logger := logrus.New()

	logger.SetOutput(os.Stdin)
	logger.SetLevel(cfg.Log.Level)

	return logrus.NewEntry(logger), nil
}

func Error(ctx context.Context, args ...interface{}) {
	err := di.FromContext(ctx).Invoke(func(l *logrus.Entry) {
		l.Error(args...)
	})
	if err != nil {
		log.Println(err)
	}
}

func Info(ctx context.Context, args ...interface{}) {
	err := di.FromContext(ctx).Invoke(func(l *logrus.Entry) {
		l.Info(args...)
	})
	if err != nil {
		log.Println(err)
	}
}

func Warn(ctx context.Context, args ...interface{}) {
	err := di.FromContext(ctx).Invoke(func(l *logrus.Entry) {
		l.Warn(args...)
	})
	if err != nil {
		log.Println(err)
	}
}
