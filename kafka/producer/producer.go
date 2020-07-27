package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type Producer struct {
	Conn *kafka.Conn
}

func MustNewProducer2(addr string, topic string) *Producer {
	conn, e := kafka.DialLeader(context.Background(), "tcp", addr, topic, 0)
	if e != nil {
		panic(e)
	}
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	return &Producer{Conn: conn}
}

//"10.13.58.111:9092"
func MustNewProducer(addr string, topic string, numPartition int) *Producer {
	producer, e := kafka.DefaultDialer.Dial("tcp", addr)
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
