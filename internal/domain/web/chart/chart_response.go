package webchart

type ChartResponse struct {
	Id         int     `json:"id"`
	Qty        int     `json:"qty"`
	Total      float64 `json:"total"`
	CustomerId int     `json:"customer_id"`
	ProductId  int     `json:"product_id"`
}
