package adminservice

import (
	"fmt"
	"log"
	"pancakaki/internal/domain/entity"
	webadmin "pancakaki/internal/domain/web/admin"
	webbank "pancakaki/internal/domain/web/bank"
	webcustomer "pancakaki/internal/domain/web/customer"
	webowner "pancakaki/internal/domain/web/owner"
	adminrepository "pancakaki/internal/repository/admin"
	bankrepository "pancakaki/internal/repository/bank"
	customerrepository "pancakaki/internal/repository/customer"
	ownerrepository "pancakaki/internal/repository/owner"
	"pancakaki/utils/helper"
)

type AdminServiceImpl struct {
	AdminRepository    adminrepository.AdminRepository
	BankRepository     bankrepository.BankRepository
	OwnerRepository    ownerrepository.OwnerRepository
	CustomerRepository customerrepository.CustomerRepository
}

func NewAdminService(adminRepository adminrepository.AdminRepository, bankRepository bankrepository.BankRepository, ownerRepository ownerrepository.OwnerRepository, customerRepository customerrepository.CustomerRepository) AdminService {
	return &AdminServiceImpl{
		AdminRepository:    adminRepository,
		BankRepository:     bankRepository,
		OwnerRepository:    ownerRepository,
		CustomerRepository: customerRepository,
	}
}

func (adminService *AdminServiceImpl) Register(req webadmin.AdminCreateRequest) (webadmin.AdminResponse, error) {

	admin := entity.Admin{
		Username: req.Username,
		Password: req.Passowrd,
	}

	adminData, _ := adminService.AdminRepository.Create(&admin)

	adminResponse := webadmin.AdminResponse{
		Id:       adminData.Id,
		Username: adminData.Username,
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
			Username: admin.Username,
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
		Username: admin.Username,
		Password: admin.Password,
	}

	return adminResponse, nil
}

func (adminService *AdminServiceImpl) Edit(req webadmin.AdminUpdateRequest) (webadmin.AdminResponse, error) {

	admin := entity.Admin{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
	}

	adminData, err := adminService.AdminRepository.Update(&admin)
	helper.PanicErr(err)

	adminResponse := webadmin.AdminResponse{
		Id:       adminData.Id,
		Username: adminData.Username,
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
		Username: adminData.Username,
		Password: adminData.Password,
	}

	return adminResponse, nil
}

func (adminService *AdminServiceImpl) RegisterBank(req webbank.BankCreateRequest, reqBank webbank.BankAdminCreateRequest) (webbank.BankResponse, error) {

	bank := entity.Bank{
		Name:        req.Name,
		BankAccount: req.BankAccount,
		AccountName: req.AccountName,
	}

	log.Println(reqBank.AdminId, "adakah")
	log.Println(bank)
	fmt.Scanln()
	bankData, _ := adminService.BankRepository.Create(&bank)

	bankAdmin := entity.BankAdmin{
		AdminId: reqBank.AdminId,
		BankId:  bankData.Id,
	}

	log.Println(bankData, "Nill kah?")
	fmt.Scanln()
	adminService.BankRepository.CreateBankAdmin(&bankAdmin)

	bankResponse := webbank.BankResponse{
		Id:          bankData.Id,
		Name:        bankData.Name,
		AccountName: bankData.AccountName,
		BankAccount: bankData.BankAccount,
	}
	return bankResponse, nil
}

func (adminService *AdminServiceImpl) ViewOneBank(name string) (webbank.BankResponse, error) {

	bankData, _ := adminService.BankRepository.FindByName(name)

	bankResponse := webbank.BankResponse{
		Id:          bankData.Id,
		Name:        bankData.Name,
		AccountName: bankData.AccountName,
		BankAccount: bankData.BankAccount,
	}
	return bankResponse, nil
}

