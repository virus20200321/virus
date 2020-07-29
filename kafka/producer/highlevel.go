package producer

import (
	"context"
	"github.com/fatih/color"
	"github.com/segmentio/kafka-go"
)

func RunWriter(topic string, addr string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{addr},
		Topic:    topic,
		Balancer: &kafka.Hash{},
	})

	e := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("666"),
		},
	)
	color.Yellow("writer writes error: %v", e)

	e = w.Close()
	color.Yellow("writer close error: %v", e)

}
