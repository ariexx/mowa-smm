package services

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type LoggerService interface {
	PrintStdout(ctx context.Context, level zapcore.Level, message string, fields ...zapcore.Field)
}

type loggerService struct {
	log *zap.Logger
}

func NewLoggerService() LoggerService {
	//check env
	if os.Getenv("APP_ENV") == "production" {
		logger, _ := zap.NewProduction()
		return &loggerService{
			log: logger,
		}
	}

	return &loggerService{
		log: zap.Must(zap.NewDevelopment()),
	}
}

func (l *loggerService) PrintStdout(ctx context.Context, level zapcore.Level, message string, fields ...zapcore.Field) {
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

}
