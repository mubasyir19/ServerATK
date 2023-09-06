package main

import (
	"ServerATK/config"
	"ServerATK/database"
	"ServerATK/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config.Config()
	database.ConnectDB()
}

func main() {
	app := fiber.New()

	routes.SetUpRoute(app)

	log.Fatal(app.Listen(":5000"))
}
