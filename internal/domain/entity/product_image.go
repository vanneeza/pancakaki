package entity

type ProductImage struct {
	Id        int    `json:"id"`
	ImageUrl  string `json:"image_url"`
	ProductId int    `json:"product_id"`
}
