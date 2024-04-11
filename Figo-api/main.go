package main

import (
	"github.com/kalpanas11/figo-api/pkg/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		//return c.SendString("Hello, Coders! Welcome to Go programming language.")
		return c.JSON(map[string]string{"message": "Figo API"})
	})
	route.Expenses(app) //run by going to http://localhost:3000/expenses
	app.Listen(":8080")
}
