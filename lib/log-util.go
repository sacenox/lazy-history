package lib

import (
	"log"
	"os"

	"github.com/kr/pretty"
)

var IsDebug = os.Getenv("DEBUG") != ""

func Debug(v ...interface{}) {
	if IsDebug {
		PrettyPrint(v)
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

func PrettyPrint(v interface{}) {
	log.Printf("%# v\n", pretty.Formatter(v))
}
