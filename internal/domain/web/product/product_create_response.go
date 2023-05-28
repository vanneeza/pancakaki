package webproduct

import "pancakaki/internal/domain/entity"

type ProductCreateResponse struct {
	Id           int                   `json:"id"`
	Name         string                `json:"name"`
	Price        int                   `json:"price"`
	Stock        int                   `json:"stock"`
	Description  string                `json:"description"`
	ShippingCost int                   `json:"shipping_cost"`
	MerkId       int                   `json:"merk_id"`
	StoreId      int                   `json:"store_id"`
	Image        []entity.ProductImage `json:"image"`
}

type ProductImageCreateResponse struct {
	Id        int    `json:"id"`
	ImageUrl  string `json:"image_url"`
	ProductId int    `json:"product_id"`
}
