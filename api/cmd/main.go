package main

import (
	"net/http"
	"order-pack-calculator-api/internal/calculator"
	"order-pack-calculator-api/internal/handler"
	"order-pack-calculator-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	calculator := calculator.NewCalculator()
	svc := service.NewPackingService(calculator)
	h := handler.NewPackingHandler(svc)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	v1.POST("/pack-calculation", h.Calculate)

	r.Run(":8080")
}
