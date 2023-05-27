package webchart

type ChartCreateRequest struct {
	Qty        int `json:"qty"`
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
}
