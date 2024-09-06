package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

var DB *sql.DB

func InitDatabase(connUrl string) *sql.DB {
	conn, err := sql.Open("pgx", connUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = SetupMigrations(conn)
	if err != nil {
		log.Fatal(err)
	}
	DB = conn
	return conn
}

func SetupMigrations(db *sql.DB) error {
	return goose.Up(db, "internal/database/migrations")
}
