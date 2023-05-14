package database

import (
	"log"

	"github.com/WeslleySanto/go-validacoes-testes-paginas-html/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Conecta() {
	dns := "./alunos.db"
	DB, err = gorm.Open(sqlite.Open(dns))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
