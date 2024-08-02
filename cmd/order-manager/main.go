package main

import (
	"ecommngn/internal/app"
	"fmt"
)

func main() {
	fmt.Println("Starting order manager...")
	orderManager := app.NewOrderManager()
	go orderManager.Start()
	defer orderManager.Stop()
}
