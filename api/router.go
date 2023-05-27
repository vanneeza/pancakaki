package api

import (
	"database/sql"
	admincontroller "pancakaki/api/controller/admin"
	chartcontroller "pancakaki/api/controller/chart"
	customercontroller "pancakaki/api/controller/customer"
	membershipcontroller "pancakaki/api/controller/membership"
	merkcontroller "pancakaki/api/controller/merk"
	ownercontroller "pancakaki/api/controller/owner"
	productcontroller "pancakaki/api/controller/product"
	productimagecontroller "pancakaki/api/controller/product_image"
	storecontroller "pancakaki/api/controller/store"
	transactioncontroller "pancakaki/api/controller/transaction"
	adminrepository "pancakaki/internal/repository/admin"
	bankrepository "pancakaki/internal/repository/bank"
	bankstorerepository "pancakaki/internal/repository/bank_store"
	chartrepository "pancakaki/internal/repository/chart"
	customerrepository "pancakaki/internal/repository/customer"
	membershiprepository "pancakaki/internal/repository/membership"
	merkrepository "pancakaki/internal/repository/merk"
	ownerrepository "pancakaki/internal/repository/owner"
	productrepository "pancakaki/internal/repository/product"
	productimagerepository "pancakaki/internal/repository/product_image"
	storerepository "pancakaki/internal/repository/store"
	transactionrepository "pancakaki/internal/repository/transaction"
	adminservice "pancakaki/internal/service/admin"
	bankservice "pancakaki/internal/service/bank"
	chartservice "pancakaki/internal/service/chart"
	customerservice "pancakaki/internal/service/customer"
	membershipservice "pancakaki/internal/service/membership"
	merkservice "pancakaki/internal/service/merk"
	ownerservice "pancakaki/internal/service/owner"
	productservice "pancakaki/internal/service/product"
	productimageservice "pancakaki/internal/service/product_image"
	storeservice "pancakaki/internal/service/store"
	transactionservice "pancakaki/internal/service/transaction"
	"pancakaki/utils/helper"

	"github.com/gin-gonic/gin"
)

