package usersHandler

import (
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func DeleteUsers(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.Users
	if err := repository.DB.First(&user, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	if err := repository.DB.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus user"})
	}

	return c.JSON(fiber.Map{
		"message": "User berhasil dihapus",
	})
}