package productHandler

import (
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Products

	if err := repository.DB.Find(&products).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data produk"})
	}

	return c.JSON(fiber.Map{
		"message": "Semua produk berhasil diambil",
		"data":    products,
	})
}