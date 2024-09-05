package model

type GeneratedTokens struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

type UserData struct {
	Id   int64
	Name string
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Body struct {
	Sub  int64  `json:"sub"`
	Name string `json:"name"`
	Exp  int64  `json:"exp"`
}

type ServiceResponse struct {
	Message   string           `json:"message"`
	NewTokens *GeneratedTokens `json:"tokens"`
}
