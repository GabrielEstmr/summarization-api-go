package main_configurations_rabbitmq

import (
	"log"
	main_configurations_yml "mpindicator/main/configurations/yml"
	main_utils "mpindicator/main/utils"

	"github.com/rabbitmq/amqp091-go"
)

const MSG_RABBITMQ_DECLARE_EXCHANGE_FAILURE = "Failed to declare an exchange"
const MSG_RABBITMQ_CONNECT_FAILURE = "Failed to connect to RabbitMQ"
const MSG_RABBITMQ_DECLARE_QUEUE_FAILURE = "Failed to declare a queue"
const MSG_ERROR_TO_CLOSE_CONNECTION = "Failed to close rabbitMQ connection"
const MSG_ERROR_TO_CLOSE_CHANNEL = "Failed to close rabbitMQ channel"
const RABBITMQ_URI = "RabbitMQ.URI"

func SetAmqpConfig() {

	rabbitMqURI := main_configurations_yml.GetBeanPropertyByName(RABBITMQ_URI)
	log.Println("========> rabbitMqURI", rabbitMqURI)

	conn, err := amqp091.Dial(rabbitMqURI)
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer closeConnection(conn)

	ch, err := conn.Channel()
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer closeChannel(ch)

	QueueDeclare(ch)

	exchanges := AmqpExchanges
	for _, value := range exchanges {
		err := ch.ExchangeDeclare(
			value,   // name
			"topic", // type
			true,    // durable
			false,   // auto-deleted
			false,   // internal
			false,   // no-wait
			nil,     // arguments
		)
		main_utils.FailOnError(err, MSG_RABBITMQ_DECLARE_EXCHANGE_FAILURE)
	}

	bindingParameters := *GetBindingParametersBean()
	for _, value := range bindingParameters {
		err := ch.QueueBind(
			value.QueueName,  // queue name
			value.RoutingKey, // routing key
			value.Exchange,   // exchange
			false,
			nil)
		main_utils.FailOnError(err, MSG_RABBITMQ_DECLARE_EXCHANGE_FAILURE)
	}
}

func QueueDeclare(ch *amqp091.Channel) {
	amqpQueues := make(map[string]amqp091.Queue)
	rabbitMQQueuesParameters := *GetBindingParametersBean()
	for key, value := range rabbitMQQueuesParameters {
		q, err := ch.QueueDeclare(
			value.QueueName,        // name
			value.Durable,          // durable
			value.DeleteWhenUnused, // delete when unused
			value.Exclusive,        // exclusive
			value.Nowait,           // no-wait
			nil,                    // arguments
		)

		main_utils.FailOnError(err, MSG_RABBITMQ_DECLARE_QUEUE_FAILURE)
		amqpQueues[key] = q
	}
}

func closeConnection(conn *amqp091.Connection) {
	errConnClose := conn.Close()
	main_utils.FailOnError(errConnClose, MSG_ERROR_TO_CLOSE_CONNECTION)
}

func closeChannel(ch *amqp091.Channel) {
	errChClose := ch.Close()
	main_utils.FailOnError(errChClose, MSG_ERROR_TO_CLOSE_CHANNEL)
}
