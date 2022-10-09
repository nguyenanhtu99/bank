package handler

import (
	"bank/pkg/executor"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() (*gin.Engine, error) {
	svc, err := executor.New()
	if err != nil {
		return nil, err
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) { c.AbortWithStatus(http.StatusOK) })
	v1 := router.Group("/v1")

	v1.POST(postAccount, CreateAccount(svc))
	v1.GET(getAccount, GetAccount(svc))

	v1.POST(postTransaction, CreateTransaction(svc))

	return router, nil
}

