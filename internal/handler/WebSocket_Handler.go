package handler

import (
    "encoding/json"
    "fmt"
    "resto_nm_api/internal/models"
    "resto_nm_api/internal/repository"

    "github.com/gofiber/contrib/websocket"
    "github.com/gofiber/fiber/v2"
)

// Middleware to check WebSocket upgrade
func WebSocketHandler(c *fiber.Ctx) error {
    // Check if the request is a WebSocket upgrade
    if websocket.IsWebSocketUpgrade(c) {
        c.Locals("allowed", true)
        return c.Next()
    }
    return fiber.ErrUpgradeRequired
}

// WebSocket connection handler for orders
func HandleOrders(c *websocket.Conn) {
    defer c.Close()

    for {
        // Read message from the WebSocket
        var (
            msg []byte
            err error
        )
        if _, msg, err = c.ReadMessage(); err != nil {
            fmt.Println("Error reading message:", err)
            break
        }

        // Parse the received JSON into an array of orders
        var orders []models.Pesnan
        err = json.Unmarshal(msg, &orders)
        if err != nil {
            fmt.Println("Invalid JSON format:", err)
            continue
        }

        // Save the orders to the database with the key "pesanan"
        orderBytes, _ := json.Marshal(orders)
        err = repository.SaveUsers(orderBytes, "pesanan")
        if err != nil {
            fmt.Println("Error saving orders to database:", err)
            continue
        }

        // Send acknowledgment back to the client
        response := map[string]string{
            "message": "Orders saved successfully",
        }
        if err = c.WriteJSON(response); err != nil {
            fmt.Println("Error sending acknowledgment:", err)
            break
        }

        fmt.Println("Orders saved successfully:", orders)
    }
}