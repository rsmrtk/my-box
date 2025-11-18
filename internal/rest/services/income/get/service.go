package get

import (
	"context"
	"math/big"
	"time"

	"github.com/rsmrtk/db-fd-model/m_income"
	di "github.com/rsmrtk/mybox/internal/rest/domain/income"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx context.Context
	req *di.GetRequest
	f   *Facade

	incomeID     string
	incomeName   string
	incomeAmount big.Rat
	incomeType   string
	incomeDate   time.Time
	createdAt    time.Time
}

func (s *service) find() error {
	// Use the Find method to get a single income by ID
	data, err := s.f.pkg.M.FinDash.Income.Find(s.ctx,
		s.req.IncomeID,
		[]m_income.Field{
			m_income.IncomeID,
			m_income.IncomeName,
			m_income.IncomeAmount,
			m_income.IncomeType,
			m_income.IncomeDate,
			m_income.CreatedAt,
		},
	)
	if err != nil {
		return errs.FailedToFindIncome
	}

	// Extract data from the model
	s.incomeID = data.IncomeID

	if data.IncomeName != nil {
		s.incomeName = *data.IncomeName
	}

	if data.IncomeAmount != nil {
		// Convert float64 to big.Rat
		s.incomeAmount = *big.NewRat(int64(*data.IncomeAmount*100), 100)
	}

	if data.IncomeType != nil {
		s.incomeType = *data.IncomeType
	}

	if data.IncomeDate != nil {
		s.incomeDate = *data.IncomeDate
	}

	if data.CreatedAt != nil {
		s.createdAt = *data.CreatedAt
	}

	return nil
}

func (s *service) reply() *di.GetResponse {
	amountValue, _ := s.incomeAmount.Float64()
	amountObj := &models.Amount{
		Amount:         amountValue,
		CurrencyCode:   "USD",
		CurrencySymbol: "$",
	}

	return &di.GetResponse{
		IncomeID:     s.incomeID,
		IncomeName:   s.incomeName,
		IncomeAmount: []*models.Amount{amountObj},
		IncomeType:   s.incomeType,
		IncomeDate:   s.incomeDate,
		CreatedAt:    s.createdAt,
	}
}
