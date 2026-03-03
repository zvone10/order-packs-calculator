// handler/packing_handler.go
package handler

import (
	"net/http"
	"order-pack-calculator-api/internal/calculator"
	"order-pack-calculator-api/internal/model"
	"order-pack-calculator-api/internal/service"

	"github.com/gin-gonic/gin"
)

type PackingHandler struct {
	svc service.PackingService
}

func NewPackingHandler(svc service.PackingService) *PackingHandler {
	return &PackingHandler{svc: svc}
}

func (h *PackingHandler) Calculate(c *gin.Context) {
	var req model.PackRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that BoxCapacity doesn't contain duplicates
	if calculator.HasSameItems(req.BoxCapacity) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "boxCapacity contains duplicate values"})
		return
	}

	result, err := h.svc.Calculate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
