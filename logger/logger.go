package logger

import (
	"context"
	"log"
	"os"

	logrus "github.com/sirupsen/logrus"
)

var (
	outfile, _ = os.Create("logs.txt") // update path for your needs
	L          = log.New(outfile, "", 0)
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}

// SetLevel altera o level do logger
func SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)
}

// GetLevel recupera o level do logger
func GetLevel() logrus.Level {
	return logrus.GetLevel()
}

// Error exibe detalhes do erro
func Error(args ...interface{}) {
	logrus.Error(args...)
}

// ErrorContext exibe detalhes do erro com o contexto
func ErrorContext(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Error(args...)
}

// Info exibe detalhes do log info
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// InfoContext exibe detalhes do log info com o contexto
func InfoContext(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Info(args...)
}

// Debug exibe detalhes do log debug
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// DebugContext exibe detalhes do log debug com o contexto
func DebugContext(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Debug(args...)
}

// Trace exibe detalhes do log trace
func Trace(args ...interface{}) {
	logrus.Trace(args...)
}

// TraceContext exibe detalhes do log trace com o contexto
func TraceContext(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Trace(args...)
}

// Fatal exibe detalhes do erro
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
