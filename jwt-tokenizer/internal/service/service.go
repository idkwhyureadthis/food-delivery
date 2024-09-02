package service

import (
	"crypto/sha256"
	"database/sql"
	"fmt"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/internal/database/db"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/tokengen"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	secretKey []byte
	db        *sql.DB
}

func New(dbPath string, secretKey []byte) *Service {
	conn := db.InitDatabase(dbPath)
	return &Service{
		secretKey: secretKey,
		db:        conn,
	}
}

func (s *Service) CreateUser(userName, password string) (*model.GeneratedTokens, error) {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))
	id, err := db.CreateUser(userName, hashedPassword)
	if err != nil {
		return nil, err
	}
	userData := model.UserData{
		Id:   id,
		Name: userName,
	}
	tokens, err := tokengen.GenerateTokens(userData, s.secretKey)
	if err != nil {
		return nil, err
	}
	cryptedRefresh, err := bcrypt.GenerateFromPassword([]byte(tokens.RefreshToken), 14)
	if err != nil {
		return nil, err
	}
	err = db.SetKey(id, string(cryptedRefresh))
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
