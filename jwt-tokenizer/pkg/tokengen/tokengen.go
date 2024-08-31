package tokengen

import (
	"fmt"
	"time"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/encoder"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
)

type TokenPart interface {
}

func GenerateTokens(userData model.UserData, secretKey []byte) (*model.GeneratedTokens, error) {
	tokens := model.GeneratedTokens{}
	headerStruct := model.Header{Alg: "HS256",
		Typ: "JWT"}
	bodyStruct := model.Body{
		Sub:  userData.Id,
		Name: userData.Name,
		Exp:  time.Now().Add(15 * time.Minute).Unix(),
	}
	fmt.Println(headerStruct, bodyStruct)
	access, err := encoder.Encode(headerStruct, bodyStruct, secretKey)
	if err != nil {
		return nil, err
	}
	tokens.AccessToken = access
	return &tokens, nil
}
