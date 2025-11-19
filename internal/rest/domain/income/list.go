package income

import (
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

// ListRequest represents the request structure for listing incomes
type ListRequest struct {
	Limit  int    `json:"limit,omitempty"`   // Optional: limit number of results
	Offset int    `json:"offset,omitempty"`  // Optional: offset for pagination
	SortBy string `json:"sort_by,omitempty"` // Optional: field to sort by
	Order  string `json:"order,omitempty"`   // Optional: asc or desc
}

// ListItem represents a single income item in the list
type ListItem struct {
	IncomeID     string           `json:"income_id"`
	IncomeName   string           `json:"income_name"`
	IncomeAmount []*models.Amount `json:"income_amount"`
	IncomeType   string           `json:"income_type"`
	IncomeDate   models.Date      `json:"income_date"`
	CreatedAt    models.Date      `json:"created_at"`
}

// ListResponse represents the response structure for listing incomes
type ListResponse struct {
	Items      []*ListItem `json:"items"`
	TotalCount int         `json:"total_count"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
}
