package main_domains

type OrderPage struct {
	Items      []SellerOrder
	Page       int
	Size       int
	TotalPages int
	Total      int
}
