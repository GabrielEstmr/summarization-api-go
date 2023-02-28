package main_gateways

import main_domains "mpindicator/main/domains"

type SendIndicatorTriggerGateway interface {
	Send(main_domains.IndicatorProcessorTrigger)
}
