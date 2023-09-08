package main

import (
	"ServerATK/config"
	"ServerATK/database"
	"ServerATK/models"
)

func init() {
	config.Config()
	database.ConnectDB()
}

func Migration() {
	database.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Product{},
		&models.OrderDetail{},
		&models.Cart{},
		&models.Payment{},
	)
}

func CreateTable() {
	migrator := database.DB.Migrator()

	migrator.CreateTable(&models.User{})
	migrator.CreateTable(&models.Category{})
	migrator.CreateTable(&models.Product{})
	migrator.CreateTable(&models.Product{})
	migrator.CreateTable(&models.OrderDetail{})
	migrator.CreateTable(&models.Cart{})
	migrator.CreateTable(&models.Payment{})
}

func DropTable() {
	migrator := database.DB.Migrator()

	migrator.DropTable(&models.User{})
	migrator.DropTable(&models.Category{})
	migrator.DropTable(&models.Product{})
	migrator.DropTable(&models.Order{})
	migrator.DropTable(&models.OrderDetail{})
	migrator.DropTable(&models.Cart{})
	migrator.DropTable(&models.Payment{})
}

func main() {
	// Select one of the below, then run => go run migrate/migrate.go
	Migration()
	// CreateTable()
	// DropTable()
}
