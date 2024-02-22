package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"io"
)

type config interface {
	SecretKey() string
}

type EncryptionManager struct {
	gcm cipher.AEAD
}

func NewEncryptionManager(cfg config) *EncryptionManager {
	hasher := md5.New()
	hasher.Write([]byte(cfg.SecretKey()))

	block, _ := aes.NewCipher(hasher.Sum(nil))
	gcm, _ := cipher.NewGCM(block)

	return &EncryptionManager{
		gcm: gcm,
	}
}

func (em *EncryptionManager) Encrypt(data string) (string, error) {
	nonce := make([]byte, em.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := em.gcm.Seal(nonce, nonce, []byte(data), nil)
	return string(ciphertext), nil
}

func (em *EncryptionManager) Decrypt(text string) (string, error) {
	data := []byte(text)

	nonceSize := em.gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := em.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
