package service

import (
	"context"
	"fmt"
	"log"

	"github.com/idkwhyureadthis/food-service/graph/model"
	"github.com/idkwhyureadthis/food-service/internal/db"
)

var lastAddedUser, lastAddedOrder, lastAddedProduct int64

func Init() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	lastAddedUser = db.GetLastId("users")
	lastAddedOrder = db.GetLastId("orders")
	lastAddedProduct = db.GetLastId("products")
	fmt.Println(lastAddedOrder, lastAddedProduct, lastAddedUser)
}

func CreateNewUser(userName string) (*model.User, error) {
	lastAddedUser++
	newUser := model.User{
		ID:     lastAddedUser,
		Name:   userName,
		Orders: []*model.Order{},
	}
	err := db.AddUser(newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func DeleteUser(ctx context.Context) (string, error) {

	return "deletion was successful", nil
}
