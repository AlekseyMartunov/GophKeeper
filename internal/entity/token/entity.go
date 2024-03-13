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
)

type Token struct {
	TokenID        int
	Name           string
	Token          string
	ExternalUserID string
	UserID         int
	IsBlocked      bool
	ExpireTime     time.Time
	CreatedTime    time.Time
}
