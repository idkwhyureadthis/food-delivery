package service

import (
	"database/sql"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/internal/database/db"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
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

func (s *Service) CreateUser() model.ServiceResponse {
	resp := model.ServiceResponse{}
	return resp
}
