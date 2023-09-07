package admin

import (
	"ServerATK/database"
	"ServerATK/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Add data",
		"result":  category,
	})
}

func EditCategory(c *fiber.Ctx) error {
	db := database.DB

	categoryID := c.Params("id")

	var category models.Category
	find := db.First(&category, "id = ?", categoryID)
	if find.Error != nil {
		if find.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "category not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "A server error occurred",
		})
	}

	name := c.FormValue("name")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Name cannot be empty",
		})
	}

	category.Name = name
	if err := db.Save(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "An error occurred while updating category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success update data",
		"result":  category,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	db := database.DB

	categoryID := c.Params("id")

	var category models.Category

	if err := db.Delete(&category, "id = ?", categoryID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "An error occurred while deleting category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success delete data category",
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

func AddProduct(c *fiber.Ctx) error {
	db := database.DB

	name := c.FormValue("name")
	description := c.FormValue("description")

	priceStr := c.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return err
	}

	stockStr := c.FormValue("stock")
	stock, err := strconv.ParseInt(stockStr, 10, 64)
	if err != nil {
		return err
	}

	// categoryIDStr := c.FormValue("categoryID")
	// if categoryIDStr == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "categoryID cannot be empty",
	// 	})
	// }
	// categoryIDUint, err := strconv.ParseUint(categoryIDStr, 10, 64)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "categoryID must be a valid number",
	// 	})
	// }
	// categoryID := uint(categoryIDUint)

	photo, err := c.FormFile("photo")
	if err != nil {
		return err
	}

	filename := "uploads/" + photo.Filename
	if err := c.SaveFile(photo, filename); err != nil {
		return err
	}

	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		Photo:       filename,
		// CategoryID:  categoryID,
	}

	result := db.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(fiber.Map{
		"message": "Success Add data",
		"result":  product,
	})
}
