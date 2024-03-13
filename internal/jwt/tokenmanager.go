package jwt

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var ErrInvalidToken = errors.New("invalid token")
var ErrExpiredToken = errors.New("token is expired")

type tokenService interface {
	GetUserIDByExternalID(ctx context.Context, externalID string) (int, error)
}

type TokenManager struct {
	tokenTTL  time.Duration
	secretKey []byte
	service   tokenService
}

type Claims struct {
	jwt.RegisteredClaims
	ID string
}

func NewTokenManager(exp time.Duration, secretKey string, service tokenService) *TokenManager {
	hasher := md5.New()
	hasher.Write([]byte(secretKey))

	return &TokenManager{
		tokenTTL:  exp,
		secretKey: hasher.Sum(nil),
		service:   service,
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

func (t *TokenManager) GetUserID(ctx context.Context, tokenString string) (int, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.secretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, ErrExpiredToken
		}
		return 0, ErrInvalidToken
	}

	if !token.Valid {
		return 0, ErrInvalidToken
	}
	id, err := t.service.GetUserIDByExternalID(ctx, claims.ID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
