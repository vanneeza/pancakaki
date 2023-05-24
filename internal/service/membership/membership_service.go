package membershipservice

import (
	entity "pancakaki/internal/domain/entity/membership"
	membershiprepository "pancakaki/internal/repository/membership"
)

type MembershipService interface {
	GetMembershipById(id int) (*entity.Membership, error)
}

type membershipService struct {
	membershipRepo membershiprepository.MembershipRepository
}

func (s *membershipService) GetMembershipById(id int) (*entity.Membership, error) {
	return s.membershipRepo.GetMembershipById(id)
}

func NewMembershipService(membershipRepo membershiprepository.MembershipRepository) MembershipService {
	return &membershipService{membershipRepo: membershipRepo}
}
