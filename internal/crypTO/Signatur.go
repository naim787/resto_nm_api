package crypTO

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GeneratePassword(seed []byte) string {
	_, privKey, _ := ed25519.GenerateKey(rand.Reader)
	hash := sha256.Sum256(seed)
	signature := ed25519.Sign(privKey, hash[:])
	final := sha256.Sum256(signature)
	return hex.EncodeToString(final[:])
}