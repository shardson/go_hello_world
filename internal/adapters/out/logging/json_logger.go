package logging

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Package   string `json:"package"`
	Function  string `json:"function"`
	Level     string `json:"level"`
}

type JSONLogger struct {
	logger *log.Logger
}

func NewJSONLogger() *JSONLogger {
	return &JSONLogger{
		logger: log.New(os.Stdout, "", 0),
	}
}

func (l *JSONLogger) Info(pkg, function, message string) {
	entry := LogEntry{
		Timestamp: time.Now().Format("2006/01/02 15:04:05"),
		Message:   message,
		Package:   pkg,
		Function:  function,
		Level:     "INFO",
	}

	l.write(entry)
}

func (l *JSONLogger) Error(pkg, function, message string) {
	entry := LogEntry{
		Timestamp: time.Now().Format("2006/01/02 15:04:05"),
		Message:   message,
		Package:   pkg,
		Function:  function,
		Level:     "ERROR",
	}

	l.write(entry)
}

func (l *JSONLogger) write(entry LogEntry) {
	data, err := json.Marshal(entry)
	if err != nil {
		l.logger.Printf(`{"timestamp":"%s","message":"failed to marshal log entry","package":"logging","function":"write","level":"ERROR"}`,
			time.Now().Format("2006/01/02 15:04:05"))
		return
	}

	l.logger.Println(string(data))
}
