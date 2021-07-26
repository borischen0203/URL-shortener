package logger

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	// Debug print out the debug log
	Debug *log.Logger
	// Info print out the info log
	Info *log.Logger
	// Error print out the error log
	Error *log.Logger
)

// Setup setup Logger
func Setup() {
	errorHandle := os.Stderr
	infoHandle := os.Stdout
	debugHandle := ioutil.Discard

	if v, _ := os.LookupEnv("VERBOSE"); strings.ToLower(v) == "true" {
		debugHandle = os.Stdout
	}

	Debug = log.New(debugHandle, "DEB: ", log.Ldate|log.Lmicroseconds|log.LUTC)
	Info = log.New(infoHandle, "INF: ", log.Ldate|log.Lmicroseconds|log.LUTC)
	Error = log.New(errorHandle, "ERR: ", log.Ldate|log.Lmicroseconds|log.LUTC)
}
