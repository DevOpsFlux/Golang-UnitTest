package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

func main() {

	getpath, _ := os.Getwd()
	fmt.Println(getpath)

	ex, err := os.Executable()
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	dir := path.Dir(ex)
	fmt.Println(dir)

	// open a file
	// ./projects/UnitT/log/test.log
	// log/test.log
	//f, err := os.OpenFile("./projects/UnitT/log/test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	f, err := os.OpenFile("./projects/UnitT/log/test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{"Fields": "TEST"}).Debug("Debug")
	logrus.SetOutput(f)

	logrus.WithFields(logrus.Fields{
		"Animal": "Logrus",
	}).Info("A logrus appears")

	logrus.Println("logrus Unit Test")


}
