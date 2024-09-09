package service

import (
	"database/sql"
	"errors"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/encoder"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/tokengen"
)

var (
	errWrongToken = errors.New("token for wrong person provided")
)

type Service struct {
	secretKey []byte
	db        *sql.DB
}

func New(dbPath string, secretKey []byte) *Service {
	return &Service{
		secretKey: secretKey,
	}
}

func (s *Service) GenerateNewTokens(userName string, id int64) (*model.GeneratedTokens, error) {
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
	tokens.CryptedRefresh = cryptedRefresh
	return tokens, nil
}

func (s *Service) RegenerateTokens(refresh, cryptedRefresh, name string) (*model.GeneratedTokens, error) {
	newTokens, err := tokengen.FromRefresh(refresh, cryptedRefresh, name, s.secretKey)
	if err != nil {
		return nil, err
	}
	newCryptedRefresh, err := encoder.CryptToken(newTokens.RefreshToken)
	if err != nil {
		return nil, err
	}
	newTokens.CryptedRefresh = string(newCryptedRefresh)
	return newTokens, err
}

func (s *Service) Verify(refresh, access, cryptedRefresh, name string, accessedId int64) (*model.ServiceResponse, error) {
	resp := model.ServiceResponse{}
	body, err := encoder.Decode(access, s.secretKey)
	if err != nil {
		if err.Error() == "token lifetime expired" {
			newTokens, err := tokengen.FromRefresh(refresh, cryptedRefresh, name, s.secretKey)
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
