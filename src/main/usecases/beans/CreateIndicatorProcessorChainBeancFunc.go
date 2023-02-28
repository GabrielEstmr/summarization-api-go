package main_usecases_beans

import (
	main_gateways "mpindicator/main/gateways"

	main_gateways_mongodb "mpindicator/main/gateways/mongodb"
	main_gateways_mongodb_repositories "mpindicator/main/gateways/mongodb/repositories"
	main_gateways_restintegrations "mpindicator/main/gateways/restintegrations"
	main_gateways_restintegrations_mpordersearch "mpindicator/main/gateways/restintegrations/mpordersearch"
	main_usecases "mpindicator/main/usecases"
)

func CreateIndicatorProcessorChainBeancFunc() *main_usecases.IndicatorProcessorChain {

	var indicatorProcessor main_usecases.IndicatorProcessor

	mpOrderSearchIntegration := main_gateways_restintegrations_mpordersearch.NewMpOrderSearchIntegration()

	var mpOrderSearchGateway main_gateways.MpOrderSearchGateway
	mpOrderSearchGatewayImpl := main_gateways_restintegrations.NewMpOrderSearchGatewayImpl(mpOrderSearchIntegration)
	mpOrderSearchGateway = mpOrderSearchGatewayImpl

	getThirtyDaysInTransitOrders := main_usecases.NewGetThirtyDaysInTransitOrders(&mpOrderSearchGateway)

	var indicatorDatabaseGateway main_gateways.IndicatorDatabaseGateway
	indicatorRepository := main_gateways_mongodb_repositories.NewIndicatorRepository()
	indicatorDatabaseGatewayImpl := main_gateways_mongodb.NewIndicatorDatabaseGatewayImpl(indicatorRepository)
	indicatorDatabaseGateway = indicatorDatabaseGatewayImpl

	saveIndicatorValues := main_usecases.NewSaveIndicatorValues(&indicatorDatabaseGateway)

	indicatorProcessorImpl := main_usecases.NewThirtyDaysInTransitOrdersProcessor(getThirtyDaysInTransitOrders, saveIndicatorValues)

	indicatorProcessor = indicatorProcessorImpl

	indicatorProcessorChain := main_usecases.NewIndicatorProcessorChain(&indicatorProcessor)

	return indicatorProcessorChain
}
