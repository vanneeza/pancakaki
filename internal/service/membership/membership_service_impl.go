package membershipservice

import (
	"errors"
	"pancakaki/internal/domain/entity"
	webmembership "pancakaki/internal/domain/web/membership"

	membershiprepository "pancakaki/internal/repository/membership"
)

type MembershipServiceImpl struct {
	MembershipRepository membershiprepository.MembershipRepository
}

func NewMembershipService(membershipRepository membershiprepository.MembershipRepository) MembershipService {
	return &MembershipServiceImpl{
		MembershipRepository: membershipRepository,
	}
}

func (membershipService *MembershipServiceImpl) Register(req webmembership.MembershipCreateRequest) (webmembership.MembershipResponse, error) {

	if req.Name == "" {
		return webmembership.MembershipResponse{}, errors.New("name required")
	}

	if req.Price == 0 {
		return webmembership.MembershipResponse{}, errors.New("price required")
	}

	if req.Tax == 0 {
		return webmembership.MembershipResponse{}, errors.New("tax required")
	}

	membership := entity.Membership{
		Name:  req.Name,
		Tax:   req.Tax,
		Price: req.Price,
	}

	membershipData, err := membershipService.MembershipRepository.Create(&membership)
	if err != nil {
		return webmembership.MembershipResponse{}, errors.New("failed to create membership")
	}

	membershipResponse := webmembership.MembershipResponse{
		Id:    membershipData.Id,
		Name:  membershipData.Name,
		Tax:   membershipData.Tax,
		Price: membershipData.Price,
	}
	return membershipResponse, nil
}

func (membershipService *MembershipServiceImpl) ViewAll() ([]webmembership.MembershipResponse, error) {

	membershipData, err := membershipService.MembershipRepository.FindAll()
	if err != nil {
		return []webmembership.MembershipResponse{}, errors.New("membership data not found")
	}

	membershipResponse := make([]webmembership.MembershipResponse, len(membershipData))
	for i, membership := range membershipData {
		membershipResponse[i] = webmembership.MembershipResponse{
			Id:    membership.Id,
			Name:  membership.Name,
			Tax:   membership.Tax,
			Price: membership.Price,
		}
	}
	return membershipResponse, nil
}

func (membershipService *MembershipServiceImpl) ViewOne(membershipId int) (webmembership.MembershipResponse, error) {

	membership, err := membershipService.MembershipRepository.FindById(membershipId)
	if err != nil {
		return webmembership.MembershipResponse{}, errors.New("membership data not found")
	}

	membershipResponse := webmembership.MembershipResponse{
		Id:    membership.Id,
		Name:  membership.Name,
		Tax:   membership.Tax,
		Price: membership.Price,
	}

	return membershipResponse, nil
}

func (membershipService *MembershipServiceImpl) Edit(req webmembership.MembershipUpdateRequest) (webmembership.MembershipResponse, error) {

	if req.Name == "" {
		return webmembership.MembershipResponse{}, errors.New("name required")
	}

	if req.Price == 0 {
		return webmembership.MembershipResponse{}, errors.New("price required")
	}

	if req.Tax == 0 {
		return webmembership.MembershipResponse{}, errors.New("tax required")
	}
	membership := entity.Membership{
		Id:    req.Id,
		Name:  req.Name,
		Tax:   req.Tax,
		Price: req.Price,
	}

	membershipData, err := membershipService.MembershipRepository.Update(&membership)

	if err != nil {
		return webmembership.MembershipResponse{}, errors.New("update data membership failed")
	}

	membershipResponse := webmembership.MembershipResponse{
		Id:    membershipData.Id,
		Name:  membershipData.Name,
		Tax:   membershipData.Tax,
		Price: membershipData.Price,
	}

	return membershipResponse, nil
}

func (membershipService *MembershipServiceImpl) Unreg(membershipId int) (webmembership.MembershipResponse, error) {

	membershipData, err := membershipService.MembershipRepository.FindById(membershipId)
	if err != nil {
		return webmembership.MembershipResponse{}, errors.New("membership data not found")
	}

	err = membershipService.MembershipRepository.Delete(membershipId)
	if err != nil {
		return webmembership.MembershipResponse{}, errors.New("membership data not found")
	}

	membershipResponse := webmembership.MembershipResponse{
		Id:    membershipData.Id,
		Name:  membershipData.Name,
		Tax:   membershipData.Tax,
		Price: membershipData.Price,
	}

	return membershipResponse, nil
}
