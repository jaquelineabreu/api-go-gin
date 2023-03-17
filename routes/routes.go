package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/api-go-gin/controllers"
)



func HandleRequests(){
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets","./assets")
	r.GET("/:nome",controllers.Saudacao)	
	r.GET("/alunos", controllers.ExibeTodosAlunos)	
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run(":8080")
}