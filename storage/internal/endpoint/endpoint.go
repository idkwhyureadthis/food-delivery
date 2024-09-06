package endpoint

import "net/http"

type Service interface {
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{s: s}
}

func (e *Endpoint) CreateUser(w http.ResponseWriter, r *http.Request) {
}

func (e *Endpoint) GetUser(w http.ResponseWriter, r *http.Request) {
}
