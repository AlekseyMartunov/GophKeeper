package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"strings"
)

type Hasher struct {
	sha  hash.Hash
	salt string
}

func NewHasher(salt string) *Hasher {
	s := sha256.New()

	return &Hasher{
		salt: salt,
		sha:  s,
	}
}

func (h *Hasher) Hash(text, key string) string {
	b := []byte(strings.Join([]string{text, key, h.salt}, ""))
	h.sha.Write(b)
	res := h.sha.Sum(nil)

	h.sha.Reset()

	return hex.EncodeToString(res)
}
