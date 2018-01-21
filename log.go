package logger

// Usage: Create a logger that writes all ERROR, WARNING and INFO messages to a file.
// DEBUG messages will not be included.
//		log := Slogger{}
//		log.Start(INFO, nil, "myproject.log")
//		log.Warning("Danger Will Smith")
//		log.Stop()

import (
	"io"
	"log"
	"os"
)

const (
	ERROR   = 1
	WARNING = 2
	INFO    = 3
	DEBUG   = 4
)

type Slogger struct {
	level int
	file  *os.File
	*log.Logger
	prefix string
	flags  int
}

// Level is required and is one of the constants above. If file name is provided, we'll create that
// file and use it's writer. If writer is provided instead of fileName, we'll write to it. Finally
// if writer is nil and fileName is "", we write to standard out.
func (logger *Slogger) Start(level int, writer io.Writer, fileName string) (err error) {
	logger.level = level
	logger.flags = log.Ldate | log.Ltime
	logger.Logger = log.New(os.Stdout, logger.prefix, logger.flags)
	if len(fileName) > 0 {
		logger.file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			logger.Logger.Println(err)
		} else {
			logger.Logger.SetOutput(logger.file)
		}
	} else if writer != nil {
		logger.Logger.SetOutput(writer)
	}
	return
}

func (logger *Slogger) Stop() (err error) {
	if logger.file != nil {
		err = logger.file.Close()
	}
	return
}

func (logger *Slogger) Error(s string) {
	if logger.level >= ERROR {
		logger.Logger.Println(s)
	}
}

func (logger *Slogger) Warning(s string) {
	if logger.level >= WARNING {
		logger.Logger.Println(s)
	}
}

func (logger *Slogger) Info(s string) {
	if logger.level >= INFO {
		logger.Logger.Println(s)
	}
}

func (logger *Slogger) Debug(s string) {
	if logger.level >= DEBUG {
		logger.Logger.Println(s)
	}
}
