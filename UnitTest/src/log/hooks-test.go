package main

import (
	"log/syslog"

	log "github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

func init() {

	// Use the Airbrake hook to report errors that have Error severity or above to
	// an exception tracker. You can create custom hooks, see the Hooks section.
	log.AddHook(airbrake.NewHook(123, "xyz", "production"))

	hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
	if err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}
}

func main() {

	log.WithFields(log.Fields{
		"Animal": "Logrus",
	}).Info("A logrus appears")

}
