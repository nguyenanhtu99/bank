package main

import (
	"bank/pkg/handler"
	"fmt"
)

func main() {
	router, err := handler.New()
	if err != nil {
		fmt.Println("Error router")
	}

	router.Run("localhost:1234")
}