package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
)

type GenerateTokensInput struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

type RefreshTokensInput struct {
	Token          string `json:"refresh"`
	CryptedRefresh string `json:"crypted_refresh"`
	Name           string `json:"name"`
}

type VerifyTokensInput struct {
	Access         string `json:"access"`
	Refresh        string `json:"refresh"`
	CryptedRefresh string `json:"crypted_refresh"`
	Name           string `json:"name"`
	AccessedId     int64  `json:"access_to"`
}

type Service interface {
	GenerateNewTokens(userName string, id int64) (*model.GeneratedTokens, error)
	RegenerateTokens(refresh, cryptedRefresh, name string) (*model.GeneratedTokens, error)
	Verify(refresh, access, cryptedRefresh, name string, accessedId int64) (*model.ServiceResponse, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{s: s}
}

func (e *Endpoint) GenerateTokens(w http.ResponseWriter, r *http.Request) {
	input := GenerateTokensInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := e.s.GenerateNewTokens(input.Name, input.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dat, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(dat)
}

func (e *Endpoint) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	input := RefreshTokensInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := e.s.RegenerateTokens(input.Token, input.CryptedRefresh, input.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
	w.Write(data)
}

func (e *Endpoint) VerifyUser(w http.ResponseWriter, r *http.Request) {
	input := VerifyTokensInput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := e.s.Verify(input.Refresh, input.Access, input.CryptedRefresh, input.Name, input.AccessedId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
	w.Write(data)
}
