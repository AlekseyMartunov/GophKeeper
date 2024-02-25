package userhandlers

const (
	userAlreadyExists   = "this login already used by another user"
	userDoseNotExist    = "user with this login and password dose not exists"
	internalServerError = "internal server error"
	requestParsingError = "the request form is incorrect or the request does not contain the required field"
	invalidToken        = "your token is invalid"
	expireToken         = "token has expired"
)
