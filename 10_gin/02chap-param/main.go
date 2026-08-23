package main

import (
	"param/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET(":id", handler.ParaHandler)

	r.GET("serch/:id", serchHandler)

	r.Run(":3000")
}

func serchHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": id,
	})
}
