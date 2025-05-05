package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	_ "os"
)

func Connect() *sql.DB {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("$DATABASE_URL is not set")
		os.Exit(1)
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
	return db

}
