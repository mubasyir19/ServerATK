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

func CreateTable() {
	migrator := database.DB.Migrator()

	migrator.CreateTable(&models.User{})
	migrator.CreateTable(&models.Kategori{})
	migrator.CreateTable(&models.Produk{})
	migrator.CreateTable(&models.Pesanan{})
	migrator.CreateTable(&models.ItemPesanan{})
	migrator.CreateTable(&models.Keranjang{})
	migrator.CreateTable(&models.Pembayaran{})
}

func DropTable() {
	migrator := database.DB.Migrator()

	migrator.DropTable(&models.User{})
	migrator.DropTable(&models.Kategori{})
	migrator.DropTable(&models.Produk{})
	migrator.DropTable(&models.Pesanan{})
	migrator.DropTable(&models.ItemPesanan{})
	migrator.DropTable(&models.Keranjang{})
	migrator.DropTable(&models.Pembayaran{})
}

func main() {
	// Select one of the below, then run => go run migrate/migrate.go
	CreateTable()
	// DropTable()
}
