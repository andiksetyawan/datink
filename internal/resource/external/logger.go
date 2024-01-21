package external

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"datink/config"
)

func NewLogger(config *config.Config) (*zap.Logger, error) {
	loggerConfig := zap.NewDevelopmentConfig()

	loggerConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
		TimeKey:    "time",
		EncodeTime: zapcore.RFC3339TimeEncoder,
	}

	level := zapcore.InfoLevel

	if l, err := zapcore.ParseLevel(config.AppLogLevel); err == nil {
		level = l
	}

	log.Println(level.String())

	loggerConfig.Level = zap.NewAtomicLevelAt(level)

	logger, err := loggerConfig.Build(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.Fields(zap.String("app_name", config.AppName)))

	if err != nil {
		return nil, err
	}

	return logger, nil
}
