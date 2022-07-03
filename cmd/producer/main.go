package main

import (
	"flag"
	"fmt"

	"github.com/jamiekieranmartin/kafka"
)

const cliVersion = "0.0.1"

const helpMessage = `
kafka-producer v%s

	Example:

	kafka-producer -username=my-username -password=my-password -broker=my-broker -topic=my-topic

	`

func main() {

	flag.Usage = func() {
		fmt.Printf(helpMessage, cliVersion)
		flag.PrintDefaults()
	}

	// cli arguments
	version := flag.Bool("version", false, "Print version string and exit")
	help := flag.Bool("help", false, "Print help message and exit")

	username := flag.String("username", "", "Kafka username")
	password := flag.String("password", "", "Kafka password")
	broker := flag.String("broker", "", "Kafka broker")
	topic := flag.String("topic", "", "Kafka topic")

	flag.Parse()

	// if asked for version, disregard everything else
	if *version {
		fmt.Printf("go-template v%s\n", cliVersion)
		return
	} else if *help {
		flag.Usage()
		return
	}

	if *username == "" || *password == "" || *broker == "" || *topic == "" {
		flag.Usage()
		return
	}

	k := kafka.New(&kafka.Config{
		Username: *username,
		Password: *password,
		Brokers:  []string{*broker},
	})

	k.Produce(*topic)
}
