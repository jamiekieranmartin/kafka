package kafka

import (
	"crypto/tls"
	"log"

	kafka "github.com/segmentio/kafka-go"
	scram "github.com/segmentio/kafka-go/sasl/scram"
)

type Config struct {
	Username string
	Password string
	Brokers  []string
}

type Client struct {
	config *Config
	dialer *kafka.Dialer
}

func New(config *Config) *Client {
	mechanism, err := scram.Mechanism(scram.SHA256, config.Username, config.Password)
	if err != nil {
		log.Fatalln(err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	return &Client{
		config: config,
		dialer: dialer}
}
