package handler

import (
	"github.com/gin-gonic/gin"
)

func ParaHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
	})
}
