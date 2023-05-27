package webtransaction

type TransactionResponse struct {
	CustomerName   string  `json:"customer_name"`
	MerkName       string  `json:"merk_name"`
	ProductName    string  `json:"product_name"`
	ProductPrice   int     `json:"product_price"`
	ShippingCost   int     `json:"shipping_cost"`
	Qty            int     `json:"quantity"`
	Tax            float64 `json:"tax"`
	TotalPrice     int     `json:"total_price"`
	BuyDate        string  `json:"buy_date"`
	Status         string  `json:"Status"`
	StoreName      string  `json:"store_name"`
	VirtualAccount int     `json:"virtual_account"`
}

type ProductResponse struct {
	MerkName     string  `json:"merk_name"`
	ProductName  string  `json:"product_name"`
	ProductPrice int     `json:"product_price"`
	Qty          int     `json:"quantity"`
	Total        float64 `json:"total"`
}

type TransactionMultiplerResponse struct {
	CustomerName   string            `json:"customer_name"`
	Product        []ProductResponse `json:"product"`
	ShippingCost   int               `json:"shipping_cost"`
	Tax            float64           `json:"tax"`
	TotalPrice     int               `json:"total_price"`
	BuyDate        string            `json:"buy_date"`
	Status         string            `json:"Status"`
	StoreName      string            `json:"store_name"`
	VirtualAccount int               `json:"virtual_account"`
}

type CustomerPaymentResponse struct {
	CustomerName   string            `json:"customer_name"`
	Product        []ProductResponse `json:"product"`
	ShippingCost   int               `json:"shipping_cost"`
	Tax            float64           `json:"tax"`
	TotalPrice     float64           `json:"total_price"`
	BuyDate        string            `json:"buy_date"`
	Status         string            `json:"status"`
	StoreName      string            `json:"store_name"`
	VirtualAccount int               `json:"virtual_account"`
	Photo          string            `json:"photo"`
}
