package usersHandler

import (
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"

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