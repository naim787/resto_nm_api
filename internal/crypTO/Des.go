package crypTO

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// ===== 3. Decrypt dan verifikasi HMAC =====
func Decrypt(ciphertext, password []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize+32 {
		return nil, fmt.Errorf("data tidak valid atau terlalu pendek")
	}

	// Ambil bagian-bagian
	iv := ciphertext[:aes.BlockSize]
	data := ciphertext[aes.BlockSize : len(ciphertext)-32]
	hmacStored := ciphertext[len(ciphertext)-32:]

	// Hash password
	key := sha256.Sum256(password)

	// Verifikasi HMAC
	h := hmac.New(sha256.New, key[:])
	h.Write(data)
	expectedHMAC := h.Sum(nil)
	if !hmac.Equal(hmacStored, expectedHMAC) {
		return nil, fmt.Errorf("HMAC tidak cocok! Data rusak atau dimodifikasi")
	}

	// Dekripsi
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, iv)
	decrypted := make([]byte, len(data))
	stream.XORKeyStream(decrypted, data)

	return decrypted, nil
}
