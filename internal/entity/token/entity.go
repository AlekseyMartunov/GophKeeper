package token

import (
	"errors"
	"time"
)

var (
	ErrTokenIsExpired     = errors.New("token is expired")
	ErrTokenAlreadyExists = errors.New("token already exist")
	ErrTokenIsInvalid     = errors.New("invalid token")
	ErrTokenIsBlocked     = errors.New("token blocked")
	ErrNoTokenFound       = errors.New("no such token exists")
)

type Token struct {
	TokenID        int
	Name           string
	Token          string
	ExternalUserID string
	InternalUserID int
	IsBlocked      bool
	CreatedTime    time.Time
}
