package closer

type Logger interface {
	Info(message string)
	Error(message string)
}

type noopLogger struct{}

func (*noopLogger) Info(string) {}

func (*noopLogger) Error(string) {}

func newNoopLogger() Logger {
	return &noopLogger{}
}
