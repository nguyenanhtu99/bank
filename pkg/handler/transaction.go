package handler

import (
	"bank/pkg/executor"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateTransaction(svc executor.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload executor.CreateTransactionInput

		if err := c.ShouldBindJSON(&payload); err != nil {
			fmt.Printf("Handler: Failed to parse payload: %v\n", err)
			c.AbortWithStatusJSON(400, gin.H{
				"message": "Invalid input",
				"error": err.Error(),
			})
			return
		}

		if err := validator.New().Struct(payload); err != nil {
			fmt.Printf("Handler: Invalid input: %v\n", err)
			c.AbortWithStatusJSON(400, gin.H{
				"message": "Invalid input",
				"error": err.Error(),
			})
			return
		}

		if err := svc.CreateTransaction(&payload); err != nil {
			fmt.Printf("Handler: Failed to create transaction: %v\n", err)
			c.AbortWithStatusJSON(500, gin.H{
				"message": "Failed to create transaction",
				"error": err.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(200, gin.H{
			"message": "Create transaction successfully!",
		})
	}
}

func GetTransaction(svc executor.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Query("userName")		
		transaction, err := svc.GetTransaction(userName)
		if err != nil {
			fmt.Printf("Handler: Can't get account: %v\n", err)
		}

		c.AbortWithStatusJSON(200, gin.H{
			"data": transaction,
			"message": "Get account successfully!",
		})
	}
}