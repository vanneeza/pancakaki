package ownercontroller

import (
	"fmt"
	"log"
	"net/http"
	"pancakaki/internal/domain/entity"
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
	GetOwnerByNoHp(ctx *gin.Context)
	UpdateOwner(ctx *gin.Context)
	DeleteOwner(ctx *gin.Context)
}

type ownerHandler struct {
	ownerService      ownerservice.OwnerService
	membershipService membershipservice.MembershipService
	bankService       bankservice.BankService
}

func (h *ownerHandler) CreateOwner(ctx *gin.Context) {
	var owner entity.Owner

	err := ctx.ShouldBindJSON(&owner)
	helper.StatusBadRequest(err, ctx)

	getMembershipById, err := h.membershipService.ViewOne(owner.MembershipId)
	log.Println(getMembershipById, "membershipId")
	fmt.Scanln()
	helper.InternalServerError(err, ctx)
	if err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result) //buat ngirim respon
		return
	}

	// bANK
	getBankAdminById, err := h.bankService.GetBankAdminById(1)
	helper.InternalServerError(err, ctx)

	// fmt.Println(getBankAdminById)

	newOwner, err := h.ownerService.CreateOwner(&owner)
	helper.InternalServerError(err, ctx)
	// if err != nil {
	// 	result := web.WebResponse{
	// 		Code:    http.StatusInternalServerError,
	// 		Status:  "INTERNAL_SERVER_ERROR",
	// 		Message: "status internal server error",
	// 		Data:    err.Error(),
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, result) //buat ngirim respon
	// 	return
	// }
	getOwnerById, err := h.ownerService.GetOwnerById(newOwner.Id)
	helper.InternalServerError(err, ctx)
	// newOwnerId := strconv.Itoa(newOwner.Id)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = strconv.Itoa(newOwner.Id)
	claims["role"] = getOwnerById.Role
	claims["nohp"] = getOwnerById.NoHp
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	var jwtKeyByte = []byte(jwtKey)
	tokenString, err := token.SignedString(jwtKeyByte)
	helper.InternalServerError(err, ctx)

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
		Status:  "CREATED",
		Message: "success create owner",
		Data:    resultOwner,
	}
	ctx.JSON(http.StatusCreated, result)
}

func (h *ownerHandler) GetOwnerById(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["id"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)

	role := claims["role"].(string)
	// fmt.Println(role)
	if role != "owner" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		ctx.JSON(http.StatusUnauthorized, result)
		return
	}

	ownerById, err := h.ownerService.GetOwnerById(ownerIdInt)
	helper.InternalServerError(err, ctx)

	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success get owner with id " + ownerId,
		Data:    ownerById,
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *ownerHandler) GetOwnerByNoHp(ctx *gin.Context) {
	// var owner entity.Owner
	// err := ctx.ShouldBindJSON(&owner)
	hp := ctx.Param("hp")
	// storeIdInt, _ := strconv.Atoi(storeId)
	// helper.InternalServerError(err, ctx)

	ownerByNoHp, err := h.ownerService.GetOwnerByNoHp(hp)
	helper.InternalServerError(err, ctx)

	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success get owner with hp " + hp,
		Data:    ownerByNoHp,
	}
	ctx.JSON(http.StatusOK, result)
}
func (h *ownerHandler) UpdateOwner(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["id"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)

	role := claims["role"].(string)
	if role != "owner" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		ctx.JSON(http.StatusUnauthorized, result)
		return
	}

	var owner entity.Owner

	err := ctx.ShouldBindJSON(&owner)
	helper.InternalServerError(err, ctx)

	var statusMembership []entity.Bank
	getMembershipOwnerById, err := h.ownerService.GetOwnerById(ownerIdInt)
	helper.InternalServerError(err, ctx)

	if owner.MembershipId != getMembershipOwnerById.MembershipId {
		getBankAdminById, err := h.bankService.GetBankAdminById(1)
		helper.InternalServerError(err, ctx)
		statusMembership = getBankAdminById
	} else {
		statusMembership = nil
	}

	getMembershipById, err := h.membershipService.ViewOne(owner.MembershipId)
	helper.InternalServerError(err, ctx)

	owner.Id = ownerIdInt
	ownerUpdate, err := h.ownerService.UpdateOwner(&owner)
	helper.InternalServerError(err, ctx)

	resultOwner := webowner.OwnerUpdateResponse{
		Id:              owner.Id,
		Name:            ownerUpdate.Name,
		NoHp:            ownerUpdate.NoHp,
		Email:           ownerUpdate.Email,
		Password:        ownerUpdate.Password,
		MembershipName:  getMembershipById.Name,
		MembershipPrice: getMembershipById.Price,
		Bank:            statusMembership,
	}

	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success update owner with id " + ownerId,
		Data:    resultOwner,
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *ownerHandler) DeleteOwner(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	ownerId := claims["ownerId"].(string)
	ownerIdInt, _ := strconv.Atoi(ownerId)

	role := claims["role"].(string)
	if role != "owner" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		ctx.JSON(http.StatusUnauthorized, result)
		return
	}

	err := h.ownerService.DeleteOwner(ownerIdInt)
	helper.InternalServerError(err, ctx)
	result := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "success delete owner with id " + ownerId,
		Data:    err,
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
