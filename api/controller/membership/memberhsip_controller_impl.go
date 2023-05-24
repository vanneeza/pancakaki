package membershipcontroller

import (
	"net/http"
	"pancakaki/internal/domain/web"
	webmembership "pancakaki/internal/domain/web/membership"
	membershipservice "pancakaki/internal/service/membership"
	"pancakaki/utils/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MembershipControllerImpl struct {
	membershipService membershipservice.MembershipService
}

func NewMembershipController(membershipService membershipservice.MembershipService) MembershipController {
	return &MembershipControllerImpl{
		membershipService: membershipService,
	}
}

func (membershipController *MembershipControllerImpl) Register(context *gin.Context) {

	var membership webmembership.MembershipCreateRequest

	err := context.ShouldBind(&membership)
	helper.InternalServerError(err, context)

	membershipResponse, err := membershipController.membershipService.Register(membership)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "The data has been successfully added",
		Data:    membershipResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"membership": webResponse})

}

func (membershipController *MembershipControllerImpl) ViewAll(context *gin.Context) {
	membershipResponse, err := membershipController.membershipService.ViewAll()
	helper.InternalServerError(err, context)

	web_response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all membership data",
		Data:    membershipResponse,
	}
	context.JSON(http.StatusOK, gin.H{"membership": web_response})
}

func (membershipController *MembershipControllerImpl) ViewOne(context *gin.Context) {
	membershipId, _ := strconv.Atoi(context.Param("id"))
	membershipResponse, err := membershipController.membershipService.ViewOne(membershipId)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "membership data by membership ID",
		Data:    membershipResponse,
	}
	context.JSON(http.StatusOK, gin.H{"membership": webResponses})
}

func (membershipController *MembershipControllerImpl) Edit(context *gin.Context) {
	var membership webmembership.MembershipUpdateRequest
	err := context.ShouldBindJSON(&membership)
	helper.InternalServerError(err, context)

	membershipId, _ := strconv.Atoi(context.Param("id"))
	membership.Id = membershipId

	membershipResponse, err := membershipController.membershipService.Edit(membership)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the data has been updated",
		Data:    membershipResponse,
	}
	context.JSON(http.StatusCreated, gin.H{"membership": webResponse})
}

func (membershipController *MembershipControllerImpl) Unreg(context *gin.Context) {
	membershipId, _ := strconv.Atoi(context.Param("id"))
	membershipResponse, err := membershipController.membershipService.Unreg(membershipId)
	helper.InternalServerError(err, context)

	webResponses := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "membership data has been deleted",
		Data:    membershipResponse,
	}
	context.JSON(http.StatusOK, gin.H{"membership": webResponses})
}
