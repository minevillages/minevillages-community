package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Postgres() {

	// Construct connection string using environment variables
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_SSL_MODE"),
	)

	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Ping the database to verify connection
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database!")

}
