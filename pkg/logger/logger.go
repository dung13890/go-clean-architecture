package logger

import (
	// "fmt"
	"log"
	"os"
	// "runtime"
)

var (
	info  *log.Logger
	debug *log.Logger
	err   *log.Logger
)

// InitLogger init
func InitLogger() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	err = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info Log
func Info(v ...any) *log.Logger {
	return info
}

// Debug Log
func Debug(v ...any) *log.Logger {
	return debug
}

// Error Log
func Error(v ...any) *log.Logger {
	return err
}
