package main_gateways_restintegrations_mpordersearch_responses

import main_domains "mpindicator/main/domains"

type OrderSearchSellerOrderResponse struct {
	OrderNumber string                    `json:"orderNumber,omitempty"`
	Seller      OrderSearchSellerResponse `json:"seller,omitempty"`
}

func (orderResponse *OrderSearchSellerOrderResponse) ToDomain() main_domains.SellerOrder {
	return main_domains.SellerOrder{
		OrderNumber: orderResponse.OrderNumber,
		SellerId:    orderResponse.Seller.Code,
	}
}
