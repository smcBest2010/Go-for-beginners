package main

import (
	"post/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/post", handler.PostCreate)

	r.Run()
}
