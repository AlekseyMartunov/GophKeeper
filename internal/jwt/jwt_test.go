package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const id = "123456"

func TestCreatingToken(t *testing.T) {
	manager := NewTokenManager(time.Second*10, "some key")

	token, err := manager.CreateToken(id)
	assert.NoError(t, err,
		"Error creating token")

	idFromToken, err := manager.GetUserID(token)
	assert.NoError(t, err,
		"Error parsing token")

	assert.Equal(t, id, idFromToken,
		"id dose not math with original id")

	_, err = manager.GetUserID(token[:len(token)-4])
	assert.Equal(t, err, ErrInvalidToken,
		"token manager dose not return invalidTokenError")

}

func TestTimeToLeaveToken(t *testing.T) {
	manager := NewTokenManager(time.Second*1, "some key")
	token, err := manager.CreateToken(id)
	assert.NoError(t, err,
		"Error creating token")

	time.Sleep(time.Second * 2)

	_, err = manager.GetUserID(token)
	assert.Equal(t, err, ErrExpiredToken)
}
