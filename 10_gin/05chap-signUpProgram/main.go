package main

import (
	"signUp/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	api.POST("/signup", handler.SignUpHandler)
	api.POST("/login", handler.LoginHandler)

	router.Run()
}
