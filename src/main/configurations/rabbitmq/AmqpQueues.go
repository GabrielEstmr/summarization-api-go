package main_configurations_rabbitmq

var AmqpQueues = map[string]string{
	"mp-indicator-aqmpq-test":             "AmqpRoutingKey",
	"mp-indicator-go-indicator-processor": "IndicatorProcessor",
}
