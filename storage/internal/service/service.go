package service

import (
	"database/sql"
	"errors"

	"github.com/idkwhyureadthis/food-delivery/storage/internal/database/db"
)

var (
	errWrongToken = errors.New("token for wrong person provided")
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
