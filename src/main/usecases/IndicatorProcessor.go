package main_usecases

import (
	"context"
	main_domains "mpindicator/main/domains"
)

type IndicatorProcessor interface {
	Execute(context.Context, main_domains.IndicatorProcessorTrigger)
}
