package main

import (
	"ServerATK/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetUpRoute(app)

	log.Fatal(app.Listen(":5000"))
}
