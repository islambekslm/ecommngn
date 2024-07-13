package main

import (
	"ecommngn/internal/app"
	"fmt"
)

func main() {
	fmt.Println("Starting engine...")

	go app.Start()
}
