package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InternalServerError(err error, context *gin.Context) {
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //buat ngirim respon
		return
	}
}
