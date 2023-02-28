package main_domains

import (
	"time"
)

type Indicator struct {
	Id               string
	SellerId         string
	Type             IndicatorType
	CreatedDate      time.Time
	LastModifiedDate time.Time
	Value            float64
}
