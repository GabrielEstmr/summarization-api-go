package main_usecases

import (
	main_domains "mpindicator/main/domains"
	main_gateways "mpindicator/main/gateways"
	main_utils "mpindicator/main/utils"
)

type GetThirtyDaysInTransitOrders struct {
	mpOrderSearchGateway main_gateways.MpOrderSearchGateway
}

func NewGetThirtyDaysInTransitOrders(mpOrderSearchGateway *main_gateways.MpOrderSearchGateway) *GetThirtyDaysInTransitOrders {
	return &GetThirtyDaysInTransitOrders{mpOrderSearchGateway: *mpOrderSearchGateway}
}

func (thisUseCase *GetThirtyDaysInTransitOrders) Execute(page int64) main_domains.OrderPage {

	result := thisUseCase.mpOrderSearchGateway.GetOrders(main_domains.OrderFilter{
		Page:            page,
		OrderStatusList: []string{"Shipped"},
		PageSize:        1000,
		OrderBeginDate:  main_utils.GetThirtyDaysInThePastAtStartOfDayAsString(main_domains.YYYYMMDD),
		OrderEndDate:    main_utils.GetTodayAsString(main_domains.YYYYMMDD),
	})

	return result
}
