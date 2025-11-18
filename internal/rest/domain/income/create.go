package income

import (
	"time"

	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type CreateRequest struct {
	IncomeName   string           `json:"income_name"`
	IncomeAmount []*models.Amount `json:"income_amount"`
	IncomeType   string           `json:"income_type"`
	IncomeDate   time.Time        `json:"income_date"`
}

type CreateResponse struct {
	IncomeID     string           `json:"income_id"`
	IncomeName   string           `json:"income_name"`
	IncomeAmount []*models.Amount `json:"income_amount"`
	IncomeType   string           `json:"income_type"`
	IncomeDate   time.Time        `json:"income_date"`
	CreatedAt    time.Time        `json:"created_at"`
}
