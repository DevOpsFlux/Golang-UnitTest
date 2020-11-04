package main

import (
	"log/syslog"

	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

func main() {
	log := logrus.New()
	hook, err := lSyslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")

	if err == nil {
		log.Hooks.Add(hook)
	}

}
