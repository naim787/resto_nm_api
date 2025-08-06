package middleware

import (
	"resto_nm_api/internal/crypTO"
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RoleGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()

		// 1. Publik: bisa akses tanpa login
		public := []string{"/", "/login", "/create-users", "/verify-token", "/favicon.ico"}
		for _, p := range public {
			if strings.HasPrefix(path, p) {
				return c.Next()
			}
		}

		// 2. Ambil dan validasi token
		token := c.Cookies("token")
		if token == "" {
			return fiber.ErrNotFound
		}

		var users []models.Users
		repository.DB.Find(&users)

		var me *models.Users
		for _, u := range users {
			if string(crypTO.GeneratePassword([]byte(u.Email))) == token {
				me = &u
				break
			}
		}
		if me == nil {
			return fiber.ErrNotFound
		}

		// 3. Proteksi rute /no → hanya "user"
		if strings.HasPrefix(path, "/no") && me.Role != "user" {
			return fiber.ErrNotFound
		}

		// 4. Proteksi rute /admin → hanya "admin"
		if strings.HasPrefix(path, "/admin") && me.Role != "admin" {
			return fiber.ErrNotFound
		}

		// 5. Proteksi rute /helper → bisa user & admin
		if strings.HasPrefix(path, "/helper") {
			// boleh lanjut
			return c.Next()
		}

		// 6. Semua rute jika termasuk selain di atas dianggap tidak ada
		return fiber.ErrNotFound
	}
}
