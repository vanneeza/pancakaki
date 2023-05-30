package membershipservice

import (
	"errors"
	"fmt"
	"log"
	"pancakaki/internal/domain/entity"
	webmembership "pancakaki/internal/domain/web/membership"

	membershiprepository "pancakaki/internal/repository/membership"
	"pancakaki/utils/helper"
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

	membership := entity.Membership{
		Name:  req.Name,
		Tax:   req.Tax,
		Price: req.Price,
	}

	membershipData, _ := membershipService.MembershipRepository.Create(&membership)

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
	helper.PanicErr(err)

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
	// helper.PanicErr(err)
	if err != nil {
		return webmembership.MembershipResponse{}, errors.New("membership not found")
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

	membership := entity.Membership{
		Id:    req.Id,
		Name:  req.Name,
		Tax:   req.Tax,
		Price: req.Price,
	}

	log.Println(membership, "service")
	fmt.Scanln()
	membershipData, err := membershipService.MembershipRepository.Update(&membership)
	helper.PanicErr(err)

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
	helper.PanicErr(err)

	err = membershipService.MembershipRepository.Delete(membershipId)
	helper.PanicErr(err)

	membershipResponse := webmembership.MembershipResponse{
		Id:    membershipData.Id,
		Name:  membershipData.Name,
		Tax:   membershipData.Tax,
		Price: membershipData.Price,
	}

	return membershipResponse, nil
}
