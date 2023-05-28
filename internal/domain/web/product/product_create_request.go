package webproduct

type ProductCreateRequest struct {
	Name         string                      `json:"name"`
	Price        int                         `json:"price"`
	Stock        int                         `json:"stock"`
	Description  string                      `json:"description"`
	ShippingCost int                         `json:"shipping_cost"`
	MerkId       int                         `json:"merk_id"`
	StoreId      int                         `json:"store_id"`
	Image        []ProductImageCreateRequest `json:"image"`
}

type ProductImageCreateRequest struct {
	ImageUrl  string `json:"image_url"`
	ProductId int    `json:"product_id"`
}
