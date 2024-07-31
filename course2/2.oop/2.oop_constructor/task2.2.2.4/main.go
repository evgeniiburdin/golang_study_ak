package main

import (
	"fmt"
	"io"
	"os"
)

// Logger interface defines the method for logging
type Logger interface {
	Log(message string)
}

// FileLogger struct represents a logger that writes to a file
type FileLogger struct {
	file *os.File
}

// Log writes a message to the file
func (f FileLogger) Log(message string) {
	fmt.Fprintln(f.file, message)
}

// ConsoleLogger struct represents a logger that writes to the console
type ConsoleLogger struct {
	out io.Writer
}

// Log writes a message to the console
func (c ConsoleLogger) Log(message string) {
	fmt.Fprintln(c.out, message)
}

// LogSystem struct represents the logging system
type LogSystem struct {
	logger Logger
}

// LogOption defines the type for functional options
type LogOption func(*LogSystem)

// WithLogger sets the logger for the LogSystem
func WithLogger(logger Logger) LogOption {
	return func(ls *LogSystem) {
		ls.logger = logger
	}
}

// NewLogSystem creates a new LogSystem with the provided options
func NewLogSystem(opts ...LogOption) *LogSystem {
	logSystem := &LogSystem{}
	for _, opt := range opts {
		opt(logSystem)
	}
	return logSystem
}

// Log logs a message using the configured logger
func (ls *LogSystem) Log(message string) {
	if ls.logger != nil {
		ls.logger.Log(message)
	}
}
