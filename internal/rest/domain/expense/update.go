package expense

import (
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

// UpdateRequest represents the request structure for updating an expense
type UpdateRequest struct {
	ExpenseID     string           `json:"expense_id" binding:"required"`
	ExpenseName   string           `json:"expense_name,omitempty"`
	ExpenseAmount []*models.Amount `json:"expense_amount,omitempty"`
	ExpenseType   string           `json:"expense_type,omitempty"`
	ExpenseDate   *models.Date     `json:"expense_date,omitempty"`
}

// UpdateResponse represents the response structure for updating an expense
type UpdateResponse struct {
	ExpenseID     string           `json:"expense_id"`
	ExpenseName   string           `json:"expense_name"`
	ExpenseAmount []*models.Amount `json:"expense_amount"`
	ExpenseType   string           `json:"expense_type"`
	ExpenseDate   models.Date      `json:"expense_date"`
	UpdatedAt     models.Date      `json:"updated_at"`
}
