package grpc

import (
	"fmt"
	"os"
	"strconv"

	"github.com/betchi/zapper"
)

// Config is settings of jaeger logger
type Config struct {
	Noop bool
}

// Logger is a logger for gorp
type Logger struct {
	noop bool
	v    int
}

var (
	logger *Logger
)

// InitGlobalLogger initialize global logger for jaeger
func InitGlobalLogger(config *Config) {
	logger = &Logger{
		noop: config.Noop,
	}

	vLevel := os.Getenv("GRPC_GO_LOG_VERBOSITY_LEVEL")
	if vl, err := strconv.Atoi(vLevel); err == nil {
		logger.v = vl
	}
}

// GlobalLogger retrieve global logger for jaeger
func GlobalLogger() *Logger {
	return logger
}

// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Info(args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Info(fmt.Sprint(args...))
}

// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Infoln(args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Info(fmt.Sprintln(args...))
}

// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Info(fmt.Sprintf(format, args...))
}

// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Warning(args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Warn(fmt.Sprint(args...))
}

// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Warningln(args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Warn(fmt.Sprintln(args...))
}

// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warningf(format string, args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Warn(fmt.Sprint(args...))
}

// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Error(args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Error(fmt.Sprint(args...))
}

// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Errorln(args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Error(fmt.Sprintln(args...))
}

// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Error(fmt.Sprintf(format, args...))
}

// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *Logger) Fatal(args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Fatal(fmt.Sprint(args...))
}

// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *Logger) Fatalln(args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Fatal(fmt.Sprintln(args...))
}

// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	if l == nil || l.noop {
		return
	}

	zapper.GlobalLogger().Fatal(fmt.Sprintf(format, args...))
}

// V reports whether verbosity level l is at least the requested verbose level.
func (l *Logger) V(v int) bool {
	return v <= l.v
}
