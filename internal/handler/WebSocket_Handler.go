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
        var msg []byte
        _, msg, err := c.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            break
        }

        var orders []models.Pesnan
        if err := json.Unmarshal(msg, &orders); err != nil {
            fmt.Println("Invalid JSON format:", err)
            continue
        }

        for _, order := range orders {
            orderJSON, _ := json.Marshal(order.Products)
            order.Products = string(orderJSON)
            if err := repository.DB.Create(&order).Error; err != nil {
                fmt.Println("Error saving order:", err)
                continue
            }
        }

        response := map[string]string{"message": "Orders saved successfully"}
        if err := c.WriteJSON(response); err != nil {
            fmt.Println("Error sending acknowledgment:", err)
            break
        }
    }
}