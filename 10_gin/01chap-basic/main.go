// Gin앱 만들기
package main

import (
	"basic/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handler.HomeHandler)
	router.GET("/index", indexHandler)
	router.Run()
}

func indexHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "메인화면"})
}
