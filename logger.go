package logger

import (
	"fmt"
	"io"
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
	prefixInfo  = "INFO: "
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

// Init loggers for each log levels
func New(verbose bool) *Logger {
	flags := log.Ldate | log.Ltime | log.Lmicroseconds
	l := Logger{
		infoLog:  log.New(os.Stdout, prefixInfo, flags),
		errorLog: log.New(os.Stderr, prefixError, flags),
		fatalLog: log.New(os.Stderr, prefixFatal, flags),
		verbose:  verbose,
	}
	return &l
}

// Sets the output destination for the loggers
// By default destination is stdout, you can change that
// with this function
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.infoLog.SetOutput(w)
	l.errorLog.SetOutput(w)
	l.fatalLog.SetOutput(w)
}

// enable the verbose mode
func (l *Logger) SetVerbose(verbose bool) {
	l.verbose = verbose
}

// Return Error Logger
func (l *Logger) ErrorLogger() *log.Logger {
	return l.errorLog
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
