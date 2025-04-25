package service

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUniqueID() string {
	// Seed untuk memastikan angka acak berbeda setiap kali dijalankan
	rand.Seed(time.Now().UnixNano())

	// Hasilkan angka acak dengan panjang 4 digit
	id := rand.Intn(10000) // Angka acak antara 0 dan 9999

	// Format angka menjadi string dengan panjang 4 karakter (padding dengan nol jika perlu)
	return fmt.Sprintf("%04d", id)
}