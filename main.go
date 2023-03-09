package main

import (
	"github.com/jaquelineabreu/api-go-gin/database"
	"github.com/jaquelineabreu/api-go-gin/models"
	"github.com/jaquelineabreu/api-go-gin/routes"
)



func main() {	
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome:"Jaqueline Abreu",CPF:"00000000000000", RG:"0100000X"},
		{Nome:"Maria Oliveira",CPF:"00000000000001", RG:"0200000X"},
	}

	routes.HandleRequests()
	
}