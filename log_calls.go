package log

import (
	"fmt"
	"os"
)

// Info writes into infologger as same as basic fmt.Print
func Info(v ...interface{}) {
	infoLogger.Output(2, fmt.Sprint(v...))
}

// Infoln writes into infologger as same as basic fmt.Println
func Infoln(v ...interface{}) {
	infoLogger.Output(2, fmt.Sprintln(v...))
}

// Infof writes into infologger as same as basic Outputf
func Infof(format string, v ...interface{}) {
	infoLogger.Output(2, fmt.Sprintf(format, v...))
}

// Warning writes warning messages into warninglogger as same as basic log.Output
func Warning(v ...interface{}) {
	warningLogger.Output(2, fmt.Sprint(v...))
}

// Warningln writes warning messages into warninglogger as same as basic log.Outputln
func Warningln(v ...interface{}) {
	warningLogger.Output(2, fmt.Sprintln(v...))
}

// Warningf writes warning messages into warninglogger as same as basic log.Outputf
func Warningf(format string, v ...interface{}) {
	warningLogger.Output(2, fmt.Sprintf(format, v...))
}

// Error writes error messages into errorlogger as same as basic log.Error
func Error(v ...interface{}) {
	errorLogger.Output(2, fmt.Sprint(v...))
}

// Errorln writes error messages into errorlogger as same as basic log.Errorln
func Errorln(v ...interface{}) {
	errorLogger.Output(2, fmt.Sprintln(v...))
}

// Errorf writes error messages into errorlogger as same as basic log.Errorf
func Errorf(format string, v ...interface{}) {
	errorLogger.Output(2, fmt.Sprintf(format, v...))
}

// Fatal writes fatal error messages into errorlogger and exit as same as basic log.Fatal
func Fatal(v ...interface{}) {
	fatalLogger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalln writes fatal error messages into errorlogger and exit as same as basic log.Fatalln
func Fatalln(v ...interface{}) {
	fatalLogger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Fatalf writes fatal error messages into errorlogger and exit as same as basic log.Fataf
func Fatalf(format string, v ...interface{}) {
	fatalLogger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Panic writes panic error messages into errorlogger and exit as same as basic log.Panic
func Panic(v ...interface{}) {
	panicLogger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Panicln writes Panic error messages into errorlogger and exit as same as basic log.Panicln
func Panicln(v ...interface{}) {
	panicLogger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Panicf writes Panic error messages into errorlogger and exit as same as basic log.Panicf
func Panicf(format string, v ...interface{}) {
	panicLogger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
