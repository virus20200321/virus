package consumer

import (
	"context"
	"github.com/fatih/color"
	"github.com/segmentio/kafka-go"
	"time"
)

func RunReader(ctx context.Context, addr string, groupID string, topic string) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{addr},
		GroupID:        groupID,
		Topic:          topic,
		CommitInterval: time.Second,
	})
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		color.Green("message at topic/partition/offset %v/%v/%v: %s = %s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	e := r.Close()
	color.Yellow("reader close error %v", e)
}
