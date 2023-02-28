package main_usecases

import (
	main_domains "mpindicator/main/domains"
	main_gateways "mpindicator/main/gateways"
)

type ApiIndicatorTrigger struct {
	UserDatabaseGateway main_gateways.SendIndicatorTriggerGateway
}

func NewApiIndicatorTrigger(userDatabaseGateway *main_gateways.SendIndicatorTriggerGateway) *ApiIndicatorTrigger {
	return &ApiIndicatorTrigger{UserDatabaseGateway: *userDatabaseGateway}
}

func (thisUseCase *ApiIndicatorTrigger) Execute(user main_domains.IndicatorProcessorTrigger) {

	thisUseCase.UserDatabaseGateway.Send(user)

}