func (adminService *AdminServiceImpl) EditBank(req webbank.BankUpdateRequest, reqBank webbank.BankAdminUpdateRequest) (webbank.BankResponse, error) {

	bank := entity.Bank{
		Id:          reqBank.AdminId,
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

func (adminService *AdminServiceImpl) ViewTransactionAllOwner() ([]webadmin.TransactionOwnerResponse, error) {

	adminData, err := adminService.AdminRepository.FindTransactionAllOwner()
	helper.PanicErr(err)

	adminResponse := make([]webadmin.TransactionOwnerResponse, len(adminData))
	for i, admin := range adminData {
		adminResponse[i] = webadmin.TransactionOwnerResponse{
			OwnerName:    admin.OwnerName,
			ProductName:  admin.NameProduct,
			MerkName:     admin.NameMerk,
			Price:        admin.Price,
			Qty:          admin.Qty,
			BuyDate:      admin.BuyDate,
			TotalPrice:   admin.TotalPrice,
			Status:       admin.Status,
			CustomerName: admin.CustomerName,
		}
	}
	return adminResponse, nil
}

func (adminService *AdminServiceImpl) ViewTransactionOwnerByName(ownerNmae string) (webadmin.TransactionOwnerResponse, error) {
	transactionOwner, err := adminService.AdminRepository.FindTransactionOwnerByName(ownerNmae)
	helper.PanicErr(err)

	adminResponse := webadmin.TransactionOwnerResponse{
		OwnerName:    transactionOwner.OwnerName,
		ProductName:  transactionOwner.NameProduct,
		MerkName:     transactionOwner.NameMerk,
		Price:        transactionOwner.Price,
		Qty:          transactionOwner.Qty,
		BuyDate:      transactionOwner.BuyDate,
		TotalPrice:   transactionOwner.TotalPrice,
		Status:       transactionOwner.Status,
		CustomerName: transactionOwner.CustomerName,
	}

	return adminResponse, nil
}

func (adminService *AdminServiceImpl) ViewAllOwner() ([]webadmin.FindOwnerResponse, error) {

	ownerData, err := adminService.AdminRepository.FindOwner()
	helper.PanicErr(err)

	adminResponse := make([]webadmin.FindOwnerResponse, len(ownerData))
	for i, owner := range ownerData {
		adminResponse[i] = webadmin.FindOwnerResponse{
			OwnerName:      owner.OwnerName,
			NoHp:           owner.NoHp,
			Email:          owner.Email,
			Password:       owner.Password,
			NameMembership: owner.NameMembership,
			NameStore:      owner.NameStore,
		}
	}
	return adminResponse, nil
}

func (adminService *AdminServiceImpl) ViewOwnerByName(ownerName string) (webadmin.FindOwnerResponse, error) {
	ownerData, err := adminService.AdminRepository.FindOwnerByName(ownerName)
	helper.PanicErr(err)

	ownerResponse := webadmin.FindOwnerResponse{
		OwnerName:      ownerData.OwnerName,
		NoHp:           ownerData.NoHp,
		Email:          ownerData.Email,
		Password:       ownerData.Password,
		NameMembership: ownerData.NameMembership,
		NameStore:      ownerData.NameStore,
	}

	return ownerResponse, nil
}

func (adminService *AdminServiceImpl) UnregOwner(ownerId int) (webowner.OwnerResponse, error) {
	ownerData, err := adminService.OwnerRepository.GetOwnerById(ownerId)
	helper.PanicErr(err)
	err = adminService.OwnerRepository.DeleteOwner(ownerId)
	helper.PanicErr(err)

	ownerResponse := webowner.OwnerResponse{
		Id:           ownerData.Id,
		Name:         ownerData.Name,
		NoHp:         ownerData.NoHp,
		Email:        ownerData.Email,
		Password:     ownerData.Password,
		MembershipId: ownerData.MembershipId,
	}
	return ownerResponse, nil
}

func (adminService *AdminServiceImpl) ViewTransactionCustomerById(customerId int) ([]webcustomer.TransactionCustomer, error) {
	transactionCustomer, err := adminService.CustomerRepository.FindTransactionCustomerById(customerId)
	helper.PanicErr(err)

	transactionCustomerResponse := make([]webcustomer.TransactionCustomer, len(transactionCustomer))
	for i, txCustomer := range transactionCustomer {
		transactionCustomerResponse[i] = webcustomer.TransactionCustomer{
			OwnerName:    txCustomer.OwnerName,
			ProductName:  txCustomer.NameProduct,
			MerkName:     txCustomer.NameMerk,
			Price:        txCustomer.Price,
			Qty:          txCustomer.Qty,
			BuyDate:      txCustomer.BuyDate,
			TotalPrice:   txCustomer.TotalPrice,
			Status:       txCustomer.Status,
			CustomerName: txCustomer.CustomerName,
		}
	}
	return transactionCustomerResponse, nil
}
