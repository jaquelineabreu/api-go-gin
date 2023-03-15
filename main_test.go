package main

import(
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/api-go-gin/controllers"
	"github.com/jaquelineabreu/api-go-gin/database"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"fmt"
	"github.com/jaquelineabreu/api-go-gin/models"
) 

var ID int


func RotasDeTeste() *gin.Engine{
	rotas := gin.Default()

	return rotas
}

func CriaAlunoMock(){
	aluno := models.Aluno{Nome: "Nome do aluno Teste", CPF: "12345678901", RG:"098765432"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock(){
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeComParametro(t *testing.T){
	r := RotasDeTeste()
	r.GET("/:nome",controllers.Saudacao)	
	req, _ := http.NewRequest("GET","/jaqu", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)	

	//Verificando status
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API diz:":"E ai jaqu, tudo bem?"}`

	//Verificando corpo da requisição
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoAlunos(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := RotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)	
	req, _ := http.NewRequest("GET","/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	//fmt.Println(resposta.Body)
}