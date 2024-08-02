package app

import (
	"fmt"

	"github.com/rs/zerolog"
)

type productCatalogue struct {
	logger zerolog.Logger
}

func NewProductCatalogue() *productCatalogue {
	return &productCatalogue{}
}

func (pc *productCatalogue) Start() {
	fmt.Print("App started")
}

func (pc *productCatalogue) Stop() {
	fmt.Print("App stopped")
}
