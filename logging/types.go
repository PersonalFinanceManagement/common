package logging

import "go.uber.org/zap"

// Logger is an interface that defines the methods for logging.
type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type logger struct {
	*zap.SugaredLogger
}

type LoggerConfig struct {
	// sets the logging level i.e. debug, info, warn, error, fatl
	Level string
	// Incase of production it logs with json output format
	Production bool
	// development mode enables features such as stracktraces, caller info etc.
	Development bool
	// defines the log output destinations like to a file or stdout or so
	OutputPaths []string
	// defines the log error output destinations like stderr or to a file
	ErrorOutputPaths []string
	//Encoding of the log, json or console
	Encoding string
	// adds constant fields to all log entries, like request id or so
	InitialFields map[string]interface{}
	//Includes caller info; like line number + file name
	EnableCaller bool
}
