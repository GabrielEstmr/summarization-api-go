package main_usecases

import (
	"context"
	"log"
	main_domains "mpindicator/main/domains"
	main_gateways "mpindicator/main/gateways"
	main_utils "mpindicator/main/utils"
)

type SaveIndicatorValues struct {
	IndicatorDatabaseGateway main_gateways.IndicatorDatabaseGateway
}

func NewSaveIndicatorValues(
	indicatorDatabaseGateway *main_gateways.IndicatorDatabaseGateway) *SaveIndicatorValues {
	return &SaveIndicatorValues{IndicatorDatabaseGateway: *indicatorDatabaseGateway}
}

func (thisUseCase *SaveIndicatorValues) Execute(
	ctx context.Context,
	indicatorType main_domains.IndicatorType,
	sellerValues map[string]float64) {

	var indicators []main_domains.Indicator
	for key, value := range sellerValues {
		indicators = append(indicators, main_domains.Indicator{
			SellerId: key,
			Type:     indicatorType,
			Value:    value,
		})
	}

	ids, err := thisUseCase.IndicatorDatabaseGateway.SaveAll(ctx, indicators)
	main_utils.FailOnError(err, "MSG")
	log.Println(ids[len(ids)-1])

}
