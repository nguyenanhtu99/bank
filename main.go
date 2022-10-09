package main

import (
	"bank/pkg/handler"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx         context.Context
	mongoClient *mongo.Client
)

func main() {
	router, err := handler.New()
	if err != nil {
		fmt.Println("Error router")
	}

	router.Run("localhost:1234")
}