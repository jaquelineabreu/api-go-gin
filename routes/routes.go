package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/api-go-gin/controllers"
)



func HandleRequests(){
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.Run(":8000")
}