package productHandler

import (
	"fmt"
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"
	"resto_nm_api/internal/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateProducts(c *fiber.Ctx) error {
	// Ambil field dari form
	name := c.FormValue("name")
	description := c.FormValue("description")
	category := c.FormValue("category")
	price := c.FormValue("price")
	stock := c.FormValue("stock")

	// Convert price & stock
	priceFloat, _ := strconv.ParseFloat(price, 64)
	stockInt, _ := strconv.Atoi(stock)

	// Handle file upload
	fileHeader, err := c.FormFile("image_url")
	var imageURL string

	if err == nil && fileHeader != nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
		path := fmt.Sprintf("../static/%s", filename) // <- FIXED

		fmt.Println("Menyimpan ke:", path)

		if err := c.SaveFile(fileHeader, path); err != nil {
			fmt.Println("Gagal simpan file:", err)
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan gambar"})
		}

		imageURL = "/static/" + filename
	} else {
		fmt.Println("File tidak dikirim atau error:", err)
	}


	// Generate ID & simpan ke DB
	product := models.Products{
		ID:          service.GenerateUniqueID(),
		Name:        name,
		Description: description,
		Category:    category,
		Price:       priceFloat,
		Stock:       stockInt,
		ImageURL:    imageURL,
	}

	result := repository.DB.Create(&product)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create product"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Product created successfully",
		"data":    product,
	})
}

// func CreateProducts(c *fiber.Ctx) error {
//     var product models.Products
//     if err := c.BodyParser(&product); err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
//     }

//     product.ID = service.GenerateUniqueID()
//     result := repository.DB.Create(&product)
//     if result.Error != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Failed to create product"})
//     }

//     return c.Status(201).JSON(fiber.Map{
//         "message": "Product created successfully",
//         "data":    product,
//     })
// }
