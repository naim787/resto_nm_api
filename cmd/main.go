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
    // CARI SEMUA DATA
    app.Get("/users", handler.GetUsers)

    // CARI ID USERS
    app.Get("/user", handler.GetUserByID)

    // HAPUS SEMUA DATA USERS
    app.Get("/delete-all-users", handler.DeleteUsers)

    // HAPUS DATA USERS BERDASARKAN ID
    app.Get("/delete-users", handler.DeleteUsersById)
     
    // BUAT DATA USERS BARU
    app.Post("/create-users", handler.CreateUsers)

    // BUAT DATA PRODUCT BARU
    app.Post("/create-products", handler.CreateProducts)


     // Middleware for WebSocket upgrade
     app.Use("/ws", handler.WebSocketHandler)

     // WebSocket route for orders
     app.Get("/ws/orders", websocket.New(handler.HandleOrders))

    // port := os.Getenv("PORT")
	// log.Fatal(app.Listen(":" + port))
    log.Fatal(app.Listen(":3000"))
}