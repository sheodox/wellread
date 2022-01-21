package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", ConnectionString())
}

func ConnectionString() string {
	pgUser := os.Getenv("PGUSER")
	pgPassword := os.Getenv("PGPASSWORD")
	pgHost := os.Getenv("PGHOST")
	pgDatabase := os.Getenv("PGDATABASE")
	return fmt.Sprintf("postgres://%v:%v@%v:5432/%v?sslmode=disable", pgUser, pgPassword, pgHost, pgDatabase)
}
