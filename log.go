package logger
// Usage: Create a logger that writes all ERROR, WARNING and INFO messages to a file.
// DEBUG messages will not be included.
//		log := Logger{}
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

type Logger struct {
	level int
	file  *os.File
}

// Level is required and is one of the constants above. If file name is provided, we'll create that
// file and use it's writer. If writer is provided instead of fileName, we'll write to it. Finally
// if writer is nil and fileName is "", we write to standard out.
func (logger *Logger) Start(level int, writer io.Writer, fileName string) (err error) {
	logger.level = level
	if len(fileName) > 0 {
		logger.file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Println(err)
		}
		log.SetOutput(logger.file)
	} else if writer != nil {
		log.SetOutput(writer)
	} else {
		log.SetOutput(os.Stdout)
	}

	return
}

func (logger *Logger) Stop() (err error) {
	if logger.file != nil {
		err = logger.file.Close()
	}
	return
}

func (logger *Logger) Error(s string) {
	if logger.level >= ERROR {
		log.Println(s)
	}
}

func (logger *Logger) Warning(s string) {
	if logger.level >= WARNING {
		log.Println(s)
	}
}

func (logger *Logger) Info(s string) {
	if logger.level >= INFO {
		log.Println(s)
	}
}

func (logger *Logger) Debug(s string) {
	if logger.level >= DEBUG {
		log.Println(s)
	}
}
