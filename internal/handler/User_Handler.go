package handler

import (
	"encoding/json"
	"resto_nm_api/internal/models"
	"resto_nm_api/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/syndtr/goleveldb/leveldb"
)

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]any{
		"message": "hallo",
	})
}

func DeleteUsers(c *fiber.Ctx) error {
	err := repository.DeleteUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal menghapus data users",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Berhasil menghapus data users",
		"data":    []models.Users{},
	})
}


func CreateUsers(c *fiber.Ctx) error{
	var users models.Users

	err := c.BodyParser(&users);
	if err != nil {return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})}

	db, err := repository.RedUsers();

	// jika data nnya tidak ada maka abil data dari body dan gabugkan deggan users
	if err == leveldb.ErrNotFound {
		arrUsers := []models.Users{users}
		// go ke byte
		userBytes, err := json.Marshal(arrUsers)

		if err != nil {return c.Status(500).JSON(fiber.Map{"error": "Failed to encode users"})}

		// simpan data
		err = repository.SaveUsers(userBytes);
		if err != nil {return c.Status(500).JSON(fiber.Map{"error": "Failed to save users"})}

		// jika berhasil kembalikan data tersebut
		return c.Status(201).JSON(arrUsers)


	}else if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to read users from DB"})
	}


	arrUsers := []models.Users{users}
	// byte ke go
	err = json.Unmarshal(db, &arrUsers);

	if err != nil {return c.Status(500).JSON(fiber.Map{"error": "Failed to parse existing users"})}

	// jika data suda ada kembalikan data dari db
	return c.Status(201).JSON(arrUsers)
}