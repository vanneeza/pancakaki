package api

import (
	"database/sql"
	"fmt"
	admincontroller "pancakaki/api/controller/admin"
	adminrepository "pancakaki/internal/repository/admin"
	adminservice "pancakaki/internal/service/admin"

	"github.com/gin-gonic/gin"
)

func Run(db *sql.DB) *gin.Engine {
	r := gin.Default()

	adminRepository := adminrepository.NewAdminRepository(db)
	adminService := adminservice.NewAdminService(adminRepository)
	adminController := admincontroller.NewAdminController(adminService)

	pancakaki := r.Group("pancakaki/v1/")

	admin := pancakaki.Group("/admins")
	{
		admin.POST("/", adminController.Register)
		admin.GET("/", adminController.ViewAll)
		admin.GET("/:id", adminController.ViewOne)
		admin.PUT("/:id", adminController.Edit)
		admin.DELETE("/:id", adminController.Unreg)
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
