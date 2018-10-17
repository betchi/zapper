package elasticapm

import (
	"fmt"

	"github.com/betchi/zapper"
)

// Logger is a logger for elastic apm
type Logger struct{}

var logger = &Logger{}

// GlobalLogger retrieve global logger for elastic apm
func GlobalLogger() *Logger {
	return logger
}

// Debugf logs a message at DebugLevel for elastic apm
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l == nil {
		return
	}

	zapper.GlobalLogger().Debug(fmt.Sprintf(format, args...))
}

// Errorf logs a message at ErrorLevel for elastic apm
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l == nil {
		return
	}

	zapper.GlobalLogger().Error(fmt.Sprintf(format, args...))
}
