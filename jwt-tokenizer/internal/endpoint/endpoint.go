package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RefreshTokensInput struct {
	Token string `json:"refresh"`
}

type VerifyTokensInput struct {
	Access     string `json:"access"`
	Refresh    string `json:"refresh"`
	AccessedId int64  `json:"access_to"`
}

type Service interface {
	CreateUser(userName, password string) (*model.GeneratedTokens, error)
	RegenerateTokens(refresh string) (*model.GeneratedTokens, error)
	Verify(refresh, access string, accessedId int64) (*model.ServiceResponse, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{s: s}
}

func (e *Endpoint) CreateUser(w http.ResponseWriter, r *http.Request) {
	input := CreateUserInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := e.s.CreateUser(input.Name, input.Password)
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
	resp, err := e.s.RegenerateTokens(input.Token)
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
	resp, err := e.s.Verify(input.Refresh, input.Access, input.AccessedId)
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
