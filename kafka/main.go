package main

import (
	"bufio"
	"fmt"
	producer2 "leetcode213/kafka/producer"
	"os"
)

func main() {
	producer := producer2.MustNewProducer2(os.Args[1], "honey")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		i, e := producer.SendMsg(text)
		if i == 0 || e != nil {
			panic(e)
		}
	}
}
