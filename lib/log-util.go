package lib

import (
	"log"
	"os"
)

var IsDebug = os.Getenv("DEBUG") != ""

func Debug(v ...interface{}) {
	if IsDebug {
		log.Println(v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if IsDebug {
		log.Printf(format, v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
