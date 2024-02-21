package users

type User struct {
	Login    string
	Password string

	ID         int
	ExternalID string
}
