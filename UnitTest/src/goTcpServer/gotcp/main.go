package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

)

type Callback struct{}

func (this *Callback) OnConnect(c *Conn) bool {
	addr := c.GetRawConn().RemoteAddr()
	c.PutExtraData(addr)
	fmt.Println("OnConnect:", addr)
	return true
}

func (this *Callback) OnMessage(c *Conn, p Packet) bool {
	echoPacket := p.(*EchoPacket)
	fmt.Printf("OnMessage:[%v] [%v]\n", echoPacket.GetLength(), string(echoPacket.GetBody()))
	c.AsyncWritePacket(NewEchoPacket(echoPacket.Serialize(), true), time.Second)
	return true
}

func (this *Callback) OnClose(c *Conn) {
	fmt.Println("OnClose:", c.GetExtraData())
}


func main() {
	// creates a tcp listener
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", ":19309")
	tcpAddr, err := net.ResolveTCPAddr("211.43.192.78", ":19309")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// creates a server
	config := &Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}
	srv := NewServer(config, &Callback{}, &EchoProtocol{})

	// starts service
	go srv.Start(listener, time.Second)
	fmt.Println("listening:", listener.Addr())

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)

	// stops service
	srv.Stop()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
