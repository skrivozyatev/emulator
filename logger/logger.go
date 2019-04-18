package logger

import (
	"log"
	"os"
)

var info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
var err = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)

func Info(v ...interface{}) {
	info.Print(v...)
}

func Infof(format string, v ...interface{}) {
	info.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	err.Fatalln(v...)
}

func Fatalf(format string, v ...interface{}) {
	err.Fatalf(format, v...)
}
