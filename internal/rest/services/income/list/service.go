package list

import (
	"context"
	"math/big"
	"time"

	"github.com/rsmrtk/db-fd-model/m_income"
	"github.com/rsmrtk/mybox/internal/rest/domain/income"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx   context.Context
	req   *income.ListRequest
	f     *Facade
	items []*m_income.Data
	total int
}

func (s *service) list() error {
	// Set default values if not provided
	if s.req.Limit <= 0 {
		s.req.Limit = 100 // Default limit
	}
	if s.req.Offset < 0 {
		s.req.Offset = 0
	}

	// Build query parameters
	queryParams := []m_income.QueryParam{}

	// Add pagination parameters if needed
	// Note: The actual implementation depends on the QueryParam type
	// This is a simplified version

	// Fetch all incomes
	var err error
	s.items, err = s.f.pkg.M.FinDash.Income.List(s.ctx, queryParams)
	if err != nil {
		return errs.FailedToListIncomes
	}

	// Calculate total count
	s.total = len(s.items)

	// Apply manual pagination to the results
	if s.req.Offset < len(s.items) {
		end := s.req.Offset + s.req.Limit
		if end > len(s.items) {
			end = len(s.items)
		}
		s.items = s.items[s.req.Offset:end]
	} else {
		s.items = []*m_income.Data{}
	}

	return nil
}

func (s *service) reply() *income.ListResponse {
	items := make([]*income.ListItem, 0, len(s.items))

	for _, data := range s.items {
		var incomeName, incomeType string
		var incomeAmount float64
		var incomeDate, createdAt time.Time

		// Extract nullable fields
		if data.IncomeName != nil {
			incomeName = *data.IncomeName
		}
		if data.IncomeAmount != nil {
			incomeAmount = *data.IncomeAmount
		}
		if data.IncomeType != nil {
			incomeType = *data.IncomeType
		}
		if data.IncomeDate != nil {
			incomeDate = *data.IncomeDate
		}
		if data.CreatedAt != nil {
			createdAt = *data.CreatedAt
		}

		// Convert float64 to big.Rat for precise decimal handling
		incomeAmountRat := *big.NewRat(int64(incomeAmount*100), 100)
		amountValue, _ := incomeAmountRat.Float64()

		item := &income.ListItem{
			IncomeID:   data.IncomeID,
			IncomeName: incomeName,
			IncomeAmount: []*models.Amount{{
				Amount:         amountValue,
				CurrencyCode:   "USD",
				CurrencySymbol: "$",
			}},
			IncomeType: incomeType,
			IncomeDate: models.NewDate(incomeDate),
			CreatedAt:  models.NewDate(createdAt),
		}

		items = append(items, item)
	}

	return &income.ListResponse{
		Items:      items,
		TotalCount: s.total,
		Limit:      s.req.Limit,
		Offset:     s.req.Offset,
	}
}
