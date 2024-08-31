package model

type GeneratedTokens struct {
	AccessToken, RefreshToken string
}

type UserData struct {
	Id   int64
	Name string
}

type ServiceResponse struct {
	Answer string
	Err    error
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
