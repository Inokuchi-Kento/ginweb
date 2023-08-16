package controller

import "github.com/gin-gonic/gin"

type Hello interface {
	HelloController(c *gin.Context)
}

type hello struct{}

func NewHelloController() Hello {
	return &hello{}
}

func (h *hello) HelloController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
