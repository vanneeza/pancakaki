package membershipservice

import webmembership "pancakaki/internal/domain/web/membership"

type MembershipService interface {
	Register(req webmembership.MembershipCreateRequest) (webmembership.MembershipResponse, error)
	ViewAll() ([]webmembership.MembershipResponse, error)
	ViewOne(memberwebmembershipId int) (webmembership.MembershipResponse, error)
	Edit(req webmembership.MembershipUpdateRequest) (webmembership.MembershipResponse, error)
	Unreg(memberwebmembershipId int) (webmembership.MembershipResponse, error)
}
