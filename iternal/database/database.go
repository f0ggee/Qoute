package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	_ "os"
)

var db *sql.DB

const (
	host     = "aws-0-eu-central-1.pooler.supabase.com"
	port     = 5432
	user     = "postgres.gilvehfffvzllxcicizz"
	password = "0208082n2amjej"
	dbname   = "postgres"
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
