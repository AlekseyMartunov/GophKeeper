package tokenhandlers

type tokenDTO struct {
	Token string `json: "token"`
}

type clientsDTO struct {
	Clients []string `json:"clients"`
}

type tokenDeleteDTO struct {
	TokenName string `json:"token_name"`
}

type createTokenDTO struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	TokenName string `json:"name"`
}
