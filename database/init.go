package database

import (
	"log"
	"zup-message-service/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Connection *gorm.DB
var err error

func Connect(connectionString string) {
	Connection, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Connection.AutoMigrate(&model.Message{})
	log.Println("Database Migration Completed...")
}
