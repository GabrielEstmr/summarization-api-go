package main_gateways

import main_domains "mpindicator/main/domains"

type IndicatorCachedDatabaseGateway interface {
	Save(indicator main_domains.Indicator) (main_domains.Indicator, error)
	FindById(id string) (main_domains.Indicator, error)
}
