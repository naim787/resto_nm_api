package handler

import (
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"
	"resto_nm_api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
    var users []models.Users
    result := repository.DB.Find(&users)
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve users"})
    }

    return c.Status(200).JSON(fiber.Map{
        "message": "Users retrieved successfully",
        "data":    users,
    })
}

func CreateUsers(c *fiber.Ctx) error {
    var user models.Users
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
    }

    user.ID = service.GenerateUniqueID()
    user.Role = "user"
    
    result := repository.DB.Create(&user)
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
    }

    return c.Status(201).JSON(fiber.Map{
        "message": "User created successfully",
        "data":    user,
    })
}

func DeleteUsers(c *fiber.Ctx) error {
    result := repository.DB.Delete(&models.Users{})
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to delete users"})
    }

    return c.Status(200).JSON(fiber.Map{
        "message": "All users deleted successfully",
    })
}