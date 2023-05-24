package api

import (
	"database/sql"
	"fmt"
	admincontroller "pancakaki/api/controller/admin"
	customercontroller "pancakaki/api/controller/customer"
	membershipcontroller "pancakaki/api/controller/membership"
	adminrepository "pancakaki/internal/repository/admin"
	bankrepository "pancakaki/internal/repository/bank"
	customerrepository "pancakaki/internal/repository/customer"
	membershiprepository "pancakaki/internal/repository/membership"
	adminservice "pancakaki/internal/service/admin"
	customerservice "pancakaki/internal/service/customer"
	membershipservice "pancakaki/internal/service/membership"

	"github.com/gin-gonic/gin"
)

func Run(db *sql.DB) *gin.Engine {
	r := gin.Default()

	bankRepository := bankrepository.NewBankRepository(db)
	adminRepository := adminrepository.NewAdminRepository(db)
	adminService := adminservice.NewAdminService(adminRepository, bankRepository)
	adminController := admincontroller.NewAdminController(adminService)

	customerRepository := customerrepository.NewCustomerRepository(db)
	customerService := customerservice.NewCustomerService(customerRepository)
	customerController := customercontroller.NewCustomerController(customerService)

	membershipRepository := membershiprepository.NewMembershipRepository(db)
	membershipService := membershipservice.NewMembershipService(membershipRepository)
	membershipController := membershipcontroller.NewMembershipController(membershipService)

	pancakaki := r.Group("pancakaki/v1")

	admin := pancakaki.Group("/admins")
	{
		admin.POST("/", adminController.Register)
		admin.GET("/", adminController.ViewAll)
		admin.GET("/:id", adminController.ViewOne)
		admin.PUT("/:id", adminController.Edit)
		admin.DELETE("/:id", adminController.Unreg)

		admin.POST("/bank/:id", adminController.RegisterBank)
		admin.PUT("/bank/:id", adminController.EditBank)
		admin.GET("/banks/", adminController.ViewAllBank)
		admin.GET("/bank/:name", adminController.ViewOneBank)

		admin.POST("/membership/", membershipController.Register)
		admin.GET("/memberships/", membershipController.ViewAll)
		admin.GET("/membership/:id", membershipController.ViewOne)
		admin.PUT("/membership/:id", membershipController.Edit)
		admin.DELETE("/membership/:id", membershipController.Unreg)

		admin.GET("/transaction_history/owners", adminController.ViewTransactionAllOwner)
		admin.GET("/transaction_history/owner/:name", adminController.ViewTransactionOwnerByName)

		admin.GET("/owner/profiles/", adminController.ViewAllOwner)
		admin.GET("/owner/profile/:name", adminController.ViewOwnerByName)

	}

	customer := pancakaki.Group("/customers")
	{
		customer.POST("/", customerController.Register)
		customer.GET("/", customerController.ViewAll)
		customer.GET("/:id", customerController.ViewOne)
		customer.PUT("/:id", customerController.Edit)
		customer.DELETE("/:id", customerController.Unreg)
	}

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})

		if err := db.Ping(); err != nil {
			c.JSON(500, gin.H{"message": "Koneksi database gagal"})
			return
		} else {
			c.JSON(200, gin.H{"message": "Koneksi database berhasillllllll"})
			return
		}
	})

	err := r.Run(":8000")
	if err != nil {
		panic("Gagal menjalankan server: " + err.Error())
	}

	r.Run(":8000")
	fmt.Println("Server berjalan di http://localhost:8000")

	return r
}
