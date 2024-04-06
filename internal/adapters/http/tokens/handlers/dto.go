package tokenhandlers

type tokenDTO struct {
	Token string `json: "token"`
}

type tokenDeleteDTO struct {
	TokenName string `json:"token_name"`
}

type createTokenDTO struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	TokenName string `json:"token_name"`
}
