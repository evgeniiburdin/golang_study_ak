package main

import (
	"bytes"
	"os"
	"testing"
)

// Test FileLogger
func TestFileLogger_Log(t *testing.T) {
	file, err := os.CreateTemp("", "log.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	logSystem.Log("Hello, world!")

	file.Seek(0, 0)
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	loggedMessage := buf.String()

	expectedMessage := "Hello, world!\n"
	if loggedMessage != expectedMessage {
		t.Errorf("Expected %q, got %q", expectedMessage, loggedMessage)
	}
}

// Test ConsoleLogger
func TestConsoleLogger_Log(t *testing.T) {
	var buf bytes.Buffer
	consoleLogger := ConsoleLogger{out: &buf}
	logSystem := NewLogSystem(WithLogger(consoleLogger))

	logSystem.Log("Hello, console!")

	expectedMessage := "Hello, console!\n"
	if buf.String() != expectedMessage {
		t.Errorf("Expected %q, got %q", expectedMessage, buf.String())
	}
}

// Test NewLogSystem with default options
func TestNewLogSystem_Default(t *testing.T) {
	logSystem := NewLogSystem()
	if logSystem.logger != nil {
		t.Error("Expected default logger to be nil")
	}
}

// Test WithLogger option
func TestWithLogger(t *testing.T) {
	file, err := os.CreateTemp("", "log.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	if logSystem.logger != fileLogger {
		t.Error("Expected logger to be set by WithLogger option")
	}
}
