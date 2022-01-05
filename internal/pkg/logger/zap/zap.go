package zap

import (
	"io"

	"github.com/persiangophers/userland/internal/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type log struct {
	zap *zap.SugaredLogger
}

func New(writer io.Writer, level zapcore.Level) logger.Logger {
	ws := zapcore.AddSync(writer)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	enc := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(enc, ws, level)

	z := zap.New(core)
	sugarLogger := z.Sugar()
	return &log{sugarLogger}
}

func (l *log) Sync() {
	l.zap.Sync()
}

func (l *log) Error(msg string, kv ...interface{}) {
	l.zap.Errorw(msg, kv...)
}

func (l *log) Warn(msg string, kv ...interface{}) {
	l.zap.Warnw(msg, kv...)
}

func (l *log) Info(msg string, kv ...interface{}) {
	l.zap.Infow(msg, kv...)
}
