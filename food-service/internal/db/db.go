package db

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/idkwhyureadthis/food-service/graph/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

func Setup(connectionURL string) error {
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

func AddUser(userName, cryptedPassword string) (int64, *model.Tokens, error) {
	var tokens model.TokensWithCrypted
	var addedId int64
	tx, err := DB.BeginTx(context.Background(), nil)
	if err != nil {
		return -1, nil, err
	}
	defer tx.Rollback()
	if err = tx.QueryRow("INSERT INTO users(name, hashed_password) VALUES($1, $2) RETURNING id", userName, cryptedPassword).Scan(&addedId); err != nil {
		return -1, nil, err
	}
	genData := struct {
		Name string `json:"name"`
		Id   int64  `json:"id"`
	}{
		Name: userName,
		Id:   addedId,
	}
	marshalledData, err := json.Marshal(genData)
	if err != nil {
		return -1, nil, err
	}
	resp, err := http.Post("http://localhost:8081/generate", "application/json", bytes.NewReader(marshalledData))
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, err
	}
	err = json.Unmarshal(respBody, &tokens)
	if err != nil {
		return -1, nil, err
	}
	_, err = tx.Exec("UPDATE USERS SET refresh_token = $1 WHERE id = $2", tokens.CryptedRefresh, addedId)
	if err != nil {
		return -1, nil, err
	}

	if err := tx.Commit(); err != nil {
		return -1, nil, err
	}
	tokensToReturn := model.Tokens{
		Refresh: tokens.Refresh,
		Access:  tokens.Access,
	}
	return addedId, &tokensToReturn, err
}

func Close() {
	DB.Close()
}
