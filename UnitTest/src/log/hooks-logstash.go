package main

import (
	"net"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	conn, err := net.Dial("tcp", "127.0.0.1:1470") // logstash.mydev.com:8911
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "myappName"}))

	log.Hooks.Add(hook)
	ctx := log.WithFields(logrus.Fields{
		"method": "main",
	})
	ctx.Info("Hello World!")
}
