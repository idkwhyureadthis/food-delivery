package tokengen

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/internal/database/db"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/encoder"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

func GenerateTokens(userData model.UserData, secretKey []byte) (*model.GeneratedTokens, error) {
	tokens := model.GeneratedTokens{}
	accessHeader := model.Header{Alg: "HS256",
		Typ: "JWT"}
	accessBody := model.Body{
		Sub:  userData.Id,
		Name: userData.Name,
		Exp:  time.Now().Unix(),
	}
	access, err := encoder.Encode(accessHeader, accessBody, secretKey)
	if err != nil {
		return nil, err
	}
	hs := hmac.New(sha256.New, secretKey)
	hs.Write([]byte(access))
	refresh := string(hs.Sum(nil))
	tokens.AccessToken = access
	tokens.RefreshToken = fmt.Sprint(userData.Id) + "." + fmt.Sprintf("%x", refresh)
	return &tokens, nil
}

func FromAccess(refresh string, secret []byte) (*model.GeneratedTokens, error) {
	idString := strings.Split(refresh, ".")[0]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return nil, err
	}
	cryptedRefresh, name, err := db.GetRefreshAndName(id)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(cryptedRefresh), []byte(refresh)); err != nil {
		return nil, errors.New("wrong token provided")
	}
	return GenerateTokens(model.UserData{Name: name, Id: id}, secret)
}
