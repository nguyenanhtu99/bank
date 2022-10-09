package handler

import (
	"bank/pkg/executor"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateAccount(svc executor.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Handler: creating account ...")
		var payload executor.CreateAccountInput

		if err := c.ShouldBindJSON(&payload); err != nil {
			fmt.Printf("Handler: failed to parse payload: %v\n", err)
			c.AbortWithStatusJSON(400, gin.H{
				"message": "Invalid input",
				"error": err.Error(),
			})
			return
		}

		if err := validator.New().Struct(payload); err != nil {
			fmt.Printf("Handler: invalid input: %v\n", err)
			c.AbortWithStatusJSON(400, gin.H{
				"message": "Invalid input",
				"error": err.Error(),
			})
			return
		}

		if _, err := svc.GetAccount(payload.UserName); err == nil {
			fmt.Println("Handler: Account already in use")
			c.AbortWithStatusJSON(400, gin.H{
				"message": "Failed to create account",
				"error": "Account already in use",
			})
			return
		}

		if err := svc.CreateAccount(&payload); err != nil {
			fmt.Printf("Handler: Failed to create account: %v\n", err)
			c.AbortWithStatusJSON(500, gin.H{
				"message": "Failed to create account",
				"error": err.Error(),
			})
			return
		}

		fmt.Println("Handler: create account successfully!")
		c.AbortWithStatusJSON(200, gin.H{
			"message": "Create account successfully!",
		})
	}
}

func GetAccount(svc executor.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Handler: getting account ...")
		userName := c.Query("userName")		
		account, err := svc.GetAccount(userName)
		if err != nil {
			fmt.Printf("Handler: Failed to get account: %v\n", err)
		}

		c.AbortWithStatusJSON(200, gin.H{
			"data": account,
			"message": "Get account successfully!",
		})
	}
}