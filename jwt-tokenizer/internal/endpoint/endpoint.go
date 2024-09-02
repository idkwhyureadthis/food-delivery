package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/pkg/model"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Service interface {
	CreateUser(userName, password string) (*model.GeneratedTokens, error)
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
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("failed to create user: %v", err)))
		return
	}
	dat, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("failed to create user: %v", err)))
		return
	}
	w.WriteHeader(200)
	w.Write(dat)
}
