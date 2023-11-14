package designpattern

import "fmt"

type Logger interface {
	Log(message string)
}

type FileLogger struct{}

func (l *FileLogger) Log(message string) {
	fmt.Println("Logging to file:", message)
}

type DatabaseLogger struct{}

func (l *DatabaseLogger) Log(message string) {
	fmt.Println("Logging to database:", message)
}

type LoggerFactory interface {
	CreateLogger() Logger
}

type FileLoggerFactory struct{}

func (f *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

type DatabaseLoggerFactory struct{}

func (f *DatabaseLoggerFactory) CreateLogger() Logger {
	return &DatabaseLogger{}
}