package main

import (
	"database/sql"
	"fmt"
	"log"
"github.com/google/uuid"
	_ "github.com/go-sql-driver/mysql"
)

const DATABASE_URI = "michaseyi:&KqEMEbp,67UygW@tcp(127.0.0.1:3306)/alx"
const DATABASE_NAME = "mysql"

type Student struct {
	first_name string
	last_name  string
	email      string
}

func main() {
	fmt.Println(uuid.NewUUID())
	db, err := sql.Open(DATABASE_NAME, DATABASE_URI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connected")
	}

	result := db.QueryRow(`SELECT first_name , last_name, email FROM student WHERE email = ?;`, "Tade5529@bing.com")

	var student Student
	if err = result.Scan(&student.first_name, &student.last_name, &student.email); err != nil {
		log.Fatal(err)
	}
	fmt.Println(student)

}
