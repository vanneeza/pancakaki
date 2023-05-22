package packetcontroller

import (
	"net/http"
	entity "pancakaki/internal/domain/entity/packet"
	"pancakaki/internal/domain/web"
	packetservice "pancakaki/internal/service/packet"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PacketHandler interface {
	InsertPacket(ctx *gin.Context)
	UpdatePacket(ctx *gin.Context)
	DeletePacket(ctx *gin.Context)
	FindpacketById(ctx *gin.Context)
	FindPacketByName(ctx *gin.Context)
	FindAllPacket(ctx *gin.Context)
}

type packetHandler struct {
	packetService packetservice.PacketService
}

// DeletePacket implements PacketHandler
func (h *packetHandler) DeletePacket(ctx *gin.Context) {
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
	var packet entity.Packet
	packet.Id = id

	if err := ctx.ShouldBindJSON(&packet); err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	err = h.packetService.DeletePacket(&packet)
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
		Message: "success delete packet with id " + idParam,
		Data:    err,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindAllPacket implements PacketHandler
func (h *packetHandler) FindAllPacket(ctx *gin.Context) {
	packetList, err := h.packetService.FindAllPacket()
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
		Status:  "success get packet list",
		Message: "success get packet list",
		Data:    packetList,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindPacketByName implements PacketHandler
func (h *packetHandler) FindPacketByName(ctx *gin.Context) {
	packetName := ctx.Param("name")

	packetByName, err := h.packetService.FindPacketByName(packetName)
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
		Message: "success get packet with name " + packetName,
		Data:    packetByName,
	}
	ctx.JSON(http.StatusOK, result)
}

// FindpacketById implements PacketHandler
func (h *packetHandler) FindpacketById(ctx *gin.Context) {
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

	packetById, err := h.packetService.FindpacketById(id)
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
		Message: "success get packet with id " + idParam,
		Data:    packetById,
	}
	ctx.JSON(http.StatusOK, result)
}

// InsertPacket implements PacketHandler
func (h *packetHandler) InsertPacket(ctx *gin.Context) {
	var packet entity.Packet
	if err := ctx.ShouldBindJSON(&packet); err != nil {
		result := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "bad request",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	newPacket, err := h.packetService.InsertPacket(&packet)
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
		Code:    http.StatusCreated,
		Status:  "success insert packet",
		Message: "success insert packet",
		Data:    newPacket,
	}
	ctx.JSON(http.StatusCreated, result)
}

// UpdatePacket implements PacketHandler
func (h *packetHandler) UpdatePacket(ctx *gin.Context) {
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
	var packet entity.Packet
	packet.Id = id

	if err := ctx.ShouldBindJSON(&packet); err != nil {
		result := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "status internal server error",
			Message: "status internal server error",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	packetUpdate, err := h.packetService.UpdatePacket(&packet)
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
		Message: "success update packet with id " + idParam,
		Data:    packetUpdate,
	}
	ctx.JSON(http.StatusOK, result)
}

func NewPacketHandler(packetService packetservice.PacketService) PacketHandler {
	return &packetHandler{packetService: packetService}
}
