package income

import (
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type GetRequest struct {
	IncomeID string `json:"income_id"`
}

type GetResponse struct {
	IncomeID     string           `json:"income_id"`
	IncomeName   string           `json:"income_name"`
	IncomeAmount []*models.Amount `json:"income_amount"`
	IncomeType   string           `json:"income_type"`
	IncomeDate   models.Date      `json:"income_date"`
	CreatedAt    models.Date      `json:"created_at"`
}
