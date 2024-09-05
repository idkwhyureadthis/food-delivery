package encoder

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	errTokenModified = errors.New("modified token provided")
	errWrongToken    = errors.New("wrong token provided")
	errTokenExpired  = errors.New("token lifetime expired")
)

func Encode(header model.Header, payload model.Body, secretKey []byte) (string, error) {
	h := hmac.New(sha256.New, secretKey)
	hBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	hEnc := base64.RawURLEncoding.EncodeToString(hBytes)
	bBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	bEnc := base64.RawURLEncoding.EncodeToString(bBytes)
	signatureString := hEnc + "." + bEnc
	h.Write([]byte(signatureString))
	sEnc := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	EncodedData := hEnc + "." + bEnc + "." + sEnc
	return EncodedData, nil
}

func Decode(access string, secret []byte) (*model.Body, error) {
	var headerStruct model.Header
	var payloadStruct model.Body
	tokenParts := strings.Split(access, ".")
	if len(tokenParts) != 3 {
		return nil, errWrongToken
	}
	header, payload, signature := tokenParts[0], tokenParts[1], tokenParts[2]
	headerDecoded, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		return nil, err
	}
	payloadDecoded, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(headerDecoded, &headerStruct)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(payloadDecoded, &payloadStruct)
	if err != nil {
		return nil, err
	}
	expectedEncoding, err := Encode(headerStruct, payloadStruct, secret)
	if err != nil {
		return nil, err
	}
	if strings.Split(expectedEncoding, ".")[2] != signature {
		return nil, errTokenModified
	}
	if payloadStruct.Exp < time.Now().Unix() {
		return nil, errTokenExpired
	}
	return &payloadStruct, nil
}

func CryptToken(token string) (string, error) {
	cryptedRefresh, err := bcrypt.GenerateFromPassword([]byte(token), 14)
	if err != nil {
		return "", err
	}
	return string(cryptedRefresh), nil
}
