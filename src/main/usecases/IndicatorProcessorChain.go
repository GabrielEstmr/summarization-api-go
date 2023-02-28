package main_usecases

import (
	"context"
	main_domains "mpindicator/main/domains"
)

type IndicatorProcessorChain struct {
	indicatorProcessor IndicatorProcessor
}

func NewIndicatorProcessorChain(indicatorProcessor *IndicatorProcessor) *IndicatorProcessorChain {
	return &IndicatorProcessorChain{indicatorProcessor: *indicatorProcessor}
}

func (thisUseCase *IndicatorProcessorChain) Execute(
	ctx context.Context,
	indicatorProcessorTrigger main_domains.IndicatorProcessorTrigger) {
	thisUseCase.indicatorProcessor.Execute(ctx, indicatorProcessorTrigger)
}
