package webproduct

type ProductUpdateRequest struct {
	Id           int                         `json:"product_id"`
	Name         string                      `json:"name"`
	Price        int                         `json:"price"`
	Stock        int                         `json:"stock"`
	Description  string                      `json:"description"`
	ShippingCost int                         `json:"shipping_cost"`
	MerkId       int                         `json:"merk_id"`
	StoreId      int                         `json:"store_id"`
	Image        []ProductImageUpdateRequest `json:"image"`
}

type ProductImageUpdateRequest struct {
	Id        int    `json:"product_image_id"`
	ImageUrl  string `json:"image_url"`
	ProductId int    `json:"product_id"`
}
