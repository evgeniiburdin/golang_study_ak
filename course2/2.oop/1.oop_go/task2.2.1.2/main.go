package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logger interface defines the method for logging messages
type Logger interface {
	Log(message string) error
}

// ConsoleLogger struct represents a logger that logs to the console
type ConsoleLogger struct {
	Writer io.Writer
}

// Log method for ConsoleLogger
func (c *ConsoleLogger) Log(message string) error {
	if c.Writer == nil {
		return fmt.Errorf("console writer is nil")
	}
	_, err := fmt.Fprintln(c.Writer, message)
	return err
}

// FileLogger struct represents a logger that logs to a file
type FileLogger struct {
	File *os.File
}

// Log method for FileLogger
func (f *FileLogger) Log(message string) error {
	if f.File == nil {
		return fmt.Errorf("file is nil")
	}
	_, err := fmt.Fprintln(f.File, message)
	return err
}

// RemoteLogger struct represents a logger that logs to a remote server
type RemoteLogger struct {
	Address string
}

// Log method for RemoteLogger (dummy implementation)
func (r *RemoteLogger) Log(message string) error {
	// Here you would add code to send the log message to a remote server
	// For demonstration purposes, we will just print the message
	if r.Address == "" {
		return fmt.Errorf("remote address is empty")
	}
	fmt.Printf("Logging to remote server at %s: %s\n", r.Address, message)
	return nil
}

// LogAll function logs a message using multiple loggers
func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}

func main() {
	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer file.Close()
	fileLogger := &FileLogger{File: file}
	remoteLogger := &RemoteLogger{Address: "http://example.com/log"}

	loggers := []Logger{consoleLogger, fileLogger, remoteLogger}
	LogAll(loggers, "This is a test log message.")
}
