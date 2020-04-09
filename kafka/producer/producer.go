package producer

import (
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type Producer struct {
	Conn *kafka.Conn
}

func MustNewProducer(topic string, numPartition int) *Producer {
	producer, e := kafka.DefaultDialer.Dial("tcp", "10.13.58.111:9092")
	if e != nil {
		panic(e)
	}
	e = producer.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if e != nil {
		panic(e)
	}
	e = producer.CreateTopics(kafka.TopicConfig{
		Topic:              topic,
		NumPartitions:      numPartition,
		ReplicationFactor:  -1,
		ReplicaAssignments: nil,
		ConfigEntries:      nil,
	})
	if e != nil {
		panic(e)
	}
	return &Producer{Conn: producer}
}

func (p *Producer) SendMsg(msg string) (int, error) {
	return p.Conn.WriteMessages(kafka.Message{
		Value: []byte(msg),
	})
}
