package expense

import (
	"time"

	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

// CreateRequest represents the request structure for creating an expense
type CreateRequest struct {
	ExpenseName   string           `json:"expense_name" binding:"required"`
	ExpenseAmount []*models.Amount `json:"expense_amount" binding:"required"`
	ExpenseType   string           `json:"expense_type" binding:"required"`
	ExpenseDate   time.Time        `json:"expense_date" binding:"required"`
}

// CreateResponse represents the response structure for creating an expense
type CreateResponse struct {
	ExpenseID     string           `json:"expense_id"`
	ExpenseName   string           `json:"expense_name"`
	ExpenseAmount []*models.Amount `json:"expense_amount"`
	ExpenseType   string           `json:"expense_type"`
	ExpenseDate   time.Time        `json:"expense_date"`
	CreatedAt     time.Time        `json:"created_at"`
}
