package infra

import (
	"log"

	"github.com/lucasSSalgado/firstCRUD.git/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=example dbname=stock port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("falha ao se conectar a DB ", err)
	}

	err = db.AutoMigrate(&model.Stock{})
	if err != nil {
		log.Fatal("falha ao migrar database ", err)
	}
	return db
}
