package admin

import (
	"ServerATK/database"
	"ServerATK/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB

	var products []models.Product

	result := db.Find(&products)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed save data to database",
			"error":   result.Error.Error(),
		})
	}

	if len(products) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No products found in the database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success save data to database",
		"data":    result,
	})
}
