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

// Make sure we can call the embedded logger
func TestLoggerEmbeddedType(t *testing.T) {
	var resultBuffer bytes.Buffer
	msg := "Verify we can access the Println method on the embedded logger"
	log := Logger{}
	log.Start(ERROR, &resultBuffer, "")
	log.Println(msg)

	if !strings.Contains(resultBuffer.String(), msg) {
		t.Errorf("Failed to access the Println method on logger.")
	}
}

// Make sure we can call the embedded logger
func TestLoggerAddPrefix(t *testing.T) {
	var resultBuffer bytes.Buffer
	msg := "Verify we can add a prefix to the embedded logger"
	prefix := "Prefix: "
	log := Logger{}
	log.Start(ERROR, &resultBuffer, "")
	log.SetPrefix(prefix)
	log.Println(msg)

	if !strings.Contains(resultBuffer.String(), prefix) {
		t.Errorf("Failed to add a prefix to logger.")
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

func TestLoggerInterface(t *testing.T) {
	var resultBuffer bytes.Buffer
	log := Logger{}
	log.Start(DEBUG, &resultBuffer, "")
	callFuncWithLogger(log)

	s := resultBuffer.String()

	success := strings.Contains(s, "Debug") && strings.Contains(s, "Info")

	if !success {
		t.Errorf("Test of Logger interface failed")
	}
}

func callFuncWithLogger(log ILogger) {
	log.Debug("Debug")
	log.Info("Info")
	log.Warning("Warning")
	log.Error("Error")
}

func TestSample(t *testing.T) {
	log := Logger{}
	log.Start(INFO, nil, "myproject.log")
	log.Warning("Danger Will Smith")
	log.Stop()
}
