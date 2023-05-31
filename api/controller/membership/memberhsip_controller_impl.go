package membershipcontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webmembership "pancakaki/internal/domain/web/membership"
	membershipservice "pancakaki/internal/service/membership"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = "secret_key"

type MembershipControllerImpl struct {
	membershipService membershipservice.MembershipService
}

func NewMembershipController(membershipService membershipservice.MembershipService) MembershipController {
	return &MembershipControllerImpl{
		membershipService: membershipService,
	}
}

func (membershipController *MembershipControllerImpl) Register(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, Access for administrators only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}

	var membership webmembership.MembershipCreateRequest

	err := context.ShouldBind(&membership)
	if err != nil {
		webResponse := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Invalid input, please enter the input correctly.",
			Data:    err.Error(),
		}
		context.JSON(http.StatusBadRequest, gin.H{"admin/membership": webResponse})
		return
	}

	membershipResponse, err := membershipController.membershipService.Register(membership)
	if err != nil {
		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUST",
			Message: "Invalid input, please enter the input correctly.",
			Data:    err.Error(),
		}
		context.JSON(http.StatusBadRequest, gin.H{"admin/membership": webResponse})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "The data has been successfully added",
		Data:    membershipResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"admin/membership": webResponse})

}

func (membershipController *MembershipControllerImpl) ViewAll(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, Access for administrators only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}
	membershipResponse, err := membershipController.membershipService.ViewAll()
	if err != nil {
		webResponse := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Invalid input, please enter the input correctly.",
			Data:    err.Error(),
		}
		context.JSON(http.StatusBadRequest, gin.H{"admin/membership": webResponse})
		return
	}

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all membership data",
		Data:    membershipResponse,
	}
	context.JSON(http.StatusOK, gin.H{"admin/membership": web_response})
}

func (membershipController *MembershipControllerImpl) ViewOne(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, Access for administrators only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}
	membershipId, _ := strconv.Atoi(context.Param("id"))
	membershipResponse, err := membershipController.membershipService.ViewOne(membershipId)
	if err != nil {
		webResponses := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "STATUS_NOT_FOUND",
			Message: err.Error(),
			Data:    "NULL",
		}
		context.JSON(http.StatusNotFound, gin.H{"admin/membership": webResponses})
		return
	}

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "membership data by membership ID",
		Data:    membershipResponse,
	}
	context.JSON(http.StatusOK, gin.H{"membership": webResponses})
}

func (membershipController *MembershipControllerImpl) Edit(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, Access for administrators only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}
	var membership webmembership.MembershipUpdateRequest
	err := context.ShouldBindJSON(&membership)
	if err != nil {
		webResponse := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Invalid input, please enter the input correctly.",
			Data:    err.Error(),
		}
		context.JSON(http.StatusBadRequest, gin.H{"admin/membership": webResponse})
		return
	}

	membershipId, err := strconv.Atoi(context.Param("id"))
	membership.Id = membershipId

	membershipResponse, err := membershipController.membershipService.Edit(membership)
	if err != nil {
		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid input, please enter the input correctly.",
			Data:    err.Error(),
		}
		context.JSON(http.StatusBadRequest, gin.H{"admin/membership": webResponse})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the data has been updated",
		Data:    membershipResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"membership": webResponse})
}

func (membershipController *MembershipControllerImpl) Unreg(context *gin.Context) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "admin" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "user is unauthorized, Access for administrators only",
			Data:    "NULL",
		}
		context.JSON(http.StatusUnauthorized, result)
		return
	}
	membershipId, _ := strconv.Atoi(context.Param("id"))
	membershipResponse, err := membershipController.membershipService.Unreg(membershipId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "STATUS_NOT_FOUND",
			Message: err.Error(),
			Data:    "NULL",
		}
		context.JSON(http.StatusBadRequest, gin.H{"admin/membership": webResponse})
		return
	}

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "membership data has been deleted",
		Data:    membershipResponse,
	}
	context.JSON(http.StatusOK, gin.H{"membership": webResponses})
}
