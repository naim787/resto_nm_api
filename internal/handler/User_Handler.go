package handler

import (
    "encoding/json"
    "resto_nm_api/internal/models"
    "resto_nm_api/internal/repository"

    "github.com/gofiber/fiber/v2"
    "github.com/syndtr/goleveldb/leveldb"
)

func GetUsers(c *fiber.Ctx) error {
    data, err := repository.RedUsers()
    if err == leveldb.ErrNotFound {
        return c.Status(200).JSON(fiber.Map{
            "message": "No users found",
            "data":    []models.Users{},
        })
    } else if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to read users from DB"})
    }

    var users []models.Users
    err = json.Unmarshal(data, &users)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to parse users data"})
    }

    return c.Status(200).JSON(fiber.Map{
        "message": "Users retrieved successfully",
        "data":    users,
    })
}

func DeleteUsers(c *fiber.Ctx) error {
    err := repository.DeleteUsers()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to delete users",
        })
    }

    return c.Status(200).JSON(fiber.Map{
        "message": "Users deleted successfully",
        "data":    []models.Users{},
    })
}

func CreateUsers(c *fiber.Ctx) error {
    var user models.Users

    // Parse body JSON
    err := c.BodyParser(&user)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
    }

    // Baca data dari database
    dbData, err := repository.RedUsers()
    if err == leveldb.ErrNotFound {
        // Jika data tidak ditemukan, buat array baru
        users := []models.Users{user}
        userBytes, _ := json.Marshal(users)
        err = repository.SaveUsers(userBytes)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Failed to save users"})
        }
        return c.Status(201).JSON(users)
    } else if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to read users from DB"})
    }

    // Jika data ditemukan, tambahkan user baru
    var users []models.Users
    err = json.Unmarshal(dbData, &users)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to parse existing users"})
    }

    users = append(users, user)
    userBytes, _ := json.Marshal(users)
    err = repository.SaveUsers(userBytes)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to save users"})
    }

    return c.Status(201).JSON(users)
}