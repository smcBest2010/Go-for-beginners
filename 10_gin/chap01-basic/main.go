package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON으로 받을 데이터 구조
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Gin Router 생성
	r := gin.Default()

	// GET /
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Gin",
		})
	})

	// GET /users/123
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	// GET /search?name=kim
	r.GET("/search", func(c *gin.Context) {
		name := c.Query("name")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	// POST /users
	r.POST("/users", func(c *gin.Context) {
		var user User

		// 요청 JSON → User 구조체
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "잘못된 JSON입니다.",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "유저 생성 완료",
			"user":    user,
		})
	})

	// 서버 실행
	r.Run(":8080")
}
