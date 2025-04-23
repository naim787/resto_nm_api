package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"resto_nm_api/internal/repository"
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




func FindEmail[T any](key string, emailUs string) (bool, error) {
	// membuka database
    db, err := repository.OpenDB()
    if err != nil {
        log.Fatal("Database initialization failed:", err)
    }
    defer db.Close()

    // Baca data dari database
    data, err := repository.ReadDB(key)
    if err != nil {
        return false, fmt.Errorf("failed to read database: %w", err)
    }

    // Unmarshal data JSON ke slice struct T
    var repData []T
    err = json.Unmarshal(data, &repData)
    if err != nil {
        return false, fmt.Errorf("failed to unmarshal data: %w", err)
    }

    // Periksa apakah email sudah ada
    for _, item := range repData {
        // Gunakan refleksi atau pastikan T adalah map[string]interface{}
        itemMap, err := json.Marshal(item)
        if err != nil {
            continue
        }

        var itemData map[string]interface{}
        json.Unmarshal(itemMap, &itemData)

        if itemData["Email"] == emailUs {
            return true, nil // Email ditemukan
        }
    }

    return false, nil // Email tidak ditemukan
}