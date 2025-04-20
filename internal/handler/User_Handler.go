package handler

import (
	"encoding/json"
	"fmt"
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/syndtr/goleveldb/leveldb"
)

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

func CreateUsers(c *fiber.Ctx) error {
    var users models.Users

    // Parse body JSON
    err := c.BodyParser(&users)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
    }

    // ambil data dari users revensi dari models.Users
    // useR := []models.Users{users}

    // // Use reusable function to save or update product
    // products, err := service.SaveOrUpdateData[[]models.Users]("products", useR, c)
    // if err != nil {
    //     return err
    // }

    fmt.Println(users)
    // return c.Status(201).JSON(products)
    return c.Status(200).JSON(fiber.Map{
        "message": "susess",
        "data": []models.Users{users},
    })
}



