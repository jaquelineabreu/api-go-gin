package controllers

import (
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context){
	c.JSON(200, gin.H{
		"id":"1",
		"nome":"Jaqueline Abreu",
	})
}