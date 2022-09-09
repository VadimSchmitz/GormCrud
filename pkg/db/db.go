package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tutorials/go/crud/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	host := os.Getenv("HOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	
	// create db connection str
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, port)
	//open connection to db
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})


	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("successfully connected to the databse!")
	}

    if err != nil {
        log.Fatalln(err)
    }

	//Make db migrations if they do not excist
	// db.AutoMigrate(&Person{})
	// db.AutoMigrate(&Book{})
	db.AutoMigrate(&models.Book{})


    return db
}
