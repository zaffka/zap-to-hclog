package log

import (
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
	"go.uber.org/zap"
)

// Adapt simplifies wrapping zap instance to Logger interface.
func Adapt(zap *zap.Logger) Logger {
	return Logger{Zap: zap}
}

type Level = hclog.Level
type Logger = SimplifiedLogger

// Log logs messages with four simplified levels - Debug,Warn,Error and Info as a default.
func (l Logger) Log(lvl Level, msg string, args ...interface{}) {
	switch lvl {
	case hclog.Debug:
		l.Debug(msg, args...)
	case hclog.Warn:
		l.Warn(msg, args...)
	case hclog.Error:
		l.Error(msg, args...)
	case hclog.DefaultLevel, hclog.Info, hclog.NoLevel, hclog.Off, hclog.Trace:
		l.Info(msg, args...)
	}
}

// Trace will log an info-level message in Zap.
func (l Logger) Trace(msg string, args ...interface{}) {
	l.Zap.Info(msg, convertToZapAny(args...)...)
}

// With returns a logger with always-presented key-value pairs.
func (l Logger) With(args ...interface{}) hclog.Logger {
	return &Logger{Zap: l.Zap.With(convertToZapAny(args...)...)}
}

// Named returns a logger with the specific name.
// The name string will always be presented in a log messages.
func (l Logger) Named(name string) hclog.Logger {
	return &Logger{Zap: l.Zap.Named(name), name: name}
}

// Name returns a logger's name (if presented).
func (l Logger) Name() string { return l.name }

// ResetNamed has the same implementation as Named.
func (l Logger) ResetNamed(name string) hclog.Logger {
	return &Logger{Zap: l.Zap.Named(name), name: name}
}

// StandardWriter returns os.Stderr as io.Writer.
func (l Logger) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	return hclog.DefaultOutput
}

// StandardLogger returns standard logger with os.Stderr as a writer.
func (l Logger) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	return log.New(l.StandardWriter(opts), "", log.LstdFlags)
}

// IsTrace has no implementation.
func (l Logger) IsTrace() bool { return false }

// IsDebug has no implementation.
func (l Logger) IsDebug() bool { return false }

// IsInfo has no implementation.
func (l Logger) IsInfo() bool { return false }

// IsWarn has no implementation.
func (l Logger) IsWarn() bool { return false }

// IsError has no implementation.
func (l Logger) IsError() bool { return false }

// ImpliedArgs has no implementation.
func (l Logger) ImpliedArgs() []interface{} { return nil }

// SetLevel has no implementation.
func (l Logger) SetLevel(lvl Level) {
}
