package logging

import (
	"fmt"
	"os"

	mylog "github.com/sirupsen/logrus"
)

// InitializeLogging asdas
func InitializeLogging(logFile string) {
	var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	mylog.SetOutput(file)
	//log.SetFormatter(&log.TextFormatter{})
	mylog.SetFormatter(&mylog.JSONFormatter{})
}
