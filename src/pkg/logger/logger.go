package logger

import (
	"context"
	"log"
	"os"

	"github.com/go-chi/chi/v5/middleware"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

type Config struct {
	ServiceName string
	Debug       bool
	GraylogPath string
}

var attachedLoggerKey = &ctxKey{}

var globalLogger *zap.SugaredLogger

func fromContext(ctx context.Context) *zap.SugaredLogger {
	var result = globalLogger
	if attachedLogger, ok := ctx.Value(attachedLoggerKey).(*zap.SugaredLogger); ok {
		result = attachedLogger
	}

	var requestID string
	if reqID := ctx.Value(middleware.RequestIDKey); reqID != nil {
		requestID = reqID.(string)
		if requestID != "" {
			result = result.With("trace-id", requestID)
		}
	}

	var traceID string
	if reqID := ctx.Value("trace-id"); reqID != nil {
		traceID = reqID.(string)
		if traceID != "" {
			result = result.With("trace-id", traceID)
		}
	}

	var consumer string
	if consumerName := ctx.Value("consumer"); consumerName != nil {
		consumer = consumerName.(string)
		if traceID != "" {
			result = result.With("consumer", consumer)
		}
	}

	return result
}

func Error(message string, kvs ...interface{}) {
	globalLogger.Errorw(message, kvs...)
}

func Info(message string, kvs ...interface{}) {
	globalLogger.Infow(message, kvs...)
}

func Debug(message string, kvs ...interface{}) {
	globalLogger.Debugw(message, kvs...)
}

func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Errorw(message, kvs...)
}

func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Warnw(message, kvs...)
}

func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Infow(message, kvs...)
}

func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Debugw(message, kvs...)
}

func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Fatalw(message, kvs...)
}

func AttachLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, attachedLoggerKey, logger)
}

func InitLogger(cfg Config) (syncFn func()) {
	loggingLevel := zap.InfoLevel
	if cfg.Debug {
		loggingLevel = zap.DebugLevel
	}

	encCfg := zap.NewProductionEncoderConfig()
	encCfg.TimeKey = "@ts"
	encCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encCfg),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	notSugaredLogger := zap.New(consoleCore)

	sugaredLogger := notSugaredLogger.Sugar()
	globalLogger = sugaredLogger.With(
		"service", cfg.ServiceName,
	)

	return func() {
		_ = notSugaredLogger.Sync()
	}
}

func init() {
	notSugaredLogger, err := zap.NewProduction()
	if err != nil {
		log.Panic(err)
	}

	globalLogger = notSugaredLogger.Sugar()
}
