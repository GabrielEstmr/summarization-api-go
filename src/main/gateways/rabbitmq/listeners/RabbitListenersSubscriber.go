package main_gateways_rabbitmq_listeners

func RabbitListenersSubscriber() {
	go Listen()
	go IndicatorProcessorListener()
}
