package wrapper

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/hashicorp/go-hclog"
)

type Nop struct{}

func (l *Nop) Log(level Level, msg string, args ...interface{}) {}

func (l *Nop) Trace(msg string, args ...interface{}) {}

func (l *Nop) Debug(msg string, args ...interface{}) {}

func (l *Nop) Info(msg string, args ...interface{}) {}

func (l *Nop) Warn(msg string, args ...interface{}) {}

func (l *Nop) Error(msg string, args ...interface{}) {}

func (l *Nop) IsTrace() bool { return false }

func (l *Nop) IsDebug() bool { return false }

func (l *Nop) IsInfo() bool { return false }

func (l *Nop) IsWarn() bool { return false }

func (l *Nop) IsError() bool { return false }

func (l *Nop) ImpliedArgs() []interface{} { return []interface{}{} }

func (l *Nop) With(args ...interface{}) hclog.Logger { return l }

func (l *Nop) Name() string { return "" }

func (l *Nop) Named(name string) hclog.Logger { return l }

func (l *Nop) ResetNamed(name string) hclog.Logger { return l }

func (l *Nop) SetLevel(level Level) {}

func (l *Nop) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	return log.New(l.StandardWriter(opts), "", log.LstdFlags)
}

func (l *Nop) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	return ioutil.Discard
}
