package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "os"
)

var db *sql.DB

const (
	host     = "aws-0-eu-central-1.pooler.supabase.com"
	port     = 5432
	user     = "postgres.gilvehfffvzllxcicizz"
	password = "qifvyh-japjyz-teXfi1"
	dbname   = "postgres"
)

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db

}
