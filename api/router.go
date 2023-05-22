package api

import (
	"database/sql"
	"fmt"
	merkcontroller "pancakaki/api/controller/merk"
	packetcontroller "pancakaki/api/controller/packet"
	productcontroller "pancakaki/api/controller/product"
	productimagecontroller "pancakaki/api/controller/product_image"
	merkrepository "pancakaki/internal/repository/merk"
	packetrepository "pancakaki/internal/repository/packet"
	productrepository "pancakaki/internal/repository/product"
	productimagerepository "pancakaki/internal/repository/product_image"
	merkservice "pancakaki/internal/service/merk"
	packetservice "pancakaki/internal/service/packet"
	productservice "pancakaki/internal/service/product"
	productimageservice "pancakaki/internal/service/product_image"

	"github.com/gin-gonic/gin"
)

func Run(db *sql.DB) *gin.Engine {
	r := gin.Default()

	merkRepository := merkrepository.NewMerkRepository(db)
	merkService := merkservice.NewMerkService(merkRepository)
	merkController := merkcontroller.NewMerkHandler(merkService)

	packetRepository := packetrepository.NewPacketRepository(db)
	packetService := packetservice.NewPacketService(packetRepository)
	packetController := packetcontroller.NewPacketHandler(packetService)

	productRepository := productrepository.NewProductRepository(db)
	productService := productservice.NewProductService(productRepository)
	productController := productcontroller.NewProductHandler(productService)

	productImageRepository := productimagerepository.NewProductImageRepository(db)
	productImageService := productimageservice.NewProductImageService(productImageRepository)
	productImageController := productimagecontroller.NewProductImageHandler(productImageService)

	pancakaki := r.Group("pancakaki/v1/")
	merk := pancakaki.Group("/merks")
	{
		merk.POST("/", merkController.InsertMerk)
		merk.GET("/", merkController.FindAllMerk)
		merk.GET("/:id", merkController.FindMerkById)
		merk.GET("/name/:name", merkController.FindMerkByName)
		merk.PUT("/", merkController.UpdateMerk)
		merk.PUT("/:id", merkController.DeleteMerk)
	}

	packet := pancakaki.Group("/packets")
	{
		packet.POST("/", packetController.InsertPacket)
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
