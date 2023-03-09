package database

import (
	"log"

	"github.com/jaquelineabreu/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	DB *gorm.DB
	err error
)
  
func ConectaComBancoDeDados(){
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if  err != nil {
		log.Panic("Erro ao conectar com o banco")
	}

	//cria tabela no banco de dados
	DB.AutoMigrate(&models.Aluno{})
}

