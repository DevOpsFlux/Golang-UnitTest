package main

import (
	"github.com/gogap/logrus_mate"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus_mate.Hijack(
		logrus.StandardLogger(),
		logrus_mate.ConfigString(
			`{formatter.name = "json"}`,
		),
	)

	logrus.WithField("Field", "A").Debugln("Hello JSON")
}
