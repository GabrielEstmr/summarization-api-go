package main_gateways_rabbitmq_listeners

import (
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
	main_configurations_yml "mpindicator/main/configurations/yml"
	main_domains "mpindicator/main/domains"
	main_usecases_beans "mpindicator/main/usecases/beans"
	main_utils "mpindicator/main/utils"
)

const MSG_ERROR_TO_READ_MESSAGE_FROM_TOPIC = "Error to load message from topic %s"
const MSG_ERROR_MESSAGE_RECEIVED = "Message received at topic/partition/offset %v/%v/%v: %s = %s\n"

func IndicatorProcessorListener() {

	rabbitMqURI := main_configurations_yml.GetBeanPropertyByName(RABBITMQ_URI)
	log.Println("========> rabbitMqURI LISTENER", rabbitMqURI)

	conn, err := amqp091.Dial(rabbitMqURI)
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer conn.Close()

	ch, err := conn.Channel()
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer ch.Close()

	msgs, err := ch.Consume(
		"mp-indicator-go-indicator-processor", // queue
		"",                                    // consumer
		true,                                  // auto ack
		false,                                 // exclusive
		false,                                 // no local
		false,                                 // no wait
		nil,                                   // args
	)
	main_utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x]==> %s", d.Body)
			var indicatorProcessorTrigger main_domains.IndicatorProcessorTrigger
			errorJSON := json.Unmarshal(d.Body, &indicatorProcessorTrigger)
			main_utils.FailOnError(errorJSON, "errorJSON.Error()")

			log.Println("IndicatorProcessorTrigger", indicatorProcessorTrigger)

			main_usecases_beans.CreateIndicatorProcessorChainBeancFunc().Execute(context.Background(), indicatorProcessorTrigger)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

	//consumer := main_configurations_kafka.BuildKafkaConsumerByTopic(main_configurations_kafka.MP_INDICATOR_INDICATOR_PROCESSOR)
	//
	//for {
	//	m, err := consumer.ReadMessage(context.Background())
	//	main_utils.FailOnError(err,
	//		fmt.Sprintf(MSG_ERROR_TO_READ_MESSAGE_FROM_TOPIC,
	//			main_configurations_kafka.MP_INDICATOR_INDICATOR_PROCESSOR.GetDescription()))
	//	fmt.Printf(MSG_ERROR_MESSAGE_RECEIVED, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	//

	//
	//}
}
