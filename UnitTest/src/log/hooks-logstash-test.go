package main

import (
	"fmt"
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
	ctx.Info("logrus-logstash-hook Dev Unit Test")

	for i := 1; i <= 10000; i++ {
		ctx.Info("logrus-logstash-hook Dev Unit Test" + fmt.Sprint(i))
	}
}

// {"@timestamp":"2020-11-04T18:41:51+09:00","@version":"1","level":"info","message":"logrus-logstash-hook Dev Unit Test10000","method":"main","type":"myappName"}
