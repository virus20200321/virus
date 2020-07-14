package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"

	producer2 "leetcode213/kafka/producer"
)

var (
	app         = kingpin.New("widget", "widget tool")
	producerCmd = app.Command("producer", "producer")
	addr        = producerCmd.Arg("addr", "address").Required().String()
	//consumer
	consumer = app.Command("consumer", "consumer")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case producerCmd.FullCommand():
		fmt.Println("server address:", *addr)
		producer2.RunWriter(*addr)

	}

}
