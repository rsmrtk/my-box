package income

import (
	"time"

	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

// UpdateRequest represents the request structure for updating an income
type UpdateRequest struct {
	IncomeID     string           `json:"income_id" binding:"required"`
	IncomeName   string           `json:"income_name,omitempty"`
	IncomeAmount []*models.Amount `json:"income_amount,omitempty"`
	IncomeType   string           `json:"income_type,omitempty"`
	IncomeDate   *time.Time       `json:"income_date,omitempty"`
}

// UpdateResponse represents the response structure for updating an income
type UpdateResponse struct {
	IncomeID     string           `json:"income_id"`
	IncomeName   string           `json:"income_name"`
	IncomeAmount []*models.Amount `json:"income_amount"`
	IncomeType   string           `json:"income_type"`
	IncomeDate   time.Time        `json:"income_date"`
	UpdatedAt    time.Time        `json:"updated_at"`
}
