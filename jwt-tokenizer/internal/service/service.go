package service

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/internal/database/db"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/encoder"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/tokengen"
	"golang.org/x/crypto/bcrypt"
)

var (
	errWrongToken = errors.New("token for wrong person provided")
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
	cryptedRefresh, err := encoder.CryptToken(tokens.RefreshToken)
	if err != nil {
		return nil, err
	}
	err = db.SetKey(id, string(cryptedRefresh))
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (s *Service) RegenerateTokens(refresh string) (*model.GeneratedTokens, error) {
	newTokens, err := tokengen.FromAccess(refresh, s.secretKey)
	if err != nil {
		return nil, err
	}
	cryptedRefresh, err := bcrypt.GenerateFromPassword([]byte(newTokens.RefreshToken), 14)
	if err != nil {
		return nil, err
	}
	idString := strings.Split(newTokens.RefreshToken, ".")[0]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return nil, err
	}
	db.SetKey(id, string(cryptedRefresh))
	return newTokens, err
}

func (s *Service) Verify(refresh, access string, accessedId int64) (*model.ServiceResponse, error) {
	resp := model.ServiceResponse{}
	body, err := encoder.Decode(access, s.secretKey)
	if err != nil {
		if err.Error() == "token lifetime expired" {
			newTokens, err := tokengen.FromAccess(refresh, s.secretKey)
			if err != nil {
				return nil, err
			}
			resp.NewTokens = newTokens
			access := resp.NewTokens.AccessToken
			body, err = encoder.Decode(access, s.secretKey)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	if body.Sub != accessedId {
		return nil, errWrongToken
	}
	resp.Message = "Verification Successful"
	return &resp, nil
}
