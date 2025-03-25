package logger

import (
	"context"
	"go.uber.org/zap"
)

const (
	Key = "logger"
)

type Logger struct {
	l *zap.Logger
}

func New(ctx context.Context) (context.Context, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, Key, &Logger{logger})
	return ctx, nil
}

func GetLoggerFromCtx(ctx context.Context) *Logger {
	return ctx.Value(Key).(*Logger)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)

}
func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	l.l.Fatal(msg, fields...)

}
