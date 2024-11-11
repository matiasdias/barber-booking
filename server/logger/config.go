package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ConfigLogger() (*zap.Logger, error) {
	var cfg zap.Config

	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	cfg.Encoding = "json"

	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stack_trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	cfg.OutputPaths = []string{"stdout"}      // logs normais
	cfg.ErrorOutputPaths = []string{"stderr"} // logs de erro
	cfg.InitialFields = map[string]interface{}{
		"application": "shop-barber-api",
	}

	return cfg.Build()
}

// PanicRecovery handles recovered panics
func PanicRecovery(p interface{}) (err error) {
	zap.S().Error(
		"PANIC detected: ",
		p,
	)

	return
}
