package handler

import (
	"encoding/json"
	"fmt"
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"
	"resto_nm_api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/syndtr/goleveldb/leveldb"
)

////////////// CARI USERS //////////////////
func GetUsers(c *fiber.Ctx) error {
    data, err := repository.ReadDB("users")
    if err == leveldb.ErrNotFound {
        // jika data nnya kosog
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



////////////// BUAT USERS BARU //////////////////
func CreateUsers(c *fiber.Ctx) error {
    var users models.Users

    // Parse body JSON
    err := c.BodyParser(&users)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
    }

    users.Role = "users"
    users.Id = service.GenerateUniqueID()

    // Gunakan key "users" untuk menyimpan data pengguna
    savedUsers, err := service.SaveOrUpdateData("users", []models.Users{users}, c)
    if err != nil {
        return err
    }

    fmt.Println(savedUsers)
    return c.Status(200).JSON(fiber.Map{
        "message": "susess",
        "data": savedUsers,
    })
}




////////////// CARI ID USERS //////////////////
func GetUserByID(c *fiber.Ctx) error {
    // Tangkap query parameter "id"
    id := c.Query("id")

    // Periksa apakah "id" ada
    if id == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "ID is required",
        })
    }

    // Gunakan ID untuk mencari data
    user, err := service.FindByID[models.Users]("users", id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to retrieve user",
        })
    }

    if user == nil {
        return c.Status(404).JSON(fiber.Map{
            "message": "User not found",
        })
    }

    // Kembalikan data pengguna
    return c.Status(200).JSON(fiber.Map{
        "message": "User found",
        "data":    user,
    })
}


func DeleteUsersById(c *fiber.Ctx) error {
    // Tangkap query parameter "id"
    id := c.Query("id")

    // Periksa apakah "id" ada
    if id == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "ID is required",
        })
    }

    // Gunakan FindNotByID untuk mendapatkan data yang tidak memiliki ID tertentu
    filteredUsers, err := service.FindNotByID[models.Users]("users", id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to retrieve users",
        })
    }

    // Jika tidak ada data yang tersisa, kembalikan respons
    if len(filteredUsers) == 0 {
        return c.Status(404).JSON(fiber.Map{
            "message": "No users found after deletion",
        })
    }

    // Simpan data yang diperbarui ke database
    updatedData, err := json.Marshal(filteredUsers)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to marshal updated users data",
        })
    }

    err = repository.SaveUsers(updatedData, "users")
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to save updated users data",
        })
    }

    // Kembalikan respons sukses
    return c.Status(200).JSON(fiber.Map{
        "message": "User deleted successfully",
        "remaining_users": filteredUsers,
    })
}