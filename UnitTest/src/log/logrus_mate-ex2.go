package main

import (
	"github.com/gogap/logrus_mate"
	"github.com/sirupsen/logrus"
)

func main() {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigFile(
			"mate.conf", // { mike {formatter.name = "text"} }
		),
	)

	mate.Hijack(logrus.StandardLogger(),
		"mike",
		logrus_mate.ConfigString(
			`{formatter.name = "json"}`,
		),
	)

	logrus.Errorln("hello std logger is hijack by mike")
}

/* // ex1
func main() {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			`{ mike {formatter.name = "json"} }`,
		),
		logrus_mate.ConfigFile(
			"mate.conf", // { mike {formatter.name = "text"} }
		),
	)

	mate.Hijack(
		logrus.StandardLogger(),
		"mike",
	)

	logrus.Println("hello std logger is hijack by mike")
}
*/
/* // ex2
func main() {
    mate, _ := logrus_mate.NewLogrusMate(
        logrus_mate.ConfigString(
            `{ mike {formatter.name = "json"} }`,
        ),
    )

    mikeLoger := mate.Logger("mike")
    mikeLoger.Errorln("Hello Error Level from Mike and my formatter is json")
}
*/

/* // ex3
func main() {
    mate, _ := logrus_mate.NewLogrusMate(
        logrus_mate.ConfigString(
            `{ mike {formatter.name = "json"} }`,
        ),
    )

    mate.Hijack(
        logrus.StandardLogger(),
        "mike",
    )

    logrus.Println("hello std logger is hijack by mike")
}
*/

/* // ex4
func main() {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			`{ mike {formatter.name = "json"} }`,
		),
		logrus_mate.ConfigFile(
			"mate.conf", // { mike {formatter.name = "text"} }
		),
	)

	mate.Hijack(
		logrus.StandardLogger(),
		"mike",
	)

	logrus.Println("hello std logger is hijack by mike")
}

*/
