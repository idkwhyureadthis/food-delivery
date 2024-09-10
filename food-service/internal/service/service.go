package service

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/idkwhyureadthis/food-service/graph/model"
	"github.com/idkwhyureadthis/food-service/internal/db"
)

func CreateNewUser(userName, password, about string) (*model.User, error) {
	h := sha256.New()
	h.Write([]byte(password))
	cryptedPassword := fmt.Sprintf("%x", h.Sum(nil))
	id, tokens, err := db.AddUser(userName, string(cryptedPassword), about)
	if err != nil {
		return nil, err
	}
	newUser := model.User{
		ID:     id,
		Name:   userName,
		About:  about,
		Orders: []*model.Order{},
		Tokens: tokens,
	}
	return &newUser, nil
}

func GetUser(id int64) (*model.User, error) {
	userData, err := db.GetUser(id)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func DeleteUser(ctx context.Context) (string, error) {
	return "deletion was successful", nil
}
