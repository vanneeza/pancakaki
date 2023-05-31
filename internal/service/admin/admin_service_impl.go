package adminservice

import (
	"errors"
	"pancakaki/internal/domain/entity"
	webadmin "pancakaki/internal/domain/web/admin"
	webbank "pancakaki/internal/domain/web/bank"
	adminrepository "pancakaki/internal/repository/admin"
	bankrepository "pancakaki/internal/repository/bank"
	customerrepository "pancakaki/internal/repository/customer"
	ownerrepository "pancakaki/internal/repository/owner"
	"pancakaki/utils/helper"

	"golang.org/x/crypto/bcrypt"
)

type AdminServiceImpl struct {
	AdminRepository    adminrepository.AdminRepository
	BankRepository     bankrepository.BankRepository
	OwnerRepository    ownerrepository.OwnerRepository
	CustomerRepository customerrepository.CustomerRepository
}

func NewAdminService(adminRepository adminrepository.AdminRepository,
	bankRepository bankrepository.BankRepository,
	ownerRepository ownerrepository.OwnerRepository,
	customerRepository customerrepository.CustomerRepository) AdminService {
	return &AdminServiceImpl{
		AdminRepository:    adminRepository,
		BankRepository:     bankRepository,
		OwnerRepository:    ownerRepository,
		CustomerRepository: customerRepository,
	}
}

func (adminService *AdminServiceImpl) Register(req webadmin.AdminCreateRequest) (webadmin.AdminResponse, error) {
	if req.Username == "" {
		return webadmin.AdminResponse{}, errors.New("username required")
	}

	if req.Password == "" {
		return webadmin.AdminResponse{}, errors.New("password required")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return webadmin.AdminResponse{}, errors.New("failed to generete password")
	}

	admin := entity.Admin{
		Username: req.Username,
		Password: string(encryptedPassword),
	}

	adminData, _ := adminService.AdminRepository.Create(&admin)
	adminResponse := webadmin.AdminResponse{
		Id:       adminData.Id,
		Username: adminData.Username,
		Password: adminData.Password,
		Role:     "admin",
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
			Username: admin.Username,
			Password: admin.Password,
			Role:     admin.Role,
			Token:    "NULL",
		}
	}
	return adminResponse, nil
}

func (adminService *AdminServiceImpl) ViewOne(adminId int, username string) (webadmin.AdminResponse, error) {
	admin, err := adminService.AdminRepository.FindById(adminId, username)
	if err != nil {
		return webadmin.AdminResponse{}, errors.New("NULL")
	}

	adminResponse := webadmin.AdminResponse{
		Id:       admin.Id,
		Username: admin.Username,
		Password: admin.Password,
		Role:     admin.Role,
		Token:    "NULL",
	}

	return adminResponse, nil
}

func (adminService *AdminServiceImpl) Edit(req webadmin.AdminUpdateRequest) (webadmin.AdminResponse, error) {
	if req.Username == "" {
		return webadmin.AdminResponse{}, errors.New("username required")
	}

	if req.Password == "" {
		return webadmin.AdminResponse{}, errors.New("password required")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return webadmin.AdminResponse{}, errors.New("failed to generete password")
	}

	admin := entity.Admin{
		Id:       req.Id,
		Username: req.Username,
		Password: string(encryptedPassword),
	}

	adminData, err := adminService.AdminRepository.Update(&admin)
	if err != nil {
		return webadmin.AdminResponse{}, errors.New("failed to update data")
	}

	adminResponse := webadmin.AdminResponse{
		Id:       adminData.Id,
		Username: adminData.Username,
		Password: adminData.Password,
		Role:     "admin",
		Token:    "NULL",
	}

	return adminResponse, nil
}

