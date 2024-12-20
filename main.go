package main

import (
	"database/sql"
	"github.com/ahay12/api-test/database"
	"github.com/ahay12/api-test/router"
	"log"
)

func main() {
	db := database.InitDatabase()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal("Failed to close database connection")
		}
	}(sqlDB)
	app := router.Make()
	log.Fatal(app.Listen(":4000"))
}
