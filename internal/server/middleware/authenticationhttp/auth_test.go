package authenticationhttp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"GophKeeper/internal/server/jwt"
)

func TestSpitToken(t *testing.T) {
	testCase := []struct {
		name   string
		header string
		token  string
		err    error
	}{
		{
			name:   "test_1",
			header: "Authorization: Bearer eyJhbGci.OiJIUzI1NiIXV.CJ9TJV",
			token:  "eyJhbGci.OiJIUzI1NiIXV.CJ9TJV",
		},
		{
			name:   "test_2",
			header: "Authorization : Bearer eyJhbGci.OiJIUzI1NiIXV.CJ9TJV",
			token:  "",
			err:    jwt.ErrInvalidToken,
		},
		{
			name:   "test_3",
			header: "Authorization : Bearer eyJhbGciOiJIUzI1.NiIXVCJ9TJV",
			token:  "",
			err:    jwt.ErrInvalidToken,
		},
		{
			name:   "test_4",
			header: "Authorization: bearer eyJhbGci.OiJIUzI1NiIXV.CJ9TJV",
			token:  "",
			err:    jwt.ErrInvalidToken,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			token, err := splitToken(tc.header)

			assert.Equal(t, tc.token, token)
			assert.Equal(t, tc.err, err)

		})
	}
}
