package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/streadway/amqp"
)

var (
	broker string
)

func init() {
	flag.StringVar(&broker, "broker", "rabbitmq", "Choices - rabbitmq/kafka")

	flag.Parse()
}

func main() {
	log.Printf("=> bootstrap %v", broker)

	switch broker {

	case "rabbitmq":
		setupRabbitMQ()

	case "kafka":
		setupKafka()

	default:
		log.Fatalf("incorrect message broker %v", broker)
	}
}

const (
	dlExchange       = "elk-dead-letter"
	dlQueue          = "elk-dead-letter-queue"
	unroutedExchange = "elk-unrouted"
	elkExchange      = "elk-exchange"
	elkQueue         = "elk-queue"
	unroutedQueue    = "elk-unrouted-queue"
)

func setupRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	handleErr(err)
	defer conn.Close()

	// 1. DLQ
	dqchan, err := conn.Channel()
	handleErr(err)

	err = dqchan.ExchangeDeclare(dlExchange, "direct", true, false, false, false, nil)
	handleErr(err)

	dq, err := dqchan.QueueDeclare(dlQueue, true, false, false, false, nil)
	handleErr(err)

	err = dqchan.QueueBind(dq.Name, "#", dlExchange, false, nil)
	handleErr(err)

	// 2. ALT
	achan, err := conn.Channel()
	handleErr(err)

	err = achan.ExchangeDeclare(unroutedExchange, "direct", true, false, false, false, nil)
	handleErr(err)

	aq, err := achan.QueueDeclare(unroutedQueue, true, false, false, false, nil)
	handleErr(err)

	err = achan.QueueBind(aq.Name, "#", unroutedExchange, false, nil)
	handleErr(err)

	// 3. ELK
	rqchan, err := conn.Channel()
	handleErr(err)

	err = rqchan.ExchangeDeclare(elkExchange, "x-consistent-hash", true, false, false, false,
		amqp.Table{"alternate-exchange": unroutedExchange})
	handleErr(err)

	for qID := 1; qID <= 2; qID++ {
		qName := fmt.Sprintf("%v-0%v", elkQueue, qID)
		q, err := rqchan.QueueDeclare(qName, true, false, false, false,
			amqp.Table{"x-max-priority": uint8(10), "x-dead-letter-exchange": dlExchange})
		handleErr(err)

		err = rqchan.QueueBind(q.Name, "10", elkExchange, false, nil)
		handleErr(err)
	}
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
		log.Fatalln(err)
	}
}

const (
	elkTopic = "elk-syslog"
)

func setupKafka() {
	_, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", elkTopic, 0)
	handleErr(err)
}
