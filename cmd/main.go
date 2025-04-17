package main

import (
    "log"
    "resto_nm_api/internal/handler"
    "resto_nm_api/internal/repository"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    // membuka database
    db, err := repository.OpenDB()
    if err != nil {
        log.Fatal("Database initialization failed:", err)
    }
    defer db.Close()

    // middleware
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
    }))

    // routes
    app.Get("/users", handler.GetUsers)
    app.Get("/delete-users", handler.DeleteUsers)
    app.Post("/create-users", handler.CreateUsers)

    log.Fatal(app.Listen(":3000"))
}