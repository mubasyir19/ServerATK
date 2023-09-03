package controllers

import "github.com/gofiber/fiber/v2"

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
