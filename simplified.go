package adapter

import "go.uber.org/zap"

type Simplified interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type SimplifiedLogger struct {
	Zap  *zap.Logger
	name string
}

func (l SimplifiedLogger) Debug(msg string, args ...interface{}) {
	l.Zap.Debug(msg, convertToZapAny(args...)...)
}
func (l SimplifiedLogger) Info(msg string, args ...interface{}) {
	l.Zap.Info(msg, convertToZapAny(args...)...)
}
func (l SimplifiedLogger) Warn(msg string, args ...interface{}) {
	l.Zap.Warn(msg, convertToZapAny(args...)...)
}
func (l SimplifiedLogger) Error(msg string, args ...interface{}) {
	l.Zap.Error(msg, convertToZapAny(args...)...)
}
