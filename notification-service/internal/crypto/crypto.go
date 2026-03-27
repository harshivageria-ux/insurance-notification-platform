package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// Encryptor handles encryption and decryption
type Encryptor struct {
	key []byte
}

// NewEncryptor creates a new encryptor instance
func NewEncryptor() *Encryptor {
	key := os.Getenv("ENCRYPTION_KEY")
	if key == "" {
		key = "32-character-encryption-key!!!!!" // Default for development, CHANGE IN PRODUCTION
	}

	keyBytes := make([]byte, 32)
	copy(keyBytes, key)

	return &Encryptor{
		key: keyBytes,
	}
}

// Encrypt encrypts plaintext using AES
func (e *Encryptor) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using AES
func (e *Encryptor) Decrypt(ciphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext2 := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext2, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
