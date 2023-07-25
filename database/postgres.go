package database

import (
	"final-project-3/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	host     = "localhost"
// 	user     = "postgres"
// 	password = "postgres"
// 	dbPort   = 5432
// 	dbname   = "final-project-3"
// )

const (
	host     = "containers-us-west-125.railway.app"
	user     = "postgres"
	password = "R1kXFyBRoSjacpC575bb"
	dbPort   = 5592
	dbname   = "railway"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, dbPort,
	)
	dsn := config
	// inject variable db
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = db.Debug().AutoMigrate(models.Task{}, models.Category{}, models.User{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Successfully connected to database")

}

func GetPostgresInstance() *gorm.DB {
	return db
}
