package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

// 1.Create data structure and dummy data
// 2. Implementation of the GET API endpoint:
// 3. Implementation of POST API endpoint:

// Create data structure
type progLang struct {
	Id       string
	Language string
	Creator  string
}

// dummy data
var languages = []progLang{
	{Id: "1", Language: "C", Creator: "Dennis Ritchie"},
	{Id: "2", Language: "Java", Creator: "James Gosling"},
	{Id: "3", Language: "C++", Creator: " Bjarne Stroustrup"},
	{Id: "4", Language: "Python", Creator: "Guido van Rossum"},
}

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, Coders! Welcome to Go programming language.")
	})
	// 2. Implementation of the GET API endpoint:
	app.Get("/languages", func(c fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(languages)
	})

	// 3. Implementation of POST API endpoint:
	// app.Post("/languages", func(c fiber.Ctx) error {
	// 	var language progLang

	// 	if err := c.BodyParser(&language); err != nil {
	// 		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	// 		return err
	// 	}
	// 	languages = append(languages, language)
	// 	return c.Status(http.StatusCreated).JSON(language)
	// })

	app.Post("/languages", func(c fiber.Ctx) error {
		var language progLang

		//Parse the request body
		if err := c.Body(); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
			fmt.Println("body ", c.Body())
			//return
		}
		// if err := c.Body(language); err != nil {
		// 	return err
		// }

		// Append the new language data
		languages = append(languages, language)

		return c.Status(http.StatusCreated).JSON(language)
	})

	app.Listen(":8080")
}
