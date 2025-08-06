package usersHandler

import (
	"resto_nm_api/internal/crypTO"
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {
	type LoginForm struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var form LoginForm
	if err := c.BodyParser(&form); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.Users
	if err := repository.DB.Where("email = ?", form.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Email tidak ditemukan"})
	}

	decryptedPassword, err := crypTO.Decrypt([]byte(user.Password), []byte(form.Email))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Gagal mendekripsi password"})
	}

	if string(decryptedPassword) != form.Password {
		return c.Status(401).JSON(fiber.Map{"error": "Password salah"})
	}

	token := crypTO.GeneratePassword([]byte(user.Email))

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    string(token),
		HTTPOnly: false,
		Secure:   false,
		Path:     "/",
		MaxAge:   300, // 5 menit = 300 detik
		Expires:  time.Now().Add(5 * time.Minute),
	})

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   string(token),
	})
}
