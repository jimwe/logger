package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLoggerStartWithWriter(t *testing.T) {
	var resultBuffer bytes.Buffer
	log := Logger{}
	err := log.Start(ERROR, &resultBuffer, "")

	if err != nil {
		t.Errorf("Failed to start logger with writer")
	}
}

func TestLoggerStartWithFile(t *testing.T) {
	path := "slack_msg_svc.log"
	log := Logger{}
	err := log.Start(ERROR, nil, path)

	if err != nil {
		t.Error("Failed to start logger with file: ", path)
	}
}

func TestLoggerStartWithStdOut(t *testing.T) {
	log := Logger{}
	err := log.Start(DEBUG, nil, "")
	log.Debug("This message should appear on StdOut")

	if err != nil {
		t.Error("Failed to start logger using StdOut")
	}
}

func TestLoggerDebugLevel(t *testing.T) {
	var resultBuffer bytes.Buffer
	log := Logger{}
	log.Start(DEBUG, &resultBuffer, "")

	msg := "Here is a debug message"
	log.Debug(msg)
	if !strings.Contains(resultBuffer.String(), msg) {
		t.Errorf("Debug message should be logged when log level is DEBUG.")
	}
	msg = "Here is an Info message"
	log.Info(msg)
	if !strings.Contains(resultBuffer.String(), msg) {
		t.Errorf("Informational message should be logged when log level is DEBUG")
	}
	msg = "Here is a warning message"
	log.Warning(msg)
	if !strings.Contains(resultBuffer.String(), msg) {
		t.Errorf("Warning message should be logged when log level is DEBUG")
	}
	msg = "Here is an error message"
	log.Error(msg)
	if !strings.Contains(resultBuffer.String(), msg) {
		t.Errorf("Error message should be logged when log level is DEBUG")
	}
}

func TestLoggerErrorLevel(t *testing.T) {
	var resultBuffer bytes.Buffer
	log := Logger{}
	log.Start(ERROR, &resultBuffer, "")

	msg := "Here is a debug message"
	log.Debug(msg)
	if strings.Contains(resultBuffer.String(), msg) {
		t.Errorf("Debug message should not be logged when log level is ERROR.")
	}
	msg = "Here is an Info message"
	log.Info(msg)
	if strings.Contains(resultBuffer.String(), msg) {
		t.Errorf("Informational message should not be logged when log level is ERROR.")
	}
	msg = "Here is a warning message"
	log.Warning(msg)
	if strings.Contains(resultBuffer.String(), msg) {
		t.Errorf("Warning message should not be logged when log level is ERROR.")
	}
	msg = "Here is an error message"
	log.Error(msg)
	if !strings.Contains(resultBuffer.String(), msg) {
		t.Errorf("Error message should be logged when log level is DEBUG")
	}
}

func TestSample(t *testing.T) {
	log := Logger{}
	log.Start(INFO, nil, "myproject.log")
	log.Warning("Danger Will Smith")
	log.Stop()
}
