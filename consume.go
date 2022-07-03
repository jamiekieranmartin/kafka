package kafka

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func (c *Client) Consume(topic, group string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: c.config.Brokers,
		GroupID: group,
		Topic:   topic,
		Dialer:  c.dialer,
	})

	defer r.Close()

	fmt.Printf("[%s]: Consuming from topic: %s\n", group, topic)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
