package lib

import (
	"log"
	"os"
)

func Debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Printf(format, v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
