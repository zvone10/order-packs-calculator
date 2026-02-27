package main

import (
	"fmt"
	"log"
	"net/http"
	"order-pack-calculator-api/internal/calculator"
	"order-pack-calculator-api/internal/handler"
	"order-pack-calculator-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `envconfig:"PORT" default:"8080"`
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Failed to process environment variables: %v", err)
	}

	calculator := calculator.NewCalculator()
	svc := service.NewPackingService(calculator)
	h := handler.NewPackingHandler(svc)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	v1.POST("/pack-calculation", h.Calculate)

	addr := fmt.Sprintf(":%d", cfg.Port)
	r.Run(addr)
}
