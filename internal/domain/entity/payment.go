package entity

type Payment struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Qty   int    `json:"qty"`
	Price int    `json:"price"`
}
