package expense

import (
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

// ListRequest represents the request structure for listing expenses
type ListRequest struct {
	Limit  int    `json:"limit,omitempty"`   // Optional: limit number of results
	Offset int    `json:"offset,omitempty"`  // Optional: offset for pagination
	SortBy string `json:"sort_by,omitempty"` // Optional: field to sort by
	Order  string `json:"order,omitempty"`   // Optional: asc or desc
}

// ListItem represents a single expense item in the list
type ListItem struct {
	ExpenseID     string           `json:"expense_id"`
	ExpenseName   string           `json:"expense_name"`
	ExpenseAmount []*models.Amount `json:"expense_amount"`
	ExpenseType   string           `json:"expense_type"`
	ExpenseDate   models.Date      `json:"expense_date"`
	CreatedAt     models.Date      `json:"created_at"`
}

// ListResponse represents the response structure for listing expenses
type ListResponse struct {
	Items      []*ListItem `json:"items"`
	TotalCount int         `json:"total_count"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
}
