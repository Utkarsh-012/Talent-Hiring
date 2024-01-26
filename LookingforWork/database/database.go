package database

import (
	"database/sql"

	"fmt"
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

func CreateTable() {

	// Create the table
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS candidate (
		id INT AUTO_INCREMENT PRIMARY KEY,
		fullname VARCHAR(100) NOT NULL,
		about VARCHAR(100),
		jobtitle VARCHAR(100),
		experience VARCHAR(100),
		country VARCHAR(100)

	);`
	_, err := Db.Exec(createTableQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'candidate' created successfully.")

}

func CreateTableJob() {
	createTableQuery := `
CREATE TABLE IF NOT EXISTS job (
	id INT AUTO_INCREMENT PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	jobName VARCHAR(255) NOT NULL
);`
	_, err := Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'job' created successfully.")
}

func CreateTableProfile() {
	createTableQuery := `
CREATE TABLE IF NOT EXISTS profile (
	id INT AUTO_INCREMENT PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	profilelink VARCHAR(255) NOT NULL
);`
	_, err := Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'profile' created successfully.")
}

func CreateTableSkills() {
	createTableQuery := `
CREATE TABLE IF NOT EXISTS skills (
	id INT AUTO_INCREMENT PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	skillsname VARCHAR(255) NOT NULL
);`
	_, err := Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'skills' created successfully.")
}

func CreateTableSocialMedia() {
	createTableQuery := `
CREATE TABLE IF NOT EXISTS socialmedia (
	id INT AUTO_INCREMENT PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	platform VARCHAR(255) NOT NULL,
	url VARCHAR(255) NOT NULL
);`
	_, err := Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'socialmedia' created successfully.")
}

func CreateTableLocation() {
	createTableQuery := `
CREATE TABLE IF NOT EXISTS location (
	id INT AUTO_INCREMENT PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	locationname VARCHAR(255) NOT NULL,
	locationtype VARCHAR(255) NOT NULL
);`
	_, err := Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'location' created successfully.")
}
