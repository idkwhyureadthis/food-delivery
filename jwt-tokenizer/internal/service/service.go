package service

import (
	"database/sql"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/internal/database/db"
)

type Service struct {
	db *sql.DB
}

func New(dbPath string) *Service {
	conn := db.InitDatabase(dbPath)
	return &Service{
		db: conn,
	}
}
