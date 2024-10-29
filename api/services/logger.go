package services

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type LoggerService interface {
	PrintStdout(ctx context.Context, level zapcore.Level, message string, fields ...zapcore.Field)
	handleError(ctx context.Context, err error)
}

type loggerService struct {
	log *zap.Logger
}

func NewLoggerService() LoggerService {
	return &loggerService{
		log: initializeLogger(),
	}
}

func initializeLogger() *zap.Logger {
	if isProduction() {
		logger, _ := zap.NewProduction()
		return logger
	}
	return zap.Must(zap.NewDevelopment())
}

func isProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}

func (l *loggerService) PrintStdout(ctx context.Context, level zapcore.Level, message string, fields ...zapcore.Field) {
	// Add context fields
	fields = append(fields, zapcore.Field{Key: "request_id", Type: zapcore.StringType, String: ctx.Value("requestid").(string)})

	switch level {
	case zapcore.DebugLevel:
		l.log.Debug(message, fields...)
	case zapcore.InfoLevel:

		l.log.Info(message, fields...)
	case zapcore.WarnLevel:
		l.log.Warn(message, fields...)
	case zapcore.ErrorLevel:
		l.log.Error(message, fields...)
	default:
		l.log.Info(message, fields...)
	}
}

func (l *loggerService) handleError(ctx context.Context, err error) {
	if err != nil {
		l.PrintStdout(ctx, zapcore.ErrorLevel, "error", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
	}
}
