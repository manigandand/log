// Package log is simple extended version of standard 'log' package
// based on logLevel
package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// logLevel is a severity level at which logger works
type logLevel int

// InfoLevel const value 0
// WarningLevel const value 1
// ErrorLevel const value 2
// FatalLevel const value 3
// PanicLevel const value 4
const (
	InfoLevel logLevel = iota
	WarningLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	fatalLogger   *log.Logger
	panicLogger   *log.Logger
)

// Levels map ...
var Levels = map[string]logLevel{
	"INFO":    InfoLevel,
	"WARNING": WarningLevel,
	"ERROR":   ErrorLevel,
	"FATAL":   FatalLevel,
	"PANIC":   PanicLevel,
}

func init() {
	Init(InfoLevel, nil)
}

// Init initiates the logger
func Init(lev logLevel, multiHandler io.Writer) {
	infoHandler := ioutil.Discard
	warningHandler := ioutil.Discard
	errorHandler := ioutil.Discard
	fatalHandler := ioutil.Discard
	panicHandler := ioutil.Discard

	switch lev {
	case InfoLevel:
		infoHandler = os.Stdout
		warningHandler = os.Stdout
		errorHandler = os.Stderr
		fatalHandler = os.Stderr
		panicHandler = os.Stdout
	case WarningLevel:
		warningHandler = os.Stdout
		errorHandler = os.Stderr
		fatalHandler = os.Stderr
		panicHandler = os.Stdout
	case ErrorLevel:
		errorHandler = os.Stderr
		fatalHandler = os.Stderr
		panicHandler = os.Stdout
	case FatalLevel:
		fatalHandler = os.Stderr
		panicHandler = os.Stdout
	case PanicLevel:
		panicHandler = os.Stdout
	default:
		log.Fatal("log: Invalid log level should be (0-4)")
	}

	if multiHandler != nil {
		if infoHandler == os.Stdout {
			infoHandler = io.MultiWriter(infoHandler, multiHandler)
		}
		if warningHandler == os.Stdout {
			warningHandler = io.MultiWriter(warningHandler, multiHandler)
		}
		if errorHandler == os.Stderr {
			errorHandler = io.MultiWriter(errorHandler, multiHandler)
		}
		if fatalHandler == os.Stderr {
			fatalHandler = io.MultiWriter(fatalHandler, multiHandler)
		}
		if panicHandler == os.Stderr {
			panicHandler = io.MultiWriter(panicHandler, multiHandler)
		}
	}

	infoLogger = log.New(infoHandler, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(warningHandler, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errorHandler, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	fatalLogger = log.New(fatalHandler, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	panicLogger = log.New(panicHandler, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Wrapper for changing logger's output writer anytime

// InfoSetOutput updates the info output writer
func InfoSetOutput(w io.Writer) {
	infoLogger.SetOutput(w)
}

// WarningSetOutput updates the warning output writer
func WarningSetOutput(w io.Writer) {
	warningLogger.SetOutput(w)
}

// ErrorSetOutput updates the error output writer
func ErrorSetOutput(w io.Writer) {
	errorLogger.SetOutput(w)
}

// FatalSetOutput updates the fatal output writer
func FatalSetOutput(w io.Writer) {
	fatalLogger.SetOutput(w)
}

// PanicSetOutput updates the Panic output writer
func PanicSetOutput(w io.Writer) {
	panicLogger.SetOutput(w)
}
