package kafka

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func (c *Client) Produce(topic string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: c.config.Brokers,
		Topic:   topic,
		Dialer:  c.dialer,
	})

	defer w.Close()

	fmt.Println("Producing to topic:", topic)
	for {
		u := pseudo_uuid()

		fmt.Println("message: ", u)

		w.WriteMessages(context.Background(), kafka.Message{Value: []byte(u)})

		time.Sleep(time.Duration(15000) * time.Millisecond)
	}
}

// Note - NOT RFC4122 compliant
func pseudo_uuid() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}
