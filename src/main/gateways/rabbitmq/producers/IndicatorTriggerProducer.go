package main_gateways_rabbitmq_producers

import (
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
	main_configurations_yml "mpindicator/main/configurations/yml"
	main_domains "mpindicator/main/domains"
	main_utils "mpindicator/main/utils"
	"time"
)

const MSG_ERROR_DURING_SERIALIZE = "Error during serialize"

type SendIndicatorTriggerGatewayImpl struct {
}

func NewSendIndicatorTriggerGatewayImpl() *SendIndicatorTriggerGatewayImpl {
	return &SendIndicatorTriggerGatewayImpl{}
}

func (gateway *SendIndicatorTriggerGatewayImpl) Send(indicatorProcessorTrigger main_domains.IndicatorProcessorTrigger) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rabbitMqURI := main_configurations_yml.GetBeanPropertyByName(RABBITMQ_URI)

	conn, err := amqp091.Dial(rabbitMqURI)
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer conn.Close()

	ch, err := conn.Channel()
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer ch.Close()

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	indicatorBytes, errJSON := json.Marshal(indicatorProcessorTrigger)
	main_utils.FailOnError(errJSON, MSG_ERROR_DURING_SERIALIZE)

	//message := []main_domains.Event{
	//	{
	//		Key:   []byte("1"),
	//		Value: userBytes,
	//	},
	//}

	err = ch.PublishWithContext(ctx,
		"logs_topic",         // exchange
		"IndicatorProcessor", // routing key
		false,                // mandatory
		false,                // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        indicatorBytes,
		})
	main_utils.FailOnError(err, MSG_RABBITMQ_PUBLISH_MESSAGE_FAILURE)

	log.Printf(" [x] Sent %s", indicatorProcessorTrigger)

}
