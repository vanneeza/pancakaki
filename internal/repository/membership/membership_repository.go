package membershiprepository

import "pancakaki/internal/domain/entity"

type MembershipRepository interface {
	Create(membership *entity.Membership) (*entity.Membership, error)
	FindAll() ([]entity.Membership, error)
	FindById(id int) (*entity.Membership, error)
	Update(membership *entity.Membership) (*entity.Membership, error)
	Delete(membershipId int) error
}
