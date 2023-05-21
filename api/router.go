package api

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run(db *sql.DB) *gin.Engine {
	r := gin.Default()

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
