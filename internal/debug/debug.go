package debug

import (
	"log"
	"os"
)

var IsDebug = os.Getenv("DEBUG") == "true"

// wrap logging functions with a conditional to check if the debug flag is set
func Debugf(message string, args ...interface{}) {
	if IsDebug {
		log.Printf(message, args...)
	}
}
