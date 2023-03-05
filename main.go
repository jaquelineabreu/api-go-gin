package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	
	//subir servidor
	r := gin.Default()
	r.Run(":5000")
}