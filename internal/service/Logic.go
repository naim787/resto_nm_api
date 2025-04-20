package service

import (
    "encoding/json"
    "resto_nm_api/internal/repository"

    "github.com/gofiber/fiber/v2"
    "github.com/syndtr/goleveldb/leveldb"
)

// SaveOrUpdateData is a reusable function to handle database operations
func SaveOrUpdateData[T any](key string, newItems []T, c *fiber.Ctx) ([]T, error) {
    // Baca data dari database
    dbData, err := repository.ReadDB(key)
    if err == leveldb.ErrNotFound {
        // Jika data tidak ditemukan, simpan newItems langsung
        itemBytes, _ := json.Marshal(newItems)
        err = repository.SaveUsers(itemBytes) // Simpan data ke database
        if err != nil {
            return nil, c.Status(500).JSON(fiber.Map{"error": "Failed to save data"})
        }
        return newItems, nil
    } else if err != nil {
        return nil, c.Status(500).JSON(fiber.Map{"error": "Failed to read data from DB"})
    }

    // Jika data ditemukan, tambahkan newItems ke data yang ada
    var existingItems []T
    err = json.Unmarshal(dbData, &existingItems)
    if err != nil {
        return nil, c.Status(500).JSON(fiber.Map{"error": "Failed to parse existing data"})
    }

    // Gabungkan data lama dengan data baru
    combinedItems := append(existingItems, newItems...)
    itemBytes, _ := json.Marshal(combinedItems)
    err = repository.SaveUsers(itemBytes) // Simpan data yang diperbarui ke database
    if err != nil {
        return nil, c.Status(500).JSON(fiber.Map{"error": "Failed to save updated data"})
    }

    return combinedItems, nil
}