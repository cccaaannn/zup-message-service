package database

import (
	"log"
	"zup-message-service/data/models"

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
		panic("[PostgreSQL] Cannot connect to DB")
	}
	log.Printf("[PostgreSQL] Connected to DB with %s", connectionString)
}

func Migrate() {
	log.Println("[PostgreSQL] Database migration started")
	Connection.AutoMigrate(&models.Message{})
	log.Println("[PostgreSQL] Database migration completed")
}
