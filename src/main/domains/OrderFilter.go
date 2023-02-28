package main_domains

type OrderFilter struct {
	OrderBeginDate  string
	OrderEndDate    string
	InvoiceStatus   string
	Page            int64
	PageSize        int64
	SellerId        string
	OrderStatusList []string
}
