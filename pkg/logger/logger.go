package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLog  *log.Logger
	debugLog *log.Logger
	errorLog *log.Logger
	traceLog *log.Logger
)

// Init initializes the loggers (call this in main)
func Init() {
	infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	debugLog = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	traceLog = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Llongfile)
}

// Info logs an informational message
func Info(args ...interface{}) {
	infoLog.Println(args...)
}

// Debug logs a debug message
func Debug(args ...interface{}) {
	debugLog.Println(args...)
}

// Error logs an error message with file:line
func Error(args ...interface{}) {
	if err := errorLog.Output(2, fmt.Sprint(args...)); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write error log: %v\n", err)
	}
}

// Trace logs a trace message with full file path
func Trace(args ...interface{}) {
	if err := traceLog.Output(2, fmt.Sprint(args...)); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write trace log: %v\n", err)
	}
}

// Infof logs a formatted informational message
func Infof(format string, args ...interface{}) {
	infoLog.Printf(format, args...)
}

// Debugf logs a formatted debug message
func Debugf(format string, args ...interface{}) {
	debugLog.Printf(format, args...)
}

// Errorf logs a formatted error message with file:line
func Errorf(format string, args ...interface{}) {
	if err := errorLog.Output(2, fmt.Sprintf(format, args...)); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write error log: %v\n", err)
	}
}

// Tracef logs a formatted trace message with full file path
func Tracef(format string, args ...interface{}) {
	if err := traceLog.Output(2, fmt.Sprintf(format, args...)); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write trace log: %v\n", err)
	}
}
