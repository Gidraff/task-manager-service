package logger

import (
	"log"
	"os"
)

const (
	// PathToLog specify path to lpg to
	PathToLog string = "Users/gidraffkamande/Projects/Private/go/src/github.com/Gidraff/task-manager-service/test.log"
)

func SetUpLogger() {
	logFileLocation, _ := os.OpenFile(
		PathToLog,                         // Open file at path
		os.O_CREATE|os.O_APPEND|os.O_RDWR, // specify flag
		0744,                              // file mode
	)
	log.SetOutput(logFileLocation)
}
