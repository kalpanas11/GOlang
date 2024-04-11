package route

import (
	"github.com/kalpanas11/golang/figo-api/pkg/handlerRcontroller"
)

// Expenses route
func Expenses(app *fiber.App) {
	//var h handler.ExpenseHandler
	r := app.Group("/expenses")
	r.Get("/", handlerRcontroller.ExpenseHandler.Index)
}
