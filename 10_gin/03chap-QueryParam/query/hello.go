package query

import (
	"github.com/gin-gonic/gin"
)

func QueryHandler(c *gin.Context) {
	sort := c.Query("sort")
	c.JSON(200, gin.H{
		"message": "sort방식은 " + sort,
	})
}

// 패스 파라미터
func PathHandeler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": id,
	})
}

// 패스 파라미터 + 쿼리 파라미터
func PQHandler(c *gin.Context) {
	id := c.Param("id")
	price := c.Query("price")

	c.JSON(200, gin.H{
		"message": id + "번 항목의 가격이 " + price + "원인 상품",
	})
}
