package main_usecases_beans

import (
	main_gateways "mpindicator/main/gateways"
	main_gateways_mongodb "mpindicator/main/gateways/mongodb"
	main_gateways_mongodb_repositories "mpindicator/main/gateways/mongodb/repositories"

	// main_gateways_restintegrations "mpindicator/main/gateways/restintegrations"
	// main_gateways_restintegrations_mpordersearch "mpindicator/main/gateways/restintegrations/mpordersearch"
	main_usecases "mpindicator/main/usecases"
)

func GetCreateUserBeanFunc() *main_usecases.CreateUser {
	repo := main_gateways_mongodb_repositories.NewUserRepository()

	var gateway main_gateways.UserDatabaseGateway
	gatewayImpl := main_gateways_mongodb.NewUserDatabaseGatewayImpl(repo)
	gateway = gatewayImpl

	// mpOrderSearchIntegration := main_gateways_restintegrations_mpordersearch.NewMpOrderSearchIntegration()

	// var mpOrderSearchGateway main_gateways.MpOrderSearchGateway
	// mpOrderSearchGatewayImpl := main_gateways_restintegrations.NewMpOrderSearchGatewayImpl(mpOrderSearchIntegration)
	// mpOrderSearchGateway = mpOrderSearchGatewayImpl

	useCase := main_usecases.NewCreateUser(&gateway)
	return useCase
}
