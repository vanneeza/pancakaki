package adminservice

import webadmin "pancakaki/internal/domain/web/admin"

type AdminService interface {
	Register(req webadmin.AdminCreateRequest) (webadmin.AdminResponse, error)
	ViewAll() ([]webadmin.AdminResponse, error)
	ViewOne(adminId int) (webadmin.AdminResponse, error)
	Edit(req webadmin.AdminUpdateRequest) (webadmin.AdminResponse, error)
	Unreg(adminId int) (webadmin.AdminResponse, error)
}
