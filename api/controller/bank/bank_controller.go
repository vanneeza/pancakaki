package bankcontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	bankservice "pancakaki/internal/service/bank"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BankHandler interface {
	GetBankAdminById(ctx *gin.Context)
}

type bankHandler struct {
	bankService bankservice.BankService
}

func (h *bankHandler) GetBankAdminById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	bankAdminById, err := h.bankService.GetBankAdminById(id)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "get by id success",
		Message: "success get bank admin with id " + idParam,
		Data:    bankAdminById,
	}
	ctx.JSON(http.StatusOK, result)
}

func NewBankHandler(bankService bankservice.BankService) BankHandler {
	return &bankHandler{bankService: bankService}
}
