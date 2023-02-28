package main_usecases_beans

import (
	main_gateways "mpindicator/main/gateways"
	main_gateways_rabbitmq_producers "mpindicator/main/gateways/rabbitmq/producers"
	main_usecases "mpindicator/main/usecases"
)

func GetProcessApiIndicatorTriggerBeanFunc() *main_usecases.ApiIndicatorTrigger {
	var gateway main_gateways.SendIndicatorTriggerGateway
	gatewayImpl := main_gateways_rabbitmq_producers.NewSendIndicatorTriggerGatewayImpl()
	gateway = gatewayImpl
	useCase := main_usecases.NewApiIndicatorTrigger(&gateway)
	return useCase
}
