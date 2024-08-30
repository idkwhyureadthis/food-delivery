package db

import (
	"database/sql"
	"fmt"

	"github.com/idkwhyureadthis/food-service/graph/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type DbData struct {
	User, Password, Host, DbName, Port string
}

var DB *sql.DB

func Setup(d DbData) error {
	connectionURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", d.User, d.Password, d.Host, d.Port, d.DbName)
	conn, err := sql.Open("pgx", connectionURL)
	if err != nil {
		return err
	}
	DB = conn
	err = Migrate()
	if err != nil {
		return err
	}
	return nil
}

func Migrate() error {
	err := goose.Up(DB, "internal/migrations")
	return err
}

func Reset() error {
	err := goose.Reset(DB, "internal/migrations")
	if err != nil {
		return err
	}
	return Migrate()
}

func AddUser(user model.User) error {
	fmt.Println(user)
	stmt, err := DB.Prepare("INSERT INTO users(name) VALUES($1)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name)
	return err
}

func Close() {
	DB.Close()
}

func GetLastId(tablename string) int64 {
	var lastId int64
	q, err := DB.Query(fmt.Sprintf("SELECT COALESCE(MAX(id), CAST(0 AS BIGINT)) as last_id FROM %s", tablename))
	if err != nil {
		panic(err)
	}
	q.Next()
	q.Scan(&lastId)
	return lastId
}
