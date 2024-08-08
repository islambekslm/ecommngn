package app

import (
	"ecommngn/internal/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/rs/zerolog"
)

type ProductCatalogue struct {
	logger zerolog.Logger
}

func NewProductCatalogue() *ProductCatalogue {
	return &ProductCatalogue{
		logger: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func (pc *ProductCatalogue) Start() {
	fmt.Print("App started")
	router := gin.Default()
	handlers.RegisterOrderRoutes(router)
	err := router.Run(":8080")
	if err != nil {
		pc.logger.Error().Err(err).Msg("Error starting HTTP server")
	}
}

func (pc *ProductCatalogue) Stop() {
	fmt.Print("App stopped")
}
