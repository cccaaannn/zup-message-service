package database

import (
	"log"
	"zup-message-service/models"

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
	log.Printf("Connected to DB with %s", connectionString)
}

func Migrate() {
	Connection.AutoMigrate(&models.Message{})
	Connection.AutoMigrate(&models.UserOnlineStatus{})
	log.Println("Database migration completed...")
}
