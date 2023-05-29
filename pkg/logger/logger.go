package logger

import (
	"log"
	"os"
)

var (
	info  *log.Logger
	debug *log.Logger
	err   *log.Logger
)

// InitLogger init
func InitLogger() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
	err = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
}

// Info Log
func Info(_ ...any) *log.Logger {
	return info
}

// Debug Log
func Debug(_ ...any) *log.Logger {
	return debug
}

// Error Log
func Error(_ ...any) *log.Logger {
	return err
}
