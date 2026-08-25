package handler

import (
	"signUp/modules"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var sign modules.SignUpRequest
	if err := c.ShouldBindJSON(&sign); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if len(sign.Password) <= 4 {
		c.JSON(400, gin.H{"error": "에러 비밀번호를 5자 이상으로 지어주세요"})
		return
	}
	if sign.Email == "" {
		c.JSON(400, gin.H{"error": "에러 이메일을 작성해주세요"})
		return
	}
	if sign.Username == "" {
		c.JSON(400, gin.H{"error": "에러 유저네임을 작성하세요"})
		return
	}
	c.JSON(201, gin.H{
		"success": true,
	})
}

func LoginHandler(c *gin.Context) {
	var login modules.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if login.Username != "admin" || login.Password != "12345" {
		c.JSON(401, gin.H{"error": "로그인 실패"})
		return
	}
	c.JSON(200, gin.H{"success": true})
}
