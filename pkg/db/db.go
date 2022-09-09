package db

import (
	"fmt"
	"log"

	"github.com/tutorials/go/crud/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	//open connection to db
	db, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("successfully connected to the databse!")
	}

    if err != nil {
        log.Fatalln(err)
    }

	db.AutoMigrate(&models.Movie{})

    return db
}
