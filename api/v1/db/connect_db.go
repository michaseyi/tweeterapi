package db

import (
	"database/sql"
	"fmt"
	"log"
	"tweeter/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectDB connectes to the database
// Exits if it fails to connect
func ConnectDB() {
	db, err := sql.Open(config.DATABASE_NAME, config.DATABASE_URI)
	DB = db
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connected")
}
