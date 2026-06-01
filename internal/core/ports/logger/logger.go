package logger

type Logger interface {
	Info(pkg, function, message string)
	Error(pkg, function, message string)
}
