package main_domains

type IndicatorProcessorTrigger struct {
	ExecutionId   string
	IndicatorType IndicatorType
}

func NewIndicatorProcessorTrigger(indicatorType IndicatorType) IndicatorProcessorTrigger {
	return IndicatorProcessorTrigger{ExecutionId: "ID", IndicatorType: indicatorType}
}
