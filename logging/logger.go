package logging

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(loggerConfig *LoggerConfig) (Logger, error) {
	if loggerConfig == nil {
		loggerConfig = DefaultLoggerConfig()
	}
	var zapConfig zap.Config
	if loggerConfig.Production {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}

	if len(loggerConfig.OutputPaths) > 0 {
		zapConfig.OutputPaths = loggerConfig.OutputPaths
	}
	if len(loggerConfig.ErrorOutputPaths) > 0 {
		zapConfig.ErrorOutputPaths = loggerConfig.ErrorOutputPaths
	}
	// set encoding
	if loggerConfig.Encoding != "" {
		zapConfig.Encoding = loggerConfig.Encoding
	}

	// Set log level
	if loggerConfig.Level != "" {
		level, err := zapcore.ParseLevel(strings.ToLower(loggerConfig.Level))
		if err != nil {
			return nil, err
		}
		zapConfig.Level = zap.NewAtomicLevelAt(level)
	}

	// intial fields
	if loggerConfig.InitialFields != nil {
		zapConfig.InitialFields = loggerConfig.InitialFields
	}
	zapConfig.Development = loggerConfig.Development
	zapConfig.DisableCaller = !loggerConfig.EnableCaller
	// Default
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLogger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	return &logger{zapLogger.Sugar()}, nil
}
