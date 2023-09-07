package admin

import (
	"ServerATK/database"
	"ServerATK/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllCategories(c *fiber.Ctx) error {
	db := database.DB

	var categories []models.Category

	result := db.Find(&categories)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed save data to database",
			"error":   result.Error.Error(),
		})
	}

	if len(categories) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No categories found in the database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success save data to database",
		"result":  categories,
	})

}

func AddCategory(c *fiber.Ctx) error {
	db := database.DB

	name := c.FormValue("name")

	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Name cannot be empty",
		})
	}

	var category = models.Category{
		ID:   uuid.New(),
		Name: name,
	}

	result := db.Create(&category)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(fiber.Map{
		"message": "Success Add data",
		"result":  category,
	})
}

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
		"result":  products,
	})
}
