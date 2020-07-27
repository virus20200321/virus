package main

import (
	"context"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	consumer2 "leetcode213/kafka/consumer"
	producer2 "leetcode213/kafka/producer"
	"os"
)

var (
	app         = kingpin.New("widget", "widget tool")
	producerCmd = app.Command("producer", "producer")
	addr        = producerCmd.Arg("addr", "address").Required().String()
	topic       = producerCmd.Arg("topic", "topic").Required().String()
	//consumer
	consumer = app.Command("consumer", "consumer")
	addr2    = consumer.Arg("addr", "address").Required().String()
	groupID  = consumer.Arg("group", "group").Required().String()
	topic2   = consumer.Arg("topic", "topic").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case producerCmd.FullCommand():
		fmt.Println("server address:", *addr)
		producer2.RunWriter(*topic, *addr)
	case consumer.FullCommand():
		consumer2.RunReader(context.Background(), *addr2, *groupID, *topic2)
	}

}
