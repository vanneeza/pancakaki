package ownercontroller

import (
	"net/http"
	entity "pancakaki/internal/domain/entity/owner"
	"pancakaki/internal/domain/web"
	webowner "pancakaki/internal/domain/web/owner"
	bankservice "pancakaki/internal/service/bank"
	membershipservice "pancakaki/internal/service/membership"
	ownerservice "pancakaki/internal/service/owner"
	"pancakaki/utils/helper"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = "secret_key"

type OwnerHandler interface {
	CreateOwner(ctx *gin.Context)
	GetOwnerById(ctx *gin.Context)
	UpdateOwner(ctx *gin.Context)
	DeleteOwner(ctx *gin.Context)
	LoginOwner(ctx *gin.Context)
}

type ownerHandler struct {
	ownerService      ownerservice.OwnerService
	membershipService membershipservice.MembershipService
	bankService       bankservice.BankService
}

func (h *ownerHandler) CreateOwner(ctx *gin.Context) {
	var owner entity.Owner

	if err := ctx.ShouldBindJSON(&owner); err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	getMembershipById, err := h.membershipService.GetMembershipById(owner.MembershipId)
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

	getBankAdminById, err := h.bankService.GetBankAdminById(1)
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

	// fmt.Println(getBankAdminById)
	newOwner, err := h.ownerService.CreateOwner(&owner)
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
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["ownerId"] = strconv.Itoa(newOwner.Id)
	claims["nohp"] = owner.NoHp
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	var jwtKeyByte = []byte(jwtKey)
	tokenString, err := token.SignedString(jwtKeyByte)
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

	resultOwner := webowner.OwnerCreateResponse{
		Id:              newOwner.Id,
		Name:            owner.Name,
		NoHp:            owner.NoHp,
		Email:           owner.Email,
		Password:        owner.Password,
		MembershipName:  getMembershipById.Name,
		MembershipPrice: getMembershipById.Price,
		Bank:            getBankAdminById,
		Token:           tokenString,
	}
	result := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "success create owner",
		Message: "success create owner",
		Data:    resultOwner,
	}
	ctx.JSON(http.StatusCreated, result)
}

func (h *ownerHandler) GetOwnerById(ctx *gin.Context) {
	ownerName := ctx.Param("ownername")
	getOwnerByName, err := h.ownerService.GetOwnerByName(ownerName)
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
	if getOwnerByName == nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "status not found",
			Message: "status not found",
			Data:    "owner with name " + ownerName + " not found",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["ownerId"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)

	ownerById, err := h.ownerService.GetOwnerById(ownerIdInt)
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

	if getOwnerByName.Id != ownerIdInt {
		result := web.WebResponse{
			Code:    http.StatusConflict,
			Status:  "status conflict",
			Message: "status conflict",
			Data:    "owner with name " + ownerName + " is unauthorized",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "get by id success",
		Message: "success get owner with id " + ownerId,
		Data:    ownerById,
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *ownerHandler) UpdateOwner(ctx *gin.Context) {
	// idParam := ctx.Param("id")
	// id, err := strconv.Atoi(idParam)
	// if err != nil {
	// 	result := web.WebResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Status:  "bad request",
	// 		Message: "bad request",
	// 		Data:    err.Error(),
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, result)
	// 	return
	// }

	ownerName := ctx.Param("ownername")
	getOwnerByName, err := h.ownerService.GetOwnerByName(ownerName)
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
	if getOwnerByName == nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "status not found",
			Message: "status not found",
			Data:    "owner with name " + ownerName + " not found",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["ownerId"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)

	if getOwnerByName.Id != ownerIdInt {
		result := web.WebResponse{
			Code:    http.StatusConflict,
			Status:  "status conflict",
			Message: "status conflict",
			Data:    "owner with name " + ownerName + " is unauthorized",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	var owner entity.Owner

	if err := ctx.ShouldBindJSON(&owner); err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	if owner.Id != ownerIdInt {
		ownerIdJson := strconv.Itoa(owner.Id)
		result := web.WebResponse{
			Code:    http.StatusConflict,
			Status:  "status conflict",
			Message: "status conflict",
			Data:    "owner with id " + ownerIdJson + " is unauthorized",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	owner.Id = ownerIdInt
	ownerUpdate, err := h.ownerService.UpdateOwner(&owner)
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
		Status:  "update success",
		Message: "success update owner with id " + ownerId,
		Data:    ownerUpdate,
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *ownerHandler) DeleteOwner(ctx *gin.Context) {
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

	ownerName := ctx.Param("ownername")
	getOwnerByName, err := h.ownerService.GetOwnerByName(ownerName)
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
	if getOwnerByName == nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "status not found",
			Message: "status not found",
			Data:    "owner with name " + ownerName + " not found",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["ownerId"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)

	if getOwnerByName.Id != ownerIdInt {
		result := web.WebResponse{
			Code:    http.StatusConflict,
			Status:  "status conflict",
			Message: "status conflict",
			Data:    "owner with name " + ownerName + " is unauthorized",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	if id != ownerIdInt {
		result := web.WebResponse{
			Code:    http.StatusConflict,
			Status:  "status conflict",
			Message: "status conflict",
			Data:    "owner with id " + idParam + " is unauthorized",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	err = h.ownerService.DeleteOwner(ownerIdInt)
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
		Status:  "delete success",
		Message: "success delete owner with id " + ownerId,
		Data:    err,
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *ownerHandler) LoginOwner(ctx *gin.Context) {
	var owner entity.Owner

	if err := ctx.ShouldBindJSON(&owner); err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	ownerEmail := owner.Email
	getOwnerByEmail, err := h.ownerService.GetOwnerByEmail(ownerEmail)
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
	if getOwnerByEmail == nil {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "status not found",
			Message: "status not found",
			Data:    "owner with email " + ownerEmail + " not found",
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	if ownerEmail != getOwnerByEmail.Email {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    "wrong email",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ownerPassword := owner.Password
	match := helper.CheckPasswordHash(ownerPassword, getOwnerByEmail.Password)
	if !match {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    "wrong password",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["ownerId"] = strconv.Itoa(getOwnerByEmail.Id)
	claims["nohp"] = getOwnerByEmail.NoHp
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	var jwtKeyByte = []byte(jwtKey)
	tokenString, err := token.SignedString(jwtKeyByte)

	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "login success",
		Message: "success login owner with email " + getOwnerByEmail.Email,
		Data:    tokenString,
	}
	ctx.JSON(http.StatusOK, result)

}

func NewOwnerHandler(
	ownerService ownerservice.OwnerService,
	membershipService membershipservice.MembershipService,
	bankService bankservice.BankService) OwnerHandler {
	return &ownerHandler{
		ownerService:      ownerService,
		membershipService: membershipService,
		bankService:       bankService}
}
