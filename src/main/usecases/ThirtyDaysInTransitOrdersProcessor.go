package main_usecases

import (
	"context"
	"log"
	main_domains "mpindicator/main/domains"
)

type ThirtyDaysInTransitOrdersProcessor struct {
	getThirtyDaysInTransitOrders GetThirtyDaysInTransitOrders
	saveIndicatorValues          SaveIndicatorValues
}

func NewThirtyDaysInTransitOrdersProcessor(
	getThirtyDaysInTransitOrders *GetThirtyDaysInTransitOrders,
	saveIndicatorValues *SaveIndicatorValues,
) *ThirtyDaysInTransitOrdersProcessor {
	return &ThirtyDaysInTransitOrdersProcessor{
		getThirtyDaysInTransitOrders: *getThirtyDaysInTransitOrders,
		saveIndicatorValues:          *saveIndicatorValues,
	}
}

func (thisUseCase *ThirtyDaysInTransitOrdersProcessor) Execute(
	ctx context.Context,
	indicatorProcessorTrigger main_domains.IndicatorProcessorTrigger) {

	if canExecute(indicatorProcessorTrigger) {

		log.Println("INDICATOR ALLOWED")

		// irá printar os dois
		totalPages := thisUseCase.getThirtyDaysInTransitOrders.Execute(0).TotalPages
		log.Println(totalPages)
		var channels []<-chan main_domains.OrderPage

		for i := 0; i < totalPages; i++ {
			channels = append(channels, thisUseCase.createMessageLong(int64(i)))
		}

		var orderPages []main_domains.OrderPage
		channelUnique := JoinChannels(channels)
		for i := 0; i < totalPages; i++ {
			orderPages = append(orderPages, <-channelUnique)
		}

		indicatorValues := make(map[string]float64)
		for _, value := range orderPages {
			for _, valueItems := range value.Items {
				if indicatorValues[valueItems.SellerId] != 0 {
					indicatorValues[valueItems.SellerId] = indicatorValues[valueItems.SellerId] + 1
				} else {
					indicatorValues[valueItems.SellerId] = 1
				}
			}

		}

		thisUseCase.saveIndicatorValues.Execute(ctx, indicatorProcessorTrigger.IndicatorType, indicatorValues)

	} else {
		log.Println("INDICATOR NOT ALLOWED")
	}

}

func canExecute(indicatorProcessorTrigger main_domains.IndicatorProcessorTrigger) bool {
	return indicatorProcessorTrigger.IndicatorType == main_domains.INDICATOR_TYPE_ORDER_30_DAY_IN_TRANSIT
}

func (thisUseCase *ThirtyDaysInTransitOrdersProcessor) createMessageLong(page int64) <-chan main_domains.OrderPage {
	channel := make(chan main_domains.OrderPage)

	log.Println("ENTROU CANAL ======> PAGE", page)

	go func() {
		for {
			channel <- thisUseCase.getThirtyDaysInTransitOrders.Execute(page)
		}
	}()
	return channel
}

func JoinChannels(channels []<-chan main_domains.OrderPage) <-chan main_domains.OrderPage {
	resultChannel := make(chan main_domains.OrderPage)
	go func() {
		for {

			for _, value := range channels {
				select {
				case message := <-value:
					resultChannel <- message
				}
			}

		}
	}()
	return resultChannel
}
