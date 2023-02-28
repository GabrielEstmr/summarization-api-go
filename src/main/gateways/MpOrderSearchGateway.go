package main_gateways

import main_domains "mpindicator/main/domains"

type MpOrderSearchGateway interface {
	GetOrders(main_domains.OrderFilter) main_domains.OrderPage
}
