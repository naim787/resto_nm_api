package main

import (
	"log"
	"os"
	"path/filepath"
	"resto_nm_api/internal/handler/productHandler"
	"resto_nm_api/internal/handler/usersHandler"
	"resto_nm_api/internal/handler/websocketHandler"
	"resto_nm_api/internal/middleware"
	"resto_nm_api/internal/repository"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	_, err := repository.OpenDB()
	if err != nil {
		log.Fatal("Database initialization failed:", err)
	}

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Logger middleware
	app.Use(logger.New())

	// Role guard middleware
	app.Use(middleware.RoleGuard())

	staticDir := "../frontend"
	app.Static("/static", "../static")
	app.Static("/_app", staticDir+"/_app") // asset JS/CSS
	app.Static("/favicon.svg", staticDir+"/favicon.svg")
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // atau LevelBestCompression
	}))
	// route akses
	// Baca semua file .html di folder frontend
	err = filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			routePath := strings.TrimSuffix(info.Name(), ".html")

			if routePath == "index" {
				// Route untuk "/"
				app.Get("/", func(c *fiber.Ctx) error {
					return c.SendFile(filepath.Join(staticDir, "index.html"))
				})
			} else {
				// Route untuk "/no", "/admin", etc
				route := "/" + routePath
				htmlFile := filepath.Join(staticDir, info.Name())

				// closure fix: capture htmlFile properly
				app.Get(route, func(file string) fiber.Handler {
					return func(c *fiber.Ctx) error {
						return c.SendFile(file)
					}
				}(htmlFile))
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Gagal membaca file HTML: %v", err)
	}

	app.Get("/users", usersHandler.GetUsers)
	app.Get("/menu", productHandler.GetAllProducts)

	// API routes
	app.Post("/create-users", usersHandler.CreateUsers)
	app.Delete("/users/:id", usersHandler.DeleteUsers)

	app.Post("/create-products", productHandler.CreateProducts)
	app.Delete("/products/:id", productHandler.DeleteProduct)
	
	app.Post("/login", usersHandler.LoginUser)

	// WebSocket
	app.Use("/ws", websocketHandler.WebSocketHandler)
	app.Get("/ws/orders", websocket.New(websocketHandler.HandleOrders))

	// Serve static files dari folder frontend
	app.Static("/", "./frontend")
	app.Use(func(c *fiber.Ctx) error {
		// Jika tidak ada route/api yang cocok, dan bukan file statis,
		// fallback ke index.html
		return c.SendFile("./frontend/index.html")
	})
	// log.Fatal(app.Listen(":3000"))
	log.Fatal(app.Listen("0.0.0.0:3001"))

}