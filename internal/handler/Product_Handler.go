package handler

import (
	"log"
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func CreateProducts(c *fiber.Ctx) error {
	var products  models.Products

	// memparsig body dari format json ke products
	err := c.BodyParser(&products)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
    }


	data, err := repository.ReadDB("products")
	if err == leveldb.ErrNotFound {
        return c.Status(200).JSON(fiber.Map{
            "message": "No users found",
            "data":    []models.Users{},
        })
    } else if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to read users from DB"})
    }


	return c.SendString("berhasil menambhkan products")
}