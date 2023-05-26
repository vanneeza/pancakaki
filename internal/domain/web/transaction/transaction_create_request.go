package webtransaction

import (
	"mime/multipart"
	"time"
)

type TransactionOrderDetailCreateRequest struct {
	BuyDate    time.Time `json:"buy_date"`
	Status     string    `json:"status"`
	TotalPrice int       `json:"total_price"`
	Photo      string    `json:"photo"`
	Tax        float64   `json:"tax"`
}

type TransactionOrderCreateRequest struct {
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
	Qty        int `json:"qty"`
}

type PaymentCreateRequest struct {
	Transaction_detail_order_Id int                   `form:"transaction_detail_order_id"`
	VirtualAccount              int                   `form:"virtual_account"`
	Pay                         int                   `form:"pay"`
	Photo                       *multipart.FileHeader `form:"photo"`
}
