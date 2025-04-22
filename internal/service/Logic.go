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
    // Jika data tidak ada, simpan newItems ke database
    if err == leveldb.ErrNotFound {
        itemBytes, _ := json.Marshal(newItems)
        err = repository.SaveUsers(itemBytes, key) // Simpan data ke database
        if err != nil {
            return nil, c.Status(500).JSON(fiber.Map{"error": "Failed to save data"})
        }
        return newItems, nil
        
    // jika ada kesalahan membaca data dari database
    } else if err != nil {
        return nil, c.Status(500).JSON(fiber.Map{"error": "Failed to read data from DB"})
    }


    // Jika data ditemukan, tambahkan newItems ke data yang ada
    // []T adalah data dari models yang suda di definisikan
    var existingItems []T
    err = json.Unmarshal(dbData, &existingItems)
    if err != nil {
        return nil, c.Status(500).JSON(fiber.Map{"error": "Failed to parse existing data"})
    }

    // Gabungkan data lama dengan data baru
    combinedItems := append(existingItems, newItems...)
    itemBytes, _ := json.Marshal(combinedItems)
    err = repository.SaveUsers(itemBytes, key) // Simpan data yang diperbarui ke database
    if err != nil {
        return nil, c.Status(500).JSON(fiber.Map{"error": "Failed to save updated data"})
    }

    return combinedItems, nil
}




// cari data users yang SAMA deggan data
func FindByID[T any](key string, id string) (*T, error) {
    // Baca data dari database dengan key yang diberikan
    data, err := repository.ReadDB(key)
    if err != nil {
        return nil, err // Jika terjadi kesalahan saat membaca database
    }

    // Unmarshal data JSON ke slice struct T
    var items []T
    err = json.Unmarshal(data, &items)
    if err != nil {
        return nil, err // Jika terjadi kesalahan saat parsing JSON
    }

    // Cari item berdasarkan ID
    for _, item := range items {
        // Gunakan refleksi untuk mendapatkan nilai ID
        itemMap, err := json.Marshal(item)
        if err != nil {
            continue
        }

        var itemData map[string]interface{}
        json.Unmarshal(itemMap, &itemData)

        if itemData["Id"] == id {
            return &item, nil // ID ditemukan, kembalikan data item
        }
    }

    return nil, nil // ID tidak ditemukan
}





// cari data users yang SAMA deggan data
func FindNotByID[T any](key string, id string) (*T, error) {
    // Baca data dari database dengan key yang diberikan
    data, err := repository.ReadDB(key)
    if err != nil {
        return nil, err // Jika terjadi kesalahan saat membaca database
    }

    // Unmarshal data JSON ke slice struct T
    var items []T
    err = json.Unmarshal(data, &items)
    if err != nil {
        return nil, err // Jika terjadi kesalahan saat parsing JSON
    }

    // Cari item berdasarkan ID
    for _, item := range items {
        // Gunakan refleksi untuk mendapatkan nilai ID
        itemMap, err := json.Marshal(item)
        if err != nil {
            continue
        }

        var itemData map[string]interface{}
        json.Unmarshal(itemMap, &itemData)

        if itemData["Id"] != id {
            return &item, nil // ID ditemukan, kembalikan data item
        }
    }

    return nil, nil // ID tidak ditemukan
}