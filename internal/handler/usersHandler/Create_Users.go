package usersHandler

import (
	"resto_nm_api/internal/crypTO"
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"
	"resto_nm_api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func CreateUsers(c *fiber.Ctx) error {
	var user models.Users
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user.ID = service.GenerateUniqueID()
	user.Role = "user"

	encryptedPassword, err := crypTO.Encrypt([]byte(user.Password), []byte(user.Email))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to encrypt password"})
	}
    user.Password = string(encryptedPassword)


	result := repository.DB.Create(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user,
	})
}