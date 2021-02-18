package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	data, err := os.Create("./logs/logrus-logtest-20210217.log")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()
	log.SetOutput(data)
	log.Println("Logggggg")
	log.Fatalln("Fatal Error")
}
