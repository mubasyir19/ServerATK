package routes

import (
	"ServerATK/controllers/admin"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoute(router fiber.Router) {
	route := router.Group("/api")

	route.Get("/products", admin.GetAllProducts)
	route.Get("/categories", admin.GetAllCategories)
	route.Post("/categories/add", admin.AddCategory)

}
