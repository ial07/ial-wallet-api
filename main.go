package main

import (
	"ial-wallet/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/login", handler.LoginHandler)
	v1.POST("/register", handler.RegisterHandler)

	router.Run()
}




	