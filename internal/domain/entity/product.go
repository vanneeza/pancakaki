package entity

type Product struct {
	Id           int
	Name         string
	Price        int
	Stock        int
	Description  string
	ShippingCost int
	MerkId       int
	StoreId      int
	IsDeleted    bool
}
