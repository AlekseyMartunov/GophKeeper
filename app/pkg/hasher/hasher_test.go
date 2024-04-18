package hasher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasher_Hash(t *testing.T) {
	testCase := []struct {
		name     string
		password string
		salt     string
		text     string
		encoded  string
	}{
		{
			name:     "test_1",
			password: "querty",
			salt:     "12345",
			text:     "some very important info",
			encoded:  "69bcf7a78bfd7d28fc2cb941b5b129a7ff2b8175d353cab08a3a6954d6053372",
		},
		{
			name:     "test_2",
			password: "pass",
			salt:     "ffg",
			text:     "some very important info 2",
			encoded:  "5cca67f8774e78a28ea911dbc231a765ae3a3df26182a91b1a497afa634d7339",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			hasher := NewHasher(tc.salt)

			e := hasher.Hash(tc.text, tc.password)

			assert.Equal(t, tc.encoded, e, "hash does not match expected value")
		})
	}
}
