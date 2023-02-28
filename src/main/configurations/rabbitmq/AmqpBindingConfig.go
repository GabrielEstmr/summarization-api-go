package main_configurations_rabbitmq

import "sync"

type AmqpQueueProperties struct {
	BindingId        string
	QueueName        string
	RoutingKey       string
	Exchange         string
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	Nowait           bool
}

var AmqpBindingParameters *map[string]AmqpQueueProperties = nil

var amqpPropertiesConfigList = []AmqpQueueProperties{
	{
		BindingId:        AmqpRoutingKeys["AmqpRoutingKey"] + AmqpExchanges["logs_topic"] + AmqpRoutingKeys["AmqpRoutingKey"] + "mp-indicator-aqmpq-test",
		QueueName:        "mp-indicator-aqmpq-test",
		RoutingKey:       AmqpRoutingKeys["AmqpRoutingKey"],
		Exchange:         AmqpExchanges["logs_topic"],
		Durable:          false,
		DeleteWhenUnused: false,
		Exclusive:        false,
		Nowait:           false,
	},
	{
		BindingId:        AmqpRoutingKeys["IndicatorProcessor"] + AmqpExchanges["logs_topic"] + AmqpRoutingKeys["IndicatorProcessor"] + "mp-indicator-go-indicator-processor",
		QueueName:        "mp-indicator-go-indicator-processor",
		RoutingKey:       AmqpRoutingKeys["IndicatorProcessor"],
		Exchange:         AmqpExchanges["logs_topic"],
		Durable:          false,
		DeleteWhenUnused: false,
		Exclusive:        false,
		Nowait:           false,
	},
}

func GetBindingParametersBean() *map[string]AmqpQueueProperties {
	var once sync.Once
	once.Do(func() {
		if AmqpBindingParameters == nil {
			AmqpBindingParameters = getBindingParameters()
		}
	})
	return AmqpBindingParameters
}

func getBindingParameters() *map[string]AmqpQueueProperties {
	queueMapConfigs := make(map[string]AmqpQueueProperties)
	for _, value := range amqpPropertiesConfigList {
		queueMapConfigs[value.BindingId] = value
	}
	return &queueMapConfigs
}
