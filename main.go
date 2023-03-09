package main

import (
	"github.com/jaquelineabreu/api-go-gin/database"
	"github.com/jaquelineabreu/api-go-gin/routes"
)



func main() {	
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}