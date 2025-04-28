package main

import (
	"log"
	"resto_nm_api/internal/handler/productHandler"
	"resto_nm_api/internal/handler/usersHandler"
	"resto_nm_api/internal/handler/websocketHandler"
	"resto_nm_api/internal/repository"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    // Membuka database
    _, err := repository.OpenDB()
    if err != nil {
        log.Fatal("Database initialization failed:", err)
    }

    // Middleware
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
    }))

    // Routes
    app.Get("/users", usersHandler.GetUsers)
    app.Post("/create-users", usersHandler.CreateUsers)
    app.Post("/create-products", productHandler.CreateProducts)

    // WebSocket route for orders
    app.Use("/ws", websocketHandler.WebSocketHandler)
    app.Get("/ws/orders", websocket.New(websocketHandler.HandleOrders))

    log.Fatal(app.Listen(":3000"))
}