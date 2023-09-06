package controllers

import (
	"ServerATK/database"
	"ServerATK/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
			"message": "No products found in the database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success save data to database",
		"data":    categories,
	})
}

func AddCategory(c *fiber.Ctx) error {
	db := database.DB

	name := c.FormValue("name")

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
		"result":  result,
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
	categoryIDUint, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		return err
	}
	categoryID := uint(categoryIDUint)

	photo, err := c.FormFile("photo")
	if err != nil {
		return err
	}

	filename := "uploads/" + photo.Filename
	if err := c.SaveFile(photo, filename); err != nil {
		return err
	}
	// upload := c.SaveFile(photo, filename)
	// if upload.Error != nil {
	// 	return c.JSON(fiber.Map{
	// 		"message": "Failed to upload file",
	// 		"result":  upload.Error(),
	// 	})
	// }

	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		CategoryID:  categoryID,
	}

	result := db.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(fiber.Map{
		"message": "Success Add data",
		"result":  result,
	})

}
