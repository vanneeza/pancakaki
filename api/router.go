package api

import (
	"database/sql"
	"fmt"
<<<<<<< HEAD
	merkcontroller "pancakaki/api/controller/merk"
	ownercontroller "pancakaki/api/controller/owner"
	packetcontroller "pancakaki/api/controller/packet"
	productcontroller "pancakaki/api/controller/product"
	productimagecontroller "pancakaki/api/controller/product_image"
	storecontroller "pancakaki/api/controller/store"
	bankrepository "pancakaki/internal/repository/bank"
	bankstorerepository "pancakaki/internal/repository/bank_store"
	membershiprepository "pancakaki/internal/repository/membership"
	merkrepository "pancakaki/internal/repository/merk"
	ownerrepository "pancakaki/internal/repository/owner"
	packetrepository "pancakaki/internal/repository/packet"
	productrepository "pancakaki/internal/repository/product"
	productimagerepository "pancakaki/internal/repository/product_image"
	storerepository "pancakaki/internal/repository/store"
	bankservice "pancakaki/internal/service/bank"
	membershipservice "pancakaki/internal/service/membership"
	merkservice "pancakaki/internal/service/merk"
	ownerservice "pancakaki/internal/service/owner"
	packetservice "pancakaki/internal/service/packet"
	productservice "pancakaki/internal/service/product"
	productimageservice "pancakaki/internal/service/product_image"
	storeservice "pancakaki/internal/service/store"
	"pancakaki/utils/helper"
=======
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
>>>>>>> chauzar

	"github.com/gin-gonic/gin"
)

func Run(db *sql.DB) *gin.Engine {
	r := gin.Default()

<<<<<<< HEAD
	merkRepository := merkrepository.NewMerkRepository(db)
	merkService := merkservice.NewMerkService(merkRepository)
	merkController := merkcontroller.NewMerkHandler(merkService)

	packetRepository := packetrepository.NewPacketRepository(db)
	packetService := packetservice.NewPacketService(packetRepository)
	packetController := packetcontroller.NewPacketHandler(packetService)

	productImageRepository := productimagerepository.NewProductImageRepository(db)
	productImageService := productimageservice.NewProductImageService(productImageRepository)
	productImageController := productimagecontroller.NewProductImageHandler(productImageService)

	membershipRepository := membershiprepository.NewMembershipRepository(db)
	membershipService := membershipservice.NewMembershipService(membershipRepository)
	// membershipController := membershipcontroller.NewMembershipHandler(membershipService)

	bankRepository := bankrepository.NewBankRepository(db)
	bankService := bankservice.NewBankService(bankRepository)

	bankStoreRepository := bankstorerepository.NewBankStoreRepository(db)

	ownerRepository := ownerrepository.NewOwnerRepository(db)
	ownerService := ownerservice.NewOwnerService(ownerRepository)
	ownerController := ownercontroller.NewOwnerHandler(ownerService, membershipService, bankService)

	storeRepository := storerepository.NewStoreRepository(db, bankRepository, bankStoreRepository)
	storeService := storeservice.NewStoreService(storeRepository)
	storeController := storecontroller.NewStoreHandler(storeService, ownerService)

	productRepository := productrepository.NewProductRepository(db)
	productService := productservice.NewProductService(productRepository)
	productController := productcontroller.NewProductHandler(productService, storeService)

	var jwtKey = "secret_key"
	pancakaki := r.Group("pancakaki/v1/")

	pancakaki.POST("/login", ownerController.LoginOwner)
	pancakaki.POST("/", ownerController.CreateOwner)

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

	merk := pancakaki.Group("/customers")
	{
		store := pancakaki.Group("/stores")
		{
			store.POST("/product", merkController.InsertMerk)
			merk.POST("/merk", merkController.InsertMerk)
			// packet.POST("/oacket", packetController.InsertPacket)
		}

		merk.GET("/", merkController.FindAllMerk)
		merk.GET("/:id", merkController.FindMerkById)
		merk.GET("/name/:name", merkController.FindMerkByName)
		merk.PUT("/", merkController.UpdateMerk)
		merk.PUT("/:id", merkController.DeleteMerk)
	}

	packet := pancakaki.Group("/packets")
	{

		packet.GET("/", packetController.FindAllPacket)
		packet.GET("/:id", packetController.FindpacketById)
		packet.GET("/name/:name", packetController.FindPacketByName)
		packet.PUT("/", packetController.UpdatePacket)
		packet.PUT("/:id", packetController.DeletePacket)
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
=======
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

		admin.GET("/customer/profiles/", customerController.ViewAll)
		admin.GET("/customer/profile/:name", customerController.ViewOne)
	}

	customer := pancakaki.Group("/customers")
	{
		customer.POST("/", customerController.Register)
		customer.GET("/", customerController.ViewAll)
		customer.GET("/:name", customerController.ViewOne)
		customer.PUT("/:id", customerController.Edit)
		customer.DELETE("/:name", customerController.Unreg)
>>>>>>> chauzar
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
