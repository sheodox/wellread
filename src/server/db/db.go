package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sheodox/wellread/query"
)

var Queries *query.Queries

func ConnectionString() string {
	pgUser := os.Getenv("PGUSER")
	pgPassword := os.Getenv("PGPASSWORD")
	pgHost := os.Getenv("PGHOST")
	pgDatabase := os.Getenv("PGDATABASE")
	return fmt.Sprintf("postgres://%v:%v@%v:5432/%v?sslmode=disable", pgUser, pgPassword, pgHost, pgDatabase)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db, err := sql.Open("postgres", ConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	Queries = query.New(db)
}
