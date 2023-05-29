package api

import (
	"database/sql"
	admincontroller "pancakaki/api/controller/admin"
	chartcontroller "pancakaki/api/controller/chart"
	customercontroller "pancakaki/api/controller/customer"
	logincontroller "pancakaki/api/controller/login"
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

	customerRepository := customerrepository.NewCustomerRepository(db)
	bankRepoCha := bankrepository.NewBankRepository(db)

	customerService := customerservice.NewCustomerService(customerRepository)
	customerController := customercontroller.NewCustomerController(customerService)

	ownerRepository := ownerrepository.NewOwnerRepository(db)
	ownerService := ownerservice.NewOwnerService(ownerRepository, customerRepository)
	ownerController := ownercontroller.NewOwnerHandler(ownerService, membershipService, bankService)

	productRepository := productrepository.NewProductRepository(db, productImageRepository)
	storeRepository := storerepository.NewStoreRepository(db, bankStoreRepository, productRepository)
	storeService := storeservice.NewStoreService(storeRepository, bankRepository)
	storeController := storecontroller.NewStoreHandler(storeService, ownerService)

	productService := productservice.NewProductService(productRepository, productImageRepository, storeRepository)
	productController := productcontroller.NewProductHandler(productService, storeService, productImageService)

	///////////-----------------------------------------------------------------------------------------------------------////////////////////

	membershipRepositoryCha := membershiprepository.NewMembershipRepository(db)
	membershipServiceCHa := membershipservice.NewMembershipService(membershipRepositoryCha)
	membershipController := membershipcontroller.NewMembershipController(membershipServiceCHa)

	adminRepository := adminrepository.NewAdminRepository(db)
	adminService := adminservice.NewAdminService(adminRepository, bankRepoCha, ownerRepository, customerRepository)
	adminController := admincontroller.NewAdminController(adminService)

	chartRepository := chartrepository.NewChartRepository(db)
	chartService := chartservice.NewChartService(chartRepository, productRepository)
	chartController := chartcontroller.NewChartController(chartService)

	transactionRepositoryCha := transactionrepository.NewTransactionRepository(db)
	transactionServiceCHa := transactionservice.NewTransactionService(transactionRepositoryCha, productRepository, customerRepository, ownerRepository, chartRepository)
	transactionController := transactioncontroller.NewTransactionController(transactionServiceCHa)
	loginController := logincontroller.NewLoginController(ownerService, customerService, adminService)

	var jwtKey = "secret_key"
	pancakaki := r.Group("pancakaki/v1/")
	pancakaki.POST("/login", loginController.Login)

	pancakaki.POST("register/owner", ownerController.CreateOwner)
	pancakaki.POST("register/customer", customerController.Register)
	pancakaki.POST("register/admin", adminController.Register)

	admin := pancakaki.Group("/admins")
	admin.Use(helper.AuthMiddleware(jwtKey))
	{
		admin.GET("/", adminController.ViewAll)
		admin.GET("/:id", adminController.ViewOne)
		admin.PUT("/:id", adminController.Edit)
		admin.DELETE("/:id", adminController.Unreg)

		admin.POST("/bank/", adminController.RegisterBank)
		admin.PUT("/bank/:id", adminController.EditBank)
		admin.GET("/banks/", adminController.ViewAllBank)

		admin.POST("/membership/", membershipController.Register)
		admin.GET("/memberships/", membershipController.ViewAll)
		admin.GET("/membership/:id", membershipController.ViewOne)
		admin.PUT("/membership/:id", membershipController.Edit)
		admin.DELETE("/membership/:id", membershipController.Unreg)

		admin.GET("/customer/profiles/", customerController.ViewAll)
		admin.GET("/customer/profile/:name", customerController.ViewOne)
		admin.POST("/merk/", merkcontroller.Register)
		admin.GET("/merks/", merkcontroller.ViewAll)
		admin.GET("/merk/:id", merkcontroller.ViewOne)
		admin.PUT("/merk/:id", merkcontroller.Edit)
		admin.DELETE("/merk/:id", merkcontroller.Unreg)
	}

	pancakaki.GET("/ownerhp/:hp", ownerController.GetOwnerByNoHp)

	owner := pancakaki.Group("/owner")
	owner.Use(helper.AuthMiddleware(jwtKey))
	{
		//owner
		owner.GET("/profile", ownerController.GetOwnerById)
		owner.PUT("/profile", ownerController.UpdateOwner)
		owner.DELETE("/profile", ownerController.DeleteOwner)
		//store
		owner.GET("/store", storeController.GetStoreByOwnerId)
		owner.POST("/store", storeController.CreateMainStore)
		owner.PUT("/store", storeController.UpdateMainStore)
		owner.DELETE("/store/:storeid", storeController.DeleteMainStore)
		//product
		owner.POST("/store/product", productController.InsertMainProduct)
		owner.GET("/store/:storeid/products", productController.FindAllProductByStoreIdAndOwnerId)
		owner.GET("/store/:storeid/product/:productid", productController.FindProductByStoreIdOwnerIdProductId)
		owner.PUT("/store/product", productController.UpdateMainProduct)
		owner.DELETE("/store/:storeid/product/:productid", productController.DeleteMainProduct)
	}

	product := pancakaki.Group("/products")
	{
		product.POST("/", productController.InsertMainProduct)
		product.GET("/", productController.FindAllProduct)
		// product.GET("/:id", productController.FindProductById)
		// product.GET("/name/:name", productController.FindProductByName)
		// product.PUT("/", productController.UpdateProduct)
		// product.PUT("/:id", productController.DeleteProduct)
	}

	productImage := pancakaki.Group("/product-image")
	{
		productImage.POST("/", productImageController.InsertProductImage)
		productImage.GET("/", productImageController.FindAllProductImage)
		productImage.GET("/:id", productImageController.FindProductImageById)
		productImage.GET("/name/:name", productImageController.FindProductImageByName)
		// productImage.PUT("/", productImageController.UpdateProductImage)
		// productImage.PUT("/:id", productImageController.DeleteProductImage)
	}

	customer := pancakaki.Group("/customers")
	customer.Use(helper.AuthMiddleware(jwtKey))
	{

		customer.GET("/profile", customerController.ViewOne)
		customer.PUT("/profile", customerController.Edit)
		customer.DELETE("/profile", customerController.Unreg)

		//------------ TRANSACTION CUSTOMER LANGSUNG BELI --------------------//
		customer.POST("/transaction", transactionController.MakeOrder)
		customer.POST("/transaction/multiple", transactionController.MakeMultipleOrder)
		customer.POST("/payment/:id", transactionController.CustomerPayment)

		//---------- Customer Chart ------------------- //
		customer.POST("/chart/", chartController.Register)
		customer.GET("/charts/", chartController.ViewAll)
		customer.GET("/chart/:id", chartController.ViewOne)
		customer.PUT("/chart/:id", chartController.Edit)
		customer.DELETE("/chart/:id", chartController.Unreg)

		//------- NOTIFICATION ------
		customer.GET("/notification", customerController.Notification)

		////------------------- TEST PAYMENTGATEWAY : PROSESS --------------//
		// customer.POST("/payment", transactionController.CreatePaymentIntent)
	}

	return r
}
