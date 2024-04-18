package authenticationhttp

import (
	"GophKeeper/internal/jwt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpitToken(t *testing.T) {
	testCase := []struct {
		name     string
		tokenIN  string
		tokenOUT string
		err      error
	}{
		{
			name:     "test_1",
			tokenIN:  "eyJhbGci.OiJIUzI1NiIXV.CJ9TJV",
			tokenOUT: "",
			err:      jwt.ErrInvalidToken,
		},
		{
			name:     "test_2",
			tokenIN:  "Bearer eyJhbGci.OiJIUzI1NiIXV.CJ9TJV",
			tokenOUT: "eyJhbGci.OiJIUzI1NiIXV.CJ9TJV",
		},
		{
			name:     "test_3",
			tokenIN:  "Bearer eyJhbGciOiJIUzI1.NiIXVCJ9TJV",
			tokenOUT: "",
			err:      jwt.ErrInvalidToken,
		},
		{
			name:     "test_4",
			tokenIN:  "bearer eyJhbGci.OiJIUzI1NiIXV.CJ9TJV",
			tokenOUT: "",
			err:      jwt.ErrInvalidToken,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			token, err := splitToken(tc.tokenIN)

			assert.Equal(t, tc.tokenOUT, token)
			assert.Equal(t, tc.err, err)

		})
	}
}
