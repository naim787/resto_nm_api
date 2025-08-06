package crypTO

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
)

// ===== 2. Encrypt dengan AES-CTR + HMAC-SHA256 =====
func Encrypt(data, password []byte) ([]byte, error) {
	// Generate IV (nonce)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	// Hash password
	key := sha256.Sum256(password)
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)

	// Enkripsi data
	encrypted := make([]byte, len(data))
	stream.XORKeyStream(encrypted, data)

	// Buat HMAC
	h := hmac.New(sha256.New, key[:])
	h.Write(encrypted)
	hmacSum := h.Sum(nil)

	// Gabungkan IV + data terenkripsi + HMAC
	final := append(iv, encrypted...)
	final = append(final, hmacSum...)
	return final, nil
}
