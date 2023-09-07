package routes

import (
	"ServerATK/controllers/admin"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoute(router fiber.Router) {
	route := router.Group("/api")

	route.Get("/categories", admin.GetAllCategories)
	route.Post("/categories/add", admin.AddCategory)
	route.Put("/categories/edit/:id", admin.EditCategory)
	route.Delete("/categories/delete/:id", admin.DeleteCategory)

	route.Get("/products", admin.GetAllProducts)
	route.Post("/products/add", admin.AddProduct)

}
