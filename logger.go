package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Level int

const (
	INFO Level = iota
	ERROR
	FATAL
)

const (
	prefixInfo  = "INFO : "
	prefixError = "ERROR: "
	prefixFatal = "FATAL: "
)

type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	fatalLog *log.Logger
	mu       sync.Mutex
	verbose  bool
}

func New(verbose bool) *Logger {
	flags := log.Ldate | log.Ltime
	l := Logger{
		infoLog:  log.New(os.Stdout, prefixInfo, flags),
		errorLog: log.New(os.Stdout, prefixError, flags),
		fatalLog: log.New(os.Stdout, prefixFatal, flags),
		verbose:  verbose,
	}
	return &l
}
func (l *Logger) SetVerbose(verbose bool) {
	l.verbose = verbose
}

func (l *Logger) output(level Level, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	depth := 2
	switch level {
	case INFO:
		l.infoLog.Output(depth, msg)
	case ERROR:
		l.errorLog.Output(depth, msg)
	case FATAL:
		l.fatalLog.Output(depth, msg)
	default:
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.verbose {
		l.output(INFO, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.output(ERROR, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.output(FATAL, fmt.Sprint(v...))
	os.Exit(1)
}
