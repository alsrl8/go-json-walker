package utils

import (
	"log"
	"log/slog"
	"os"
	"sync"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

var (
	once     sync.Once
	goLogger *GoLogger
)

type GoLogger struct {
	*slog.Logger
}

func GetLogger() *GoLogger {
	once.Do(func() {
		file, err := os.Create("logs/app.log")
		if err != nil {
			log.Fatal(err)
		}
		h := slog.NewJSONHandler(
			file,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		)
		slogger := slog.New(h)
		goLogger = &GoLogger{
			slogger,
		}
	})
	return goLogger
}

func (gl *GoLogger) Info(msg string, args ...interface{}) {
	gl.Logger.Info(msg, args...)
}

func (gl *GoLogger) Debug(msg string, args ...interface{}) {
	gl.Logger.Debug(msg, args...)
}

func (gl *GoLogger) Warn(msg string, args ...interface{}) {
	gl.Logger.Warn(msg, args...)
}

func (gl *GoLogger) Error(msg string, args ...interface{}) {
	gl.Logger.Error(msg, args...)
}
