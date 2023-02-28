package main_gateways_redis_documents

import (
	main_domains "mpindicator/main/domains"
	"time"
)

type CachedIndicatorDocument struct {
	Id               string    `json:"id,omitempty"`
	SellerId         string    `json:"sellerId,omitempty"`
	Type             string    `json:"type,omitempty"`
	CreatedDate      time.Time `json:"createdDate,omitempty"`
	LastModifiedDate time.Time `json:"lastModifiedDate,omitempty"`
	Value            float64   `json:"value,omitempty"`
}

func NewCachedIndicatorDocument(indicator main_domains.Indicator) CachedIndicatorDocument {
	return CachedIndicatorDocument{
		Id:               indicator.Id,
		SellerId:         indicator.SellerId,
		Type:             indicator.Type.GetDescription(),
		CreatedDate:      indicator.CreatedDate,
		LastModifiedDate: indicator.LastModifiedDate,
		Value:            indicator.Value,
	}
}

func (thisDocument *CachedIndicatorDocument) ToDomain() main_domains.Indicator {
	return main_domains.Indicator{
		Id:               thisDocument.Id,
		SellerId:         thisDocument.SellerId,
		Type:             main_domains.IndicatorType(thisDocument.Type),
		CreatedDate:      thisDocument.CreatedDate,
		LastModifiedDate: thisDocument.LastModifiedDate,
		Value:            thisDocument.Value,
	}
}
