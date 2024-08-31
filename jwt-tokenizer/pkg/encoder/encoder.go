package encoder

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
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
