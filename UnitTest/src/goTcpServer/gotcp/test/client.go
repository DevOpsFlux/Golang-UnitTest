package main

import (
	"fmt"
	"log"
	"net"
	"time"
	//"echo"
	"github.com/gansidui/gotcp/examples/echo"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "211.43.192.78:19309")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	echoProtocol := &echo.EchoProtocol{}

	// ping <--> pong
	for i := 0; i < 3; i++ {
		// write
		conn.Write(echo.NewEchoPacket([]byte("004020010171950456                      "), false).Serialize())

		// read
		p, err := echoProtocol.ReadPacket(conn)
		if err == nil {
			echoPacket := p.(*echo.EchoPacket)
			fmt.Printf("Server reply:[%v] [%v]\n", echoPacket.GetLength(), string(echoPacket.GetBody()))
		}

		time.Sleep(2 * time.Second)
	}

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}