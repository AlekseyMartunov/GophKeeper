package jwt

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var ErrInvalidToken = errors.New("invalid token")

type TokenManager struct {
	tokenTTL  time.Duration
	secretKey []byte
}

type Claims struct {
	jwt.RegisteredClaims
	ID string
}

func NewTokenManager(exp time.Duration, secretKey string) *TokenManager {
	hasher := md5.New()
	hasher.Write([]byte(secretKey))

	return &TokenManager{
		tokenTTL:  exp,
		secretKey: hasher.Sum(nil),
	}
}

func (t *TokenManager) CreateToken(ID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.tokenTTL)),
		},
		ID: ID,
	})

	tokenString, err := token.SignedString(t.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *TokenManager) GetUserID(tokenString string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.secretKey, nil
	})

	if err != nil {
		return "", ErrInvalidToken
	}

	if !token.Valid {

		return "", ErrInvalidToken
	}

	return claims.ID, nil
}
