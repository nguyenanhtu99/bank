package main

import (
	"bank/pkg/handler"
	"fmt"
)

func main() {
	router, err := handler.New()
	if err != nil {
		fmt.Printf("Error router: %v\n", err)
	}

	if err := router.Run("localhost:1234"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}