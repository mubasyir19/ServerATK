package controllers

import (
	"ServerATK/database"
	"ServerATK/models"

	"github.com/gofiber/fiber/v2"
)

func SignUpAdmin(c *fiber.Ctx) {
	// request body

	// Hashing password

	// kirim ke dbase

	// respon
}

func SignInAdmin(c *fiber.Ctx) {
	// request body (username & password)

	// Cek data user (dari body username)

	// Compare hashPass dengan pass

	// Generate jwt token

	// respon token
}

func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []models.Product

	result := db.Find(&products)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed save data to database",
			"error":  result.Error.Error(),
		})
	}

	if len(products) == 0 {
        return c.Status(404).JSON(fiber.Map{
            "message": "No products found in the database",
        })
    }

	return c.Status(200).JSON(fiber.Map{
		"message": "Success save data to database",
		"data": result,
	})
}