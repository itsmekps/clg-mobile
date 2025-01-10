package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySqlDB() (*sql.DB, error) {
	dsn := "fiber_user:fiber_password@tcp(localhost:3306)/fiber_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Ping Databse
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected to MySQL")
	// Create table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
