package main_redix

import (
	"fmt"

	"github.com/mediocregopher/radix/v3" // redis client
	log "github.com/sirupsen/logrus"
)

func main() {

	c, err := radix.NewPool("tcp", "172.0.0.1:6379", 1)
	if err != nil {
		log.Error(err)
	}

	var setVal string
	if err := c.Do(radix.Cmd(&setVal, "SET", "keyone", "valueone")); err != nil {
		log.Error(err)
	}

	var getVal string
	if err := c.Do(radix.Cmd(&getVal, "GET", "keyone")); err != nil {
		log.Error(err)
	}

	fmt.Println(setVal)
	fmt.Println(getVal)

	/*
		err := client.Do(radix.Cmd(nil, "SET", "foo", "someval"))

		var fooVal string
		err := client.Do(radix.Cmd(&fooVal, "GET", "foo"))

		var fooValB []byte
		err := client.Do(radix.Cmd(&fooValB, "GET", "foo"))

		var barI int
		err := client.Do(radix.Cmd(&barI, "INCR", "bar"))

		var bazEls []string
		err := client.Do(radix.Cmd(&bazEls, "LRANGE", "baz", "0", "-1"))

		var buzMap map[string]string
		err := client.Do(radix.Cmd(&buzMap, "HGETALL", "buz"))
	*/
}
