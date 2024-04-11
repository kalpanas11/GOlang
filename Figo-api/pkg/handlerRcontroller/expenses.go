package handlerRcontroller

import (
	"time"

	"github.com/kalpanas11/figo-api/pkg/entityRmodel"

	"github.com/gofiber/fiber/v2"
)

// ExpenseHandler type
type ExpenseHandler struct{}

// Index to list all expenses
func (h ExpenseHandler) Index(ctx fibre.Ctx) error {
	expenseHandler := []entityRmodel.Expense{
		{
			ID:         "1",
			Title:      "Lunch at MyFood",
			Total:      14.95,
			Attachment: "photo.jpg",
			CreatedAt:  time.Now(),
		},
	}
	return ctx.JSON(fiber.Map{"data": expenseHandler})
}