func Run(db *sql.DB) *gin.Engine {
	r := gin.Default()

	merkRepository := merkrepository.NewMerkRepository(db)
	merkService := merkservice.NewMerkService(merkRepository)
	merkcontroller := merkcontroller.NewMerkController(merkService)

	productImageRepository := productimagerepository.NewProductImageRepository(db)
	productImageService := productimageservice.NewProductImageService(productImageRepository)
	productImageController := productimagecontroller.NewProductImageHandler(productImageService)

	membershipRepository := membershiprepository.NewMembershipRepository(db)
	membershipService := membershipservice.NewMembershipService(membershipRepository)
	// membershipController := membershipcontroller.NewMembershipHandler(membershipService)

	bankRepository := bankstorerepository.NewBankStoreRepository(db) ///////////// mas ady
	bankService := bankservice.NewBankService(bankRepository)
	bankStoreRepository := bankstorerepository.NewBankStoreRepository(db)

	ownerRepository := ownerrepository.NewOwnerRepository(db)
	ownerService := ownerservice.NewOwnerService(ownerRepository)
	ownerController := ownercontroller.NewOwnerHandler(ownerService, membershipService, bankService)

	storeRepository := storerepository.NewStoreRepository(db, bankStoreRepository)
	storeService := storeservice.NewStoreService(storeRepository)
	storeController := storecontroller.NewStoreHandler(storeService, ownerService)

	productRepository := productrepository.NewProductRepository(db)
	productService := productservice.NewProductService(productRepository)
	productController := productcontroller.NewProductHandler(productService, storeService)

	///////////-----------------------------------------------------------------------------------------------------------////////////////////

	customerRepository := customerrepository.NewCustomerRepository(db)
	bankRepoCha := bankrepository.NewBankRepository(db)

	membershipRepositoryCha := membershiprepository.NewMembershipRepository(db)
	membershipServiceCHa := membershipservice.NewMembershipService(membershipRepositoryCha)
	membershipController := membershipcontroller.NewMembershipController(membershipServiceCHa)

	adminRepository := adminrepository.NewAdminRepository(db)
	adminService := adminservice.NewAdminService(adminRepository, bankRepoCha, ownerRepository, customerRepository)
	adminController := admincontroller.NewAdminController(adminService)

	customerService := customerservice.NewCustomerService(customerRepository)
	customerController := customercontroller.NewCustomerController(customerService)

	chartRepository := chartrepository.NewChartRepository(db)
	chartService := chartservice.NewChartService(chartRepository, productRepository)
	chartController := chartcontroller.NewChartController(chartService)

	transactionRepositoryCha := transactionrepository.NewTransactionRepository(db)
	transactionServiceCHa := transactionservice.NewTransactionService(transactionRepositoryCha, productRepository, customerRepository, ownerRepository, chartRepository)
	transactionController := transactioncontroller.NewTransactionController(transactionServiceCHa)

	var jwtKey = "secret_key"
	pancakaki := r.Group("pancakaki/v1/")

	pancakaki.POST("/login", ownerController.LoginOwner)
	pancakaki.POST("/", ownerController.CreateOwner)

	pancakaki.GET("admins/", adminController.ViewAll)
	admin := pancakaki.Group("/admin")
	{
		admin.POST("/", adminController.Register)
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
		// admin.GET("/transaction_history/customer/:id", adminController.ViewTransactionCustomerById)

		admin.GET("/owner/profiles/", adminController.ViewAllOwner)
		admin.DELETE("/owner/profile/:id", adminController.UnregOwner)
		admin.GET("/owner/profile/:name", adminController.ViewOwnerByName)

		admin.GET("/customer/profiles/", customerController.ViewAll)
		admin.GET("/customer/profile/:name", customerController.ViewOne)

		admin.POST("/merk/", merkcontroller.Register)
		admin.GET("/merks/", merkcontroller.ViewAll)
		admin.GET("/merk/:id", merkcontroller.ViewOne)
		admin.PUT("/merk/:id", merkcontroller.Edit)
		admin.DELETE("/merk/:id", merkcontroller.Unreg)
	}

	owner := pancakaki.Group("/owner")
	owner.Use(helper.AuthMiddleware(jwtKey))
	{
		//owner
		owner.GET("/:ownername/profile", ownerController.GetOwnerById)
		owner.PUT("/:ownername/profile", ownerController.UpdateOwner)
		owner.PUT("/:ownername/profile/:id", ownerController.DeleteOwner)
		//store
		owner.POST("/:ownername/store", storeController.CreateMainStore)
		owner.POST("/:ownername/store/storename", storeController.UpdateMainStore)
		owner.POST("/:ownername/store/:storename/product", productController.InsertProduct)

	}

	merk := pancakaki.Group("/testaja")
	{
		store := pancakaki.Group("/stores")
		{
			store.POST("/product", merkcontroller.Register)
		}
		merk.POST("/merk", merkcontroller.Register)
		merk.GET("/", merkcontroller.ViewAll)
		merk.GET("/:id", merkcontroller.ViewOne)
		merk.PUT("/", merkcontroller.Unreg)
		merk.DELETE("/:id", merkcontroller.Unreg)
	}

	product := pancakaki.Group("/products")
	{
		product.POST("/", productController.InsertProduct)
		product.GET("/", productController.FindAllProduct)
		product.GET("/:id", productController.FindProductById)
		product.GET("/name/:name", productController.FindProductByName)
		product.PUT("/", productController.UpdateProduct)
		product.PUT("/:id", productController.DeleteProduct)
	}

	productImage := pancakaki.Group("/product-image")
	{
		productImage.POST("/:id", productImageController.InsertProductImage)
		productImage.GET("/", productImageController.FindAllProductImage)
		productImage.GET("/:id", productImageController.FindProductImageById)
		productImage.GET("/name/:name", productImageController.FindProductImageByName)
		productImage.PUT("/", productImageController.UpdateProductImage)
		productImage.PUT("/:id", productImageController.DeleteProductImage)
	}

	customer := pancakaki.Group("/customers")
	{
		customer.POST("/", customerController.Register)
		customer.GET("/", customerController.ViewAll)
		customer.GET("/:name", customerController.ViewOne)
		customer.PUT("/:id", customerController.Edit)
		customer.DELETE("/:name", customerController.Unreg)

		//------------ TRANSACTION CUSTOMER LANGSUNG BELI --------------------//
		customer.POST("/transaction", transactionController.MakeOrder)
		customer.POST("/transactions", transactionController.MakeMultipleOrder)
		customer.POST("/payment/:id", transactionController.CustomerPayment)

		//---------- Customer Chart ------------------- //
		customer.POST("/chart/", chartController.Register)
		customer.GET("/charts/:id", chartController.ViewAll)
		customer.GET("/chart/:id", chartController.ViewOne)
		customer.PUT("/chart/:id", chartController.Edit)
		customer.DELETE("/chart/:id", chartController.Unreg)

		//------- NOTIFICATION ------
		customer.GET("/notification/:id", customerController.Notification)
		////------------------- TEST PAYMENTGATEWAY : PROSESS --------------//
		// customer.POST("/payment", transactionController.CreatePaymentIntent)
	}

	return r
}
