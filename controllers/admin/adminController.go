package admin

import (
	"ServerATK/database"
	"ServerATK/models"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllCategories(c *fiber.Ctx) error {
	db := database.DB

	var categories []models.Category

	result := db.Preload("Products").Find(&categories)
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

	result := db.Preload("Category").Find(&products)
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

	categoryIDStr := c.FormValue("categoryID")
	if categoryIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "categoryID cannot be empty",
		})
	}

	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "categoryID must be a valid UUID",
		})
	}

	photo, err := c.FormFile("photo")
	if err != nil {
		return err
	}

	filename := "uploads/" + photo.Filename
	if err := c.SaveFile(photo, filename); err != nil {
		return err
	}

	product := models.Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		Photo:       filename,
		CategoryID:  categoryID,
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

func EditProduct(c *fiber.Ctx) error {
	db := database.DB

	productID := c.Params("id")

	var product models.Product

	if err := db.First(&product, "id = ? ", productID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "A server error occurred",
		})
	}

	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	stockStr := c.FormValue("stock")
	categoryIDStr := c.FormValue("categoryID")

	if name != "" {
		product.Name = name
	}

	if description != "" {
		product.Description = description
	}

	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid price format",
			})
		}
		product.Price = price
	}

	if stockStr != "" {
		stock, err := strconv.ParseInt(stockStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid stock format",
			})
		}
		product.Price = float64(stock)
	}

	if categoryIDStr != "" {
		categoryID, err := uuid.Parse(categoryIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid category ID",
			})
		}
		product.CategoryID = categoryID
	}

	photo, err := c.FormFile("photo")
	if err == nil {
		filename := "uploads/" + photo.Filename
		if err := c.SaveFile(photo, filename); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to save photo",
			})
		}
		product.Photo = filename
	}

	if err := db.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"result":  product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	db := database.DB

	productID := c.Params("id")

	var product models.Product

	if product.Photo != "" {
		if err := os.Remove(product.Photo); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to delete product photo",
			})
		}
	}

	if err := db.Delete(&product, "id = ?", productID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "An error occurred while deleting category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success delete data product",
	})

}
