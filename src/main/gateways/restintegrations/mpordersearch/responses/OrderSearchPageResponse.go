package main_gateways_restintegrations_mpordersearch_responses

import main_domains "mpindicator/main/domains"

type OrderSearchPageResponse struct {
	Items      []OrderSearchSellerOrderResponse `json:"items,omitempty"`
	Page       int                              `json:"page,omitempty"`
	Size       int                              `json:"size,omitempty"`
	TotalPages int                              `json:"totalPages,omitempty"`
	Total      int                              `json:"total,omitempty"`
}

func (pageResponse *OrderSearchPageResponse) ToDomain() main_domains.OrderPage {

	var pageItemsSellerOrder []main_domains.SellerOrder
	pageItems := pageResponse.Items
	for _, value := range pageItems {
		pageItemsSellerOrder = append(pageItemsSellerOrder, value.ToDomain())
	}

	return main_domains.OrderPage{
		Items:      pageItemsSellerOrder,
		Page:       pageResponse.Page,
		Size:       pageResponse.Size,
		TotalPages: pageResponse.TotalPages,
		Total:      pageResponse.Total,
	}
}
