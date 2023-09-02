package routes

import "github.com/gofiber/fiber/v2"

func SetUpRoute(router fiber.Router) {
	route := router.Group("/api")

	route.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello World",
		})
	})
}
