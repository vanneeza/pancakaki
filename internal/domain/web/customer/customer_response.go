package webcustomer

type CustomerResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NoHp     int64  `json:"no_hp"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type TransactionCustomer struct {
	CustomerName   string  `json:"customer_name"`
	MerkName       string  `json:"merk_name"`
	ProductId      int     `json:"product_id"`
	ProductName    string  `json:"product_name"`
	ProductPrice   int     `json:"product_price"`
	ShippingCost   int     `json:"shipping_cost"`
	Qty            int     `json:"quantity"`
	Tax            float64 `json:"tax"`
	TotalPrice     float64 `json:"total_price"`
	BuyDate        string  `json:"buy_date"`
	Status         string  `json:"status"`
	StoreName      string  `json:"store_name"`
	VirtualAccount int     `json:"virtual_account"`
	Photo          string  `json:"photo"`
}
