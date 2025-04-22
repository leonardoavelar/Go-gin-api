package database

import (
	"log"

	"github.com/leonardoavelar/Go-gin-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Conectar() {

	connectionString := "host=localhost user=leonardo password=leonardo dbname=leonardo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Erro ao conectar no banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})

}
