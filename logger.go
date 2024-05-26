package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type Level int

const (
	INFO Level = iota
	WARNING
	ERROR
	FATAL
)

const (
	prefixInfo    = "INFO: "
	prefixWarning = "WARNING: "
	prefixError   = "ERROR: "
	prefixFatal   = "FATAL: "
)

type LogEntry struct {
	Timestamp time.Time
	Level     Level
	Message   string
}

type Logger struct {
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	fatalLog   *log.Logger
	mu         sync.Mutex
	verbose    bool
	channel    chan LogEntry
}

// Init loggers for each log levels
func New(verbose bool) *Logger {
	flags := log.Ldate | log.Ltime | log.Lmicroseconds
	l := Logger{
		infoLog:    log.New(os.Stdout, prefixInfo, flags),
		warningLog: log.New(os.Stdout, prefixWarning, flags),
		errorLog:   log.New(os.Stderr, prefixError, flags),
		fatalLog:   log.New(os.Stderr, prefixFatal, flags),
		verbose:    verbose,
	}
	return &l
}

// Sets the output to a channel
func (l *Logger) SetOutputChannel(c chan LogEntry) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.channel = c
}

// Sets the output destination for the loggers
// By default destination is stdout, you can change that
// with this function
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.infoLog.SetOutput(w)
	l.warningLog.SetOutput(w)
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
	case WARNING:
		l.warningLog.Output(depth, msg)
	case ERROR:
		l.errorLog.Output(depth, msg)
	case FATAL:
		l.fatalLog.Output(depth, msg)
	default:
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.verbose {
		if l.channel != nil {
			l.channel <- LogEntry{Timestamp: time.Now(), Level: INFO, Message: fmt.Sprintf(format, v...)}
		} else {
			l.output(INFO, fmt.Sprintf(format, v...))
		}
	}
}

func (l *Logger) Warning(format string, v ...interface{}) {
	if l.verbose {
		if l.channel != nil {
			l.channel <- LogEntry{Timestamp: time.Now(), Level: WARNING, Message: fmt.Sprintf(format, v...)}
		} else {
			l.output(WARNING, fmt.Sprintf(format, v...))
		}
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l.channel != nil {
		l.channel <- LogEntry{Timestamp: time.Now(), Level: ERROR, Message: fmt.Sprintf(format, v...)}
	} else {
		l.output(ERROR, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.channel != nil {
		l.channel <- LogEntry{Timestamp: time.Now(), Level: FATAL, Message: fmt.Sprint(v...)}
	} else {
		l.output(FATAL, fmt.Sprint(v...))
		os.Exit(1)
	}
}
