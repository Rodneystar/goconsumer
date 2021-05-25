package goconsumerapp

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func Run() {
	c := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"auto.offset.reset": "earliest",
		"group.id":          "test",
	}

	topics := []string{"test1", "test2"}
	lc := NewLoggingConsumer(c, topics)

	lc.Start()
}

type LoggingConsumer struct {
	cfg      kafka.ConfigMap
	consumer *kafka.Consumer
	topics   []string
}

func (lc *LoggingConsumer) Handle(msg *kafka.Message) {
	fmt.Printf("offset: %d, key: %s\nvalue: %s", msg.TopicPartition.Offset, msg.Key, msg.Value)
}

func (lc *LoggingConsumer) Start() {
	lc.consumer.SubscribeTopics(lc.topics, nil)
	defer func() { lc.consumer.Close() }()
	for {
		msg, err := lc.consumer.ReadMessage(-1)
		if err != nil {
			continue
		}
		lc.Handle(msg)
	}
}

func NewLoggingConsumer(c *kafka.ConfigMap, topics []string) *LoggingConsumer {
	consumer, _ := kafka.NewConsumer(c)

	return &LoggingConsumer{
		consumer: consumer,
		cfg:      *c,
		topics:   topics,
	}
}
