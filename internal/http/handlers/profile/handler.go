package handler

import (
	"BACKEND_GO/internal/domain/profile/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	service services.ProfileUseCase
}

func NewProfileHandler(service services.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	id := c.Param("id")
	profile, err := h.service.GetProfileService(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"meta": gin.H{"status": "error", "message": err.Error()},
			"data": nil,
		})
		return
	}
	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"meta": gin.H{"status": "error", "message": "Profile not found"},
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{"status": "success", "message": "Profile retrieved successfully"},
		"data": profile,
	})
}
