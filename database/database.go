package database

import (
	"fmt"
	"honeypot/models"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
    var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

    DB.Debug().AutoMigrate(&models.User{})
}
