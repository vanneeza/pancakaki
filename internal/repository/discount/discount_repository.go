package discountrepository

import entity "pancakaki/internal/domain/entity/discount"

type DiscountRepository interface {
	Create(discount *entity.Discount) (*entity.Discount, error)
	FindAll() ([]entity.Discount, error)
	FindById(id int) (*entity.Discount, error)
	Update(discount *entity.Discount) (*entity.Discount, error)
	Delete(discountId int) error
}
