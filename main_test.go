package main

import(
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/api-go-gin/controllers"
	"net/http"
	"net/http/httptest"


) 


func RotasDeTeste() *gin.Engine{
	rotas := gin.Default()

	return rotas
}

func TestVerificaStatusCodeComParametro(t *testing.T){
	r := RotasDeTeste()
	r.GET("/:nome",controllers.Saudacao)	
	req, _ := http.NewRequest("GET","/", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	if resposta.Code != http.StatusOK{
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	}


}