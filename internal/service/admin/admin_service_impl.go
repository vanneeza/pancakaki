package adminservice

import (
	"log"
	entity "pancakaki/internal/domain/entity/admin"
	webadmin "pancakaki/internal/domain/web/admin"
	adminrepository "pancakaki/internal/repository/admin"
	"pancakaki/utils/helper"
)

type AdminServiceImpl struct {
	AdminRepository adminrepository.AdminRepository
}

func NewAdminService(adminRepository adminrepository.AdminRepository) AdminService {
	return &AdminServiceImpl{
		AdminRepository: adminRepository,
	}
}

func (adminService *AdminServiceImpl) Register(req webadmin.AdminCreateRequest) (webadmin.AdminResponse, error) {

	admin := entity.Admin{
		Name:     req.Name,
		Password: req.Passowrd,
	}

	adminData, _ := adminService.AdminRepository.Create(&admin)

	adminResponse := webadmin.AdminResponse{
		Id:       adminData.Id,
		Name:     adminData.Name,
		Password: adminData.Password,
	}
	return adminResponse, nil
}

func (adminService *AdminServiceImpl) ViewAll() ([]webadmin.AdminResponse, error) {

	adminData, err := adminService.AdminRepository.FindAll()
	helper.PanicErr(err)

	adminResponse := make([]webadmin.AdminResponse, len(adminData))
	for i, admin := range adminData {
		adminResponse[i] = webadmin.AdminResponse{
			Id:       admin.Id,
			Name:     admin.Name,
			Password: admin.Password,
		}
	}
	return adminResponse, nil
}

func (adminService *AdminServiceImpl) ViewOne(adminId int) (webadmin.AdminResponse, error) {
	admin, err := adminService.AdminRepository.FindById(adminId)
	helper.PanicErr(err)

	adminResponse := webadmin.AdminResponse{
		Id:       admin.Id,
		Name:     admin.Name,
		Password: admin.Password,
	}

	return adminResponse, nil
}

func (adminService *AdminServiceImpl) Edit(req webadmin.AdminUpdateRequest) (webadmin.AdminResponse, error) {
	log.Println(req, "di service")

	admin := entity.Admin{
		Id:       req.Id,
		Name:     req.Name,
		Password: req.Password,
	}
	log.Println(admin, "admin di service")

	adminData, err := adminService.AdminRepository.Update(&admin)
	helper.PanicErr(err)
	log.Println(admin, "admin Daa di service")

	adminResponse := webadmin.AdminResponse{
		Id:       adminData.Id,
		Name:     adminData.Name,
		Password: adminData.Password,
	}

	return adminResponse, nil
}

func (adminService *AdminServiceImpl) Unreg(adminId int) (webadmin.AdminResponse, error) {

	adminData, err := adminService.AdminRepository.FindById(adminId)
	helper.PanicErr(err)

	err = adminService.AdminRepository.Delete(adminId)
	helper.PanicErr(err)

	adminResponse := webadmin.AdminResponse{
		Id:       adminData.Id,
		Name:     adminData.Name,
		Password: adminData.Password,
	}

	return adminResponse, nil
}
