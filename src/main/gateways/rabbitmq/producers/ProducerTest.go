package main_gateways_rabbitmq_producers

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
	main_configurations_yml "mpindicator/main/configurations/yml"
	main_utils "mpindicator/main/utils"
	"os"
	"strings"
)

const MSG_RABBITMQ_PUBLISH_MESSAGE_FAILURE = "Failed to publish a message"
const MSG_RABBITMQ_CONNECT_FAILURE = "Failed to connect to RabbitMQ"
const RABBITMQ_URI = "RabbitMQ.URI"

func Produce(ctxP *context.Context) {

	rabbitMqURI := main_configurations_yml.GetBeanPropertyByName(RABBITMQ_URI)

	conn, err := amqp091.Dial(rabbitMqURI)
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer conn.Close()

	ch, err := conn.Channel()
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer ch.Close()

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(*ctxP,
		"logs_topic",     // exchange
		"AmqpRoutingKey", // routing key
		false,            // mandatory
		false,            // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	main_utils.FailOnError(err, MSG_RABBITMQ_PUBLISH_MESSAGE_FAILURE)

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "OIEEEE"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}
