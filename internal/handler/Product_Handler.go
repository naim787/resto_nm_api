package handler

import (
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"
	"resto_nm_api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func CreateProducts(c *fiber.Ctx) error {
    var product models.Products
    if err := c.BodyParser(&product); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
    }

    product.ID = service.GenerateUniqueID()
    result := repository.DB.Create(&product)
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create product"})
    }

    return c.Status(201).JSON(fiber.Map{
        "message": "Product created successfully",
        "data":    product,
    })
}