package entity

type Product struct {
	Id           int
	Name         string
	Price        int
	Stock        int16
	Description  string
	ShippingCost int
	MerkId       int
	StoreId      int
	IsDeleted    bool
	// CreatedAt    time.Time
	// UpdateAt     time.Time
}
