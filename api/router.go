package api

import (
	"database/sql"
	admincontroller "pancakaki/api/controller/admin"
	customercontroller "pancakaki/api/controller/customer"
	logincontroller "pancakaki/api/controller/login"
	membershipcontroller "pancakaki/api/controller/membership"
	merkcontroller "pancakaki/api/controller/merk"
	ownercontroller "pancakaki/api/controller/owner"
	productcontroller "pancakaki/api/controller/product"
	productimagecontroller "pancakaki/api/controller/product_image"
	storecontroller "pancakaki/api/controller/store"
	adminrepository "pancakaki/internal/repository/admin"
	bankrepository "pancakaki/internal/repository/bank"
	bankstorerepository "pancakaki/internal/repository/bank_store"
	customerrepository "pancakaki/internal/repository/customer"
	membershiprepository "pancakaki/internal/repository/membership"
	merkrepository "pancakaki/internal/repository/merk"
	ownerrepository "pancakaki/internal/repository/owner"
	productrepository "pancakaki/internal/repository/product"
	productimagerepository "pancakaki/internal/repository/product_image"
	storerepository "pancakaki/internal/repository/store"
	adminservice "pancakaki/internal/service/admin"
	bankservice "pancakaki/internal/service/bank"
	customerservice "pancakaki/internal/service/customer"
	membershipservice "pancakaki/internal/service/membership"
	merkservice "pancakaki/internal/service/merk"
	ownerservice "pancakaki/internal/service/owner"
	productservice "pancakaki/internal/service/product"
	productimageservice "pancakaki/internal/service/product_image"
	storeservice "pancakaki/internal/service/store"
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

	loginController := logincontroller.NewLoginController(ownerService, customerService)

	var jwtKey = "secret_key"
	pancakaki := r.Group("pancakaki/v1/")

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
		admin.GET("/transaction_history/customer/:id", adminController.ViewTransactionCustomerById)

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

	pancakaki.POST("/login", loginController.Login)

	pancakaki.POST("/register/owner", ownerController.CreateOwner)
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
		product.POST("/", productController.InsertMainProduct)
		product.GET("/", productController.FindAllProduct)
		// product.GET("/:id", productController.FindProductById)
		// product.GET("/name/:name", productController.FindProductByName)
		// product.PUT("/", productController.UpdateProduct)
		product.PUT("/:id", productController.DeleteProduct)
	}

	productImage := pancakaki.Group("/product-image")
	{
		productImage.POST("/", productImageController.InsertProductImage)
		productImage.GET("/", productImageController.FindAllProductImage)
		productImage.GET("/:id", productImageController.FindProductImageById)
		productImage.GET("/name/:name", productImageController.FindProductImageByName)
		// productImage.PUT("/", productImageController.UpdateProductImage)
		productImage.PUT("/:id", productImageController.DeleteProductImage)
	}

	customer := pancakaki.Group("/customers")
	customer.Use(helper.AuthMiddleware(jwtKey))
	{
		customer.POST("/", customerController.Register)
		customer.GET("/", customerController.ViewAll)
		customer.GET("/:name", customerController.ViewOne)
		customer.PUT("/:id", customerController.Edit)
		customer.DELETE("/:name", customerController.Unreg)
	}
	return r
}
