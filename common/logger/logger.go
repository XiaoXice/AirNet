package logger

type Logger interface {
	Handler(msg string) bool
}
