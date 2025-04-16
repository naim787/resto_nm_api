package main

import (
	"log"
	"resto_nm_api/internal/handler"
	"resto_nm_api/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New();

	// membuka database
	db := repository.OpenDB();
	// if db != nil {log.Fatal("Database initialization failed.", db)};
	defer db.Close();



	app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
    }));

	app.Get("/users", handler.GetUsers);
	app.Get("/delete-users", handler.DeleteUsers)

	app.Post("/create-users", handler.CreateUsers);

	log.Fatal(app.Listen(":3000"));
}