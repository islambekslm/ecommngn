package app

import (
	"fmt"

	"github.com/rs/zerolog"
)

type orderManager struct {
	logger zerolog.Logger
}

func NewOrderManager() *orderManager {
	return &orderManager{}
}

func (om *orderManager) Start() {
	fmt.Println("App started")
}

func (om *orderManager) Stop() {
	fmt.Println("App stopped")
}
