package main

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func main() {

	// to create topics when auto.create.topics.enable='false'
	topic := "my-topic"
	partition := 0

	/*
		// to create topics when auto.create.topics.enable='true'
		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "my-topic", 0)
		if err != nil {
			panic(err.Error())
		}
	*/
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
