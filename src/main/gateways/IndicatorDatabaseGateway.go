package main_gateways

import (
	"context"
	main_domains "mpindicator/main/domains"
)

type IndicatorDatabaseGateway interface {
	Save(context.Context, main_domains.Indicator) (string, error)
	SaveAll(context.Context, []main_domains.Indicator) ([]string, error)
	FindById(context.Context, string) (main_domains.Indicator, error)
}
