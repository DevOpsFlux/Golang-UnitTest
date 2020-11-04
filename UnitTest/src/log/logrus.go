package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	// open a file
	f, err := os.OpenFile("log/testlogrus1.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Trace("Trace")
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Debug("Debug")
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Info("Info")
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Warn("Warn")
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Error("Error")

	logrus.SetOutput(f)

	logrus.WithFields(logrus.Fields{
		"Animal": "Logrus",
	}).Info("A logrus appears")

	logrus.Println("logrus Unit Test")

	/*
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})

		logrus.SetReportCaller(false)

		logrus.SetOutput(os.Stdout)

		logrus.SetLevel(logrus.TraceLevel)

		logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Trace("Trace")
		logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Debug("Debug")
		logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Info("Info")
		logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Warn("Warn")
		logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Error("Error")

		logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Fatal("Fatal")
		logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Panic("Panic")

		entry := logrus.WithFields(logrus.Fields{"Fields": "TEST"})
		entry.Trace("Trace logger")
		entry.Debug("Debug logger")
		entry.Info("Info logger")
		entry.Warn("Warn logger")
		entry.Error("Error logger")

		entry.Fatal("Fatal logger")
		entry.Panic("Panic logger")
	*/

}
