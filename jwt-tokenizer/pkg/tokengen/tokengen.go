package tokengen

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/encoder"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
)

type TokenPart interface {
}

func GenerateTokens(userData model.UserData, secretKey []byte) (*model.GeneratedTokens, error) {
	tokens := model.GeneratedTokens{}
	accessHeaderStruct := model.Header{Alg: "HS256",
		Typ: "JWT"}
	accessBodyStruct := model.Body{
		Sub:  userData.Id,
		Name: userData.Name,
		Exp:  time.Now().Add(15 * time.Minute).Unix(),
	}
	access, err := encoder.Encode(accessHeaderStruct, accessBodyStruct, secretKey)
	if err != nil {
		return nil, err
	}
	hs := hmac.New(sha256.New, secretKey)
	hs.Write([]byte(access))
	refresh := string(hs.Sum(nil))
	tokens.AccessToken = access
	tokens.RefreshToken = fmt.Sprintf("%x", refresh)
	return &tokens, nil
}
