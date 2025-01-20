package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() bool {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connectionString := "postgres" + "://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?" + "sslmode=disable"
	fmt.Println("connectionString is: ", connectionString)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("error in connecting to DB", err)
		return false
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users(
	userID SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL
	);
	`

	if _, err = db.Exec(createUserTable); err != nil {
		log.Fatal("error in creating users table", err)
		return false
	}

	createTaskTable := `
	CREATE TABLE IF NOT EXISTS tasks(
	taskID SERIAL PRIMARY KEY,
	userID int,
	taskName VARCHAR(255) NOT NULL,
	status bool,
	FOREIGN KEY(userID) references users(userID)
	);
	`

	if _, err = db.Exec(createTaskTable); err != nil {
		log.Fatal("error in creating tasks table", err)
		return false
	}

	DB = db
	return true
}
