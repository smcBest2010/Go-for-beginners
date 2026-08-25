package handler

import (
	"post/module"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var post module.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, post)
}

func SignUpHandler(c *gin.Context) {
	var sign module.SignUpRequest
	if err := c.ShouldBindJSON(&sign); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, sign)
}

/*
택배(JSON)가 서버에 도착했습니다.
고랭은 택배 상자 안에 든 데이터가 필요합니다.
하지만 고랭 코드 안에서는 {"title": "안녕"} 같은 문자열 상태로는 데이터를 마음대로 요리(DB 저장 등)할 수 없습니다.
그래서 고랭이 다룰 수 있는 고랭 전용 상자(구조체 변수)로 내용물을 옮겨 담아야 합니다.
*/
