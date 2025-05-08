package handlers

import (
	"muraragi/street-racer-arena-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseCarHandler struct {
	baseCarService services.BaseCarService
}

func NewBaseCarHandler(baseCarService services.BaseCarService) *BaseCarHandler {
	return &BaseCarHandler{baseCarService: baseCarService}
}

func (h *BaseCarHandler) GetAllBaseCars(c *gin.Context) {
	cars, err := h.baseCarService.GetAllBaseCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cars)
}
