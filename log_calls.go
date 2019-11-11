package log

import (
	"fmt"
	"os"
)

// Info writes into infologger as same as basic fmt.Print
func Info(v ...interface{}) {
	if err := infoLogger.Output(2, fmt.Sprint(v...)); err != nil {
		fmt.Println(err.Error())
	}
}

// Infoln writes into infologger as same as basic fmt.Println
func Infoln(v ...interface{}) {
	if err := infoLogger.Output(2, fmt.Sprintln(v...)); err != nil {
		fmt.Println(err.Error())
	}
}

// Infof writes into infologger as same as basic Outputf
func Infof(format string, v ...interface{}) {
	if err := infoLogger.Output(2, fmt.Sprintf(format, v...)); err != nil {
		fmt.Println(err.Error())
	}
}

// Warning writes warning messages into warninglogger as same as basic log.Output
func Warning(v ...interface{}) {
	if err := warningLogger.Output(2, fmt.Sprint(v...)); err != nil {
		fmt.Println(err.Error())
	}
}

// Warningln writes warning messages into warninglogger as same as basic log.Outputln
func Warningln(v ...interface{}) {
	if err := warningLogger.Output(2, fmt.Sprintln(v...)); err != nil {
		fmt.Println(err.Error())
	}
}

// Warningf writes warning messages into warninglogger as same as basic log.Outputf
func Warningf(format string, v ...interface{}) {
	if err := warningLogger.Output(2, fmt.Sprintf(format, v...)); err != nil {
		fmt.Println(err.Error())
	}
}

// Error writes error messages into errorlogger as same as basic log.Error
func Error(v ...interface{}) {
	if err := errorLogger.Output(2, fmt.Sprint(v...)); err != nil {
		fmt.Println(err.Error())
	}
}

// Errorln writes error messages into errorlogger as same as basic log.Errorln
func Errorln(v ...interface{}) {
	if err := errorLogger.Output(2, fmt.Sprintln(v...)); err != nil {
		fmt.Println(err.Error())
	}
}

// Errorf writes error messages into errorlogger as same as basic log.Errorf
func Errorf(format string, v ...interface{}) {
	if err := errorLogger.Output(2, fmt.Sprintf(format, v...)); err != nil {
		fmt.Println(err.Error())
	}
}

// Fatal writes fatal error messages into errorlogger and exit as same as basic log.Fatal
func Fatal(v ...interface{}) {
	if err := fatalLogger.Output(2, fmt.Sprint(v...)); err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

// Fatalln writes fatal error messages into errorlogger and exit as same as basic log.Fatalln
func Fatalln(v ...interface{}) {
	if err := fatalLogger.Output(2, fmt.Sprintln(v...)); err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

// Fatalf writes fatal error messages into errorlogger and exit as same as basic log.Fataf
func Fatalf(format string, v ...interface{}) {
	if err := fatalLogger.Output(2, fmt.Sprintf(format, v...)); err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

// Panic writes panic error messages into errorlogger and exit as same as basic log.Panic
func Panic(v ...interface{}) {
	if err := panicLogger.Output(2, fmt.Sprint(v...)); err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

// Panicln writes Panic error messages into errorlogger and exit as same as basic log.Panicln
func Panicln(v ...interface{}) {
	if err := panicLogger.Output(2, fmt.Sprintln(v...)); err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

// Panicf writes Panic error messages into errorlogger and exit as same as basic log.Panicf
func Panicf(format string, v ...interface{}) {
	if err := panicLogger.Output(2, fmt.Sprintf(format, v...)); err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}
