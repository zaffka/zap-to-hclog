package log

type SimplifiedNop struct{}

func (l SimplifiedNop) Debug(msg string, args ...interface{}) {}
func (l SimplifiedNop) Info(msg string, args ...interface{})  {}
func (l SimplifiedNop) Warn(msg string, args ...interface{})  {}
func (l SimplifiedNop) Error(msg string, args ...interface{}) {}
