package handler

import (
	"resto_nm_api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProducts(c *fiber.Ctx) error {
	var product models.Products

    // Parse body JSON
    err := c.BodyParser(&product)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
    }

    // productT := []models.Products{product}
    // // Use reusable function to save or update product
    // products, err := service.SaveOrUpdateData[[]models.Products]("products", productT, c)
    // if err != nil {
    //     return err
    // }

    // return c.Status(201).JSON(products)
    return c.Status(200).JSON(fiber.Map{
        "message": "susess",
        "data": []models.Products{},
    })
}