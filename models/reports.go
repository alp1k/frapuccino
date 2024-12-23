package models

type TotalSales struct {
	TotalSales int `json:"total_sales"`
}

type PopularItems struct {
	Items []OrderItem `json:"popular_items"`
}
