package logging

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         "console",
	}
}

func DefaultLoggerConfig() *LoggerConfig {
	developmentLoggerConfig := NewLoggerConfig()
	developmentLoggerConfig.Production = false
	developmentLoggerConfig.Development = true
	developmentLoggerConfig.EnableCaller = true
	developmentLoggerConfig.InitialFields = map[string]interface{}{
		"env":     "development",
		"caution": "true",
	}
	return developmentLoggerConfig
}
