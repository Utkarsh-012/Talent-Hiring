package database

import (
	"database/sql"
	"fmt"

	//"project/LookingforWork/database"

	// "fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "root:qwerty123@tcp(localhost:3306)/talent")
	if err != nil {
		log.Fatal(err)
	}

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
}

func CreateTableCompany() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS companies (
		id VARCHAR(255) PRIMARY KEY,
		profile VARCHAR(255) NOT NULL,
		companyname VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		verifiedemail VARCHAR(255) NOT NULL,
		companywebsite VARCHAR(255) NOT NULL
	);`
	_, err := Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'companies' created successfully.")
}
