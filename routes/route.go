package routes

import (
	"ServerATK/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoute(router fiber.Router) {
	route := router.Group("/api")

	route.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	route.Get("/products", controllers.GetAllProducts)
	route.Get("/categories", controllers.GetAllCategories)
	route.Post("/add/category", controllers.AddCategory)
}
