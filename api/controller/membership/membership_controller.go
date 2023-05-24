package membershipcontroller

import (
	"github.com/gin-gonic/gin"
)

// type MembershipHandler interface {
// 	GetMembershipById(ctx *gin.Context)
// }

// type membershipHandler struct {
// 	membershipService membershipservice.MembershipService
// }

// func (h *membershipHandler) GetMembershipById(ctx *gin.Context) {
// 	idParam := ctx.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		result := web.WebResponse{
// 			Code:    http.StatusBadRequest,
// 			Status:  "bad request",
// 			Message: "bad request",
// 			Data:    err.Error(),
// 		}
// 		ctx.JSON(http.StatusBadRequest, result)
// 		return
// 	}

// 	memberById, err := h.membershipService.GetMembershipById(id)
// 	if err != nil {
// 		result := web.WebResponse{
// 			Code:    http.StatusInternalServerError,
// 			Status:  "status internal server error",
// 			Message: "status internal server error",
// 			Data:    err.Error(),
// 		}
// 		ctx.JSON(http.StatusInternalServerError, result)
// 		return
// 	}
// 	result := web.WebResponse{
// 		Code:    http.StatusOK,
// 		Status:  "get by id success",
// 		Message: "success get mebership with id " + idParam,
// 		Data:    memberById,
// 	}
// 	ctx.JSON(http.StatusOK, result)
// }

// func NewMembershipHandler(membershipService membershipservice.MembershipService) MembershipHandler {
// 	return &membershipHandler{membershipService: membershipService}
// import "github.com/gin-gonic/gin"

type MembershipController interface {
	Register(context *gin.Context)
	ViewAll(context *gin.Context)
	ViewOne(context *gin.Context)
	Edit(context *gin.Context)
	Unreg(context *gin.Context)
}
