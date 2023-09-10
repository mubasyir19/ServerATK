package main

import (
	"ServerATK/config"
	"ServerATK/database"
	"ServerATK/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	config.Config()
	database.ConnectDB()
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	routes.SetUpRoute(app)

	log.Fatal(app.Listen(":5000"))
}
