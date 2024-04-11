package route

import (
	"figo-api/pkg/handlerRcontroller"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/figo-api/pkg/handlerRcontroller"
)

// Expenses route
func Expenses(app *fiber.App) {
	//var h handler.ExpenseHandler
	r := app.Group("/expenses")
	r.Get("/", handlerRcontroller.ExpenseHandler.Index)
}