func (adminService *AdminServiceImpl) Unreg(adminId int, username string) (webadmin.AdminResponse, error) {
	if adminId == 0 && username == "" {
		return webadmin.AdminResponse{}, errors.New("admin id and username required")
	}
	adminData, err := adminService.AdminRepository.FindById(adminId, username)
	if err != nil {
		return webadmin.AdminResponse{}, errors.New("NULL")
	}

	err = adminService.AdminRepository.Delete(adminId)
	if err != nil {
		return webadmin.AdminResponse{}, errors.New("NULL")
	}

	adminResponse := webadmin.AdminResponse{
		Id:       adminData.Id,
		Username: adminData.Username,
		Password: adminData.Password,
		Role:     adminData.Role,
		Token:    "NULL",
	}

	return adminResponse, nil
}

func (adminService *AdminServiceImpl) RegisterBank(req webbank.BankCreateRequest, reqBank webbank.BankAdminCreateRequest) (webbank.BankResponse, error) {
	if req.AccountName == "" {
		return webbank.BankResponse{}, errors.New("account name is required")
	}

	if req.BankAccount == 0 {
		return webbank.BankResponse{}, errors.New("bank account is required")
	}

	bank := entity.Bank{
		Name:        req.Name,
		BankAccount: req.BankAccount,
		AccountName: req.AccountName,
	}

	bankData, err := adminService.BankRepository.Create(&bank)

	if err != nil {
		return webbank.BankResponse{}, errors.New("failed to create bank")
	}
	bankAdmin := entity.BankAdmin{
		AdminId: reqBank.AdminId,
		BankId:  bankData.Id,
	}

	_, err2 := adminService.BankRepository.CreateBankAdmin(&bankAdmin)
	if err2 != nil {
		return webbank.BankResponse{}, errors.New("failed to create bank")
	}
	bankResponse := webbank.BankResponse{
		Id:          bankData.Id,
		Name:        bankData.Name,
		AccountName: bankData.AccountName,
		BankAccount: bankData.BankAccount,
	}
	return bankResponse, nil
}

func (adminService *AdminServiceImpl) EditBank(req webbank.BankUpdateRequest) (webbank.BankResponse, error) {

	bank := entity.Bank{
		Id:          req.Id,
		Name:        req.Name,
		BankAccount: req.BankAccount,
		AccountName: req.AccountName,
	}

	bankData, _ := adminService.BankRepository.Update(&bank)

	bankResponse := webbank.BankResponse{
		Id:          bankData.Id,
		Name:        bankData.Name,
		AccountName: bankData.AccountName,
		BankAccount: bankData.BankAccount,
	}
	return bankResponse, nil
}
func (adminService *AdminServiceImpl) ViewAllBank() ([]webbank.BankResponse, error) {

	bankData, err := adminService.BankRepository.FindAll()
	helper.PanicErr(err)

	bankResponse := make([]webbank.BankResponse, len(bankData))
	for i, bank := range bankData {
		bankResponse[i] = webbank.BankResponse{
			Id:          bank.Id,
			Name:        bank.Name,
			AccountName: bank.AccountName,
			BankAccount: bank.BankAccount,
		}
	}
	return bankResponse, nil
}

func (adminService *AdminServiceImpl) DeleteBank(bankId int) ([]webbank.BankResponse, error) {
	if bankId == 0 {
		return []webbank.BankResponse{}, errors.New("NULL")
	}

	bankResponse, err := adminService.BankRepository.FindById(bankId)
	if err != nil {
		return []webbank.BankResponse{}, errors.New("NULL")
	}

	err = adminService.BankRepository.Delete(bankId)

	if err != nil {
		return []webbank.BankResponse{}, errors.New("NULL")
	}

	bankResponses := make([]webbank.BankResponse, len(bankResponse))
	for i, bank := range bankResponse {
		bankResponses[i] = webbank.BankResponse{
			Id:          bank.Id,
			Name:        bank.Name,
			AccountName: bank.AccountName,
			BankAccount: bank.BankAccount,
		}

	}
	return bankResponses, nil
}
