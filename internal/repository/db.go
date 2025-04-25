package repository

import (
    "log"
    "resto_nm_api/internal/models"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Membuka database SQLite dengan GORM
func OpenDB() (*gorm.DB, error) {
    database, err := gorm.Open(sqlite.Open("restaurant.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal konek ke database!")
        return nil, err
    }

    // AutoMigrate untuk membuat tabel secara otomatis
    err = database.AutoMigrate(
        &models.Users{},
        &models.Products{},
        &models.Pesnan{},
    )
    if err != nil {
        log.Fatal("Gagal migrasi tabel!")
        return nil, err
    }

    DB = database
    return DB, nil
}