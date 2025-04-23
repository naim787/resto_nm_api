package handler

import (
	"fmt"
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func CreateProducts(c *fiber.Ctx) error {
    var product models.Products

    // Parse body JSON
    err := c.BodyParser(&product)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
    }

    product.ID = service.GenerateUniqueID()

    //ambil data dari product revensi dari models.Products
    products, err := service.SaveOrUpdateData("products", []models.Products{product}, c)
    if err != nil {
        return err
    }

    fmt.Println(products)
    // return c.Status(201).JSON(products)
    return c.Status(200).JSON(fiber.Map{
        "message": "susess",
        "data": products,
    })
}