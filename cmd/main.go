package main

import (
    "log"
    "resto_nm_api/internal/handler"
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
    app.Get("/users", handler.GetUsers)
    app.Post("/create-users", handler.CreateUsers)
    app.Post("/create-products", handler.CreateProducts)

    // WebSocket route for orders
    app.Use("/ws", handler.WebSocketHandler)
    app.Get("/ws/orders", websocket.New(handler.HandleOrders))

    log.Fatal(app.Listen(":3000"))
}