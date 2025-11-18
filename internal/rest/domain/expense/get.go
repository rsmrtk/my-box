package expense

import (
	"time"

	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

// GetRequest represents the request structure for fetching an expense
type GetRequest struct {
	ExpenseID string `json:"expense_id" binding:"required"`
}

// GetResponse represents the response structure for fetching an expense
type GetResponse struct {
	ExpenseID     string           `json:"expense_id"`
	ExpenseName   string           `json:"expense_name"`
	ExpenseAmount []*models.Amount `json:"expense_amount"`
	ExpenseType   string           `json:"expense_type"`
	ExpenseDate   time.Time        `json:"expense_date"`
	CreatedAt     time.Time        `json:"created_at"`
}
