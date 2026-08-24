package main

import (
	"post/query"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", query.QueryHandler)

	r.GET("/:id", query.PathHandeler)

	r.GET("/test/:path", urlParams)

	// r.GET("/:id/search", query.PQHandler)

	r.GET("/:id/search", params)

	r.Run(":3000")
}

func urlParams(c *gin.Context) {
	path := c.Param("path")
	query := c.Query("query")
	c.JSON(200, gin.H{
		"message": "당신의 경로는" + path + "쿼리파라미터는" + query,
	})
}

func params(c *gin.Context) {
	id := c.Param("id")
	price := c.Query("price")
	c.JSON(200, gin.H{
		"message": id + "번의 가격이 " + price + "원 입니다.",
	})
}
