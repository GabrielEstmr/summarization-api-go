package main_gateway_mongodb_documents

import (
	main_domains "mpindicator/main/domains"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IndicatorDocument struct {
	Id               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SellerId         string             `json:"sellerId,omitempty"`
	Type             string             `json:"type,omitempty"`
	CreatedDate      primitive.DateTime `json:"createdDate,omitempty"`
	LastModifiedDate primitive.DateTime `json:"lastModifiedDate,omitempty"`
	Value            float64            `json:"value,omitempty"`
}

func NewIndicatorDocument(indicator main_domains.Indicator) IndicatorDocument {
	return IndicatorDocument{
		SellerId: indicator.SellerId,
		Type:     indicator.Type.GetDescription(),
		Value:    indicator.Value,
	}
}

func (thisDocument *IndicatorDocument) ToDomain() main_domains.Indicator {
	return main_domains.Indicator{
		Id:               thisDocument.Id.Hex(),
		SellerId:         thisDocument.SellerId,
		Type:             main_domains.IndicatorType(thisDocument.Type),
		CreatedDate:      thisDocument.CreatedDate.Time(),
		LastModifiedDate: thisDocument.LastModifiedDate.Time(),
		Value:            thisDocument.Value,
	}
}
