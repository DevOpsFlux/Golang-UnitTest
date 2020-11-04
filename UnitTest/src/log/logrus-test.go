package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

func init() {

	// open a file
	f, err := os.OpenFile("log/logrus-test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	// don't forget to close it
	//defer f.Close()
}

func main() {

	log.WithFields(log.Fields{
		"Animal": "Logrus",
	}).Info("A logrus appears")

}
