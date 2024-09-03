package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

var lastCreatedUser int
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

func CreateUser(name, hashedPassword string) (int64, error) {
	var id int64
	stmt, err := DB.Prepare("INSERT INTO users (name, hashed_password, refresh_token) VALUES ($1, $2, '') RETURNING id;")
	if err != nil {
		return -1, err
	}
	resp, err := stmt.Query(name, hashedPassword)
	if err != nil {
		return -1, err
	}
	resp.Next()
	resp.Scan(&id)
	return id, nil
}

func SetKey(id int64, refresh string) error {
	_, err := DB.Exec("UPDATE USERS SET refresh_token = $1 where id = $2", refresh, id)
	if err != nil {
		DeleteUser(id)
		return err
	}
	return nil
}

func DeleteUser(id int64) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetRefreshAndName(id int64) (string, string, error) {
	var token, name string
	rows, err := DB.Query("SELECT refresh_token, name FROM users WHERE id = $1", id)
	if err != nil {
		return "", "", err
	}
	found := rows.Next()
	if !found {
		return "", "", sql.ErrNoRows
	}
	rows.Scan(&token, name)
	return token, name, nil
}
