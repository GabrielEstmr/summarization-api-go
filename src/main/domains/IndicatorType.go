package main_domains

import "errors"

type IndicatorType string

const INVALID_INDICATOR_DESCRIPTION = "Invalid Indicator description"

const (
	INDICATOR_TYPE_INVALID_INVOICES        IndicatorType = "invalid_invoice"
	INDICATOR_TYPE_ORDER_30_DAY_DISPATCH   IndicatorType = "order_30_day_dispatch"
	INDICATOR_TYPE_ORDER_30_DAY_INVOICE    IndicatorType = "order_30_day_invoice"
	INDICATOR_TYPE_ORDER_30_DAY_IN_TRANSIT IndicatorType = "order_30_day_in_transit"
)

func (s IndicatorType) GetDescription() string {
	switch s {
	case INDICATOR_TYPE_INVALID_INVOICES:
		return "invalid_invoice"
	case INDICATOR_TYPE_ORDER_30_DAY_DISPATCH:
		return "order_30_day_dispatch"
	case INDICATOR_TYPE_ORDER_30_DAY_INVOICE:
		return "order_30_day_invoice"
	case INDICATOR_TYPE_ORDER_30_DAY_IN_TRANSIT:
		return "order_30_day_in_transit"
	}
	return "unknown"
}

func FindIndicatorTypeByDescription(description string) (IndicatorType, error) {
	switch description {
	case "invalid_invoice":
		return INDICATOR_TYPE_INVALID_INVOICES, nil
	case "order_30_day_dispatch":
		return INDICATOR_TYPE_ORDER_30_DAY_DISPATCH, nil
	case "order_30_day_invoice":
		return INDICATOR_TYPE_ORDER_30_DAY_INVOICE, nil
	case "order_30_day_in_transit":
		return INDICATOR_TYPE_ORDER_30_DAY_IN_TRANSIT, nil
	}
	return "", errors.New(INVALID_INDICATOR_DESCRIPTION)
}
