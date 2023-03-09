package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/api-go-gin/database"
	"github.com/jaquelineabreu/api-go-gin/models"
)

func ExibeTodosAlunos(c *gin.Context){
	var alunos []models.Aluno
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context){
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo bem?",
	})
}

func CriaNovoAluno(c *gin.Context){
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}