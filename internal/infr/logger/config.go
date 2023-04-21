package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level string

func (l Level) toZapLevel() (zap.AtomicLevel, error) {
	switch l {
	case DebugLevel:
		return zap.NewAtomicLevelAt(zapcore.DebugLevel), nil
	case ErrorLevel:
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel), nil
	case InfoLevel:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel), nil
	default:
		return zap.AtomicLevel{}, errors.Errorf("unknown log level: %v", l)
	}
}

const (
	DebugLevel = Level("debug")
	InfoLevel  = Level("info")
	ErrorLevel = Level("error")
)

type Config struct {
	Encoding string `envconfig:"ENCODING" default:"console"`
	LogLevel Level  `envconfig:"LEVEL" default:"debug"`
}
