package productHandler

import (
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func DeleteProduct(c *fiber.Ctx) error {
	// Ambil ID dari URL param
	id := c.Params("id")

	// Cek apakah produk dengan ID tersebut ada
	var product models.Products
	if err := repository.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	// Hapus produk
	if err := repository.DB.Delete(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus produk"})
	}

	return c.JSON(fiber.Map{
		"message": "Produk berhasil dihapus",
	})
}
