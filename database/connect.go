package database

import (
	"fmt"
	"log"

	"github.com/plusiv/fiber-todo-api/config"
	"github.com/plusiv/fiber-todo-api/internals/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	/* Uncomment this if using a database different from sqlite
	portStr := config.Config("DB_PORT")
	port, err := strconv.ParseInt(portStr, 10, 32);

	if err != nil {
		log.Fatalf("Unable to convert Port %s to int", portStr)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err := gorm.Open(sqlite.Open(dsn))
	*/

	DB, err = gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", config.Config("DB_NAME"))))

	if err != nil {
		log.Fatal("Unable to connect with Database.")
	}

	log.Print("Connected to Database")
	DB.AutoMigrate(&model.ToDo{})
	log.Print("Database migrated")

}
