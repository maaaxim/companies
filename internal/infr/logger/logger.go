package logger

import (
	errors "github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func InitLogger(cfg Config) (Logger, error) {

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	level, err := cfg.LogLevel.toZapLevel()
	if err != nil {
		return Logger{}, err
	}
	zapConfig.Encoding = cfg.Encoding
	zapConfig.DisableCaller = true
	zapConfig.DisableStacktrace = true
	zapConfig.Level = level
	zapConfig.OutputPaths = []string{"stdout"}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		return Logger{}, errors.Wrap(err, "build zap logger error")
	}

	return Logger{Logger: zapLogger}, nil
}

func (l Logger) Errors(errs []error) {
	var str string
	for i, err := range errs {
		str += err.Error()
		if len(errs) > 0 && i != len(errs) {
			str += ","
		}
	}
	if len(str) > 0 {
		l.Error(str)
	}
}
