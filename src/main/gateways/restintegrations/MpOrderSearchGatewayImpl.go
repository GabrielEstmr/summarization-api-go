package main_gateways_restintegrations

import (
	main_domains "mpindicator/main/domains"
	main_gateways_restintegrations_mpordersearch "mpindicator/main/gateways/restintegrations/mpordersearch"
)

type MpOrderSearchGatewayImpl struct {
	mpOrderSearchIntegration main_gateways_restintegrations_mpordersearch.MpOrderSearchIntegration
}

func NewMpOrderSearchGatewayImpl(
	mpOrderSearchIntegration *main_gateways_restintegrations_mpordersearch.MpOrderSearchIntegration) *MpOrderSearchGatewayImpl {
	return &MpOrderSearchGatewayImpl{mpOrderSearchIntegration: *mpOrderSearchIntegration}
}

func (gateway *MpOrderSearchGatewayImpl) GetOrders(orderFilter main_domains.OrderFilter) main_domains.OrderPage {
	response := gateway.mpOrderSearchIntegration.GetOrder(
		orderFilter.OrderBeginDate,
		orderFilter.OrderEndDate,
		orderFilter.InvoiceStatus,
		orderFilter.Page,
		orderFilter.PageSize,
		orderFilter.SellerId,
		orderFilter.OrderStatusList)
	return response.ToDomain()
}
