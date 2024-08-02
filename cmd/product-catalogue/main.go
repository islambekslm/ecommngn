package main

import (
	"ecommngn/internal/app"
	"fmt"
)

func main() {
	fmt.Println("Starting product catalogue...")
	productCatalogue := app.NewProductCatalogue()
	go productCatalogue.Start()
	defer productCatalogue.Stop()
}
