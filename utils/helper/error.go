package helper

import (
	"net/http"
	"pancakaki/internal/domain/web"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InternalServerError(err error, context *gin.Context) {
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Error",
			Data:    "NULL",
		}
		context.JSON(http.StatusInternalServerError, result)
		return
	}
}

func StatusBadRequest(err error, context *gin.Context) {
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "bad request",
			Data:    err.Error(),
		}
		context.JSON(http.StatusBadRequest, result) //buat ngirim respon
		return
	}
}

func StatusNotFound(err error, context *gin.Context) {
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "not found",
			Data:    err.Error(),
		}
		context.JSON(http.StatusNotFound, result) //buat ngirim respon
		return
	}
}

func CheckName(name string) bool {
	// checkName := regexp.MustCompile(`\d`).MatchString(name)
	nameTrim := strings.ReplaceAll(name, " ", "")
	checkName := regexp.MustCompile(`^[a-zA-Z]+$`)
	isMatch := checkName.MatchString(nameTrim)

	return isMatch
}

func CheckHp(hp string) bool {
	hpTrim := strings.ReplaceAll(hp, " ", "")
	checkHp, _ := regexp.Compile(`[0-9]+`)
	isMatch := checkHp.MatchString(hpTrim)
	// if isMatch {
	// 	return true
	// }
	return isMatch
}
