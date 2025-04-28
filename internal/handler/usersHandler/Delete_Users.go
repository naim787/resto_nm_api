package usersHandler

import (
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func DeleteUsers(c *fiber.Ctx) error {
	result := repository.DB.Delete(&models.Users{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete users"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "All users deleted successfully",
	})
}