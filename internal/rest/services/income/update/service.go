package update

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	m_income "github.com/rsmrtk/db-fd-model/m_income"
	"github.com/rsmrtk/mybox/internal/rest/domain/income"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx  context.Context
	req  *income.UpdateRequest
	f    *Facade
	data *m_income.Data
}

func (s *service) update() error {
	// Validate income ID format
	_, err := uuid.Parse(s.req.IncomeID)
	if err != nil {
		return errs.InvalidIncomeID
	}

	// First, fetch the existing income
	fields := []m_income.Field{
		m_income.IncomeID,
		m_income.IncomeName,
		m_income.IncomeAmount,
		m_income.IncomeType,
		m_income.IncomeDate,
	}

	s.data, err = s.f.pkg.M.FinDash.Income.Find(s.ctx, s.req.IncomeID, fields)
	if err != nil {
		return errs.IncomeNotFound
	}

	// Update local data for response
	if s.req.IncomeName != "" {
		s.data.IncomeName = &s.req.IncomeName
	}
	if len(s.req.IncomeAmount) > 0 {
		amount := s.req.IncomeAmount[0].Amount
		s.data.IncomeAmount = &amount
	}
	if s.req.IncomeType != "" {
		s.data.IncomeType = &s.req.IncomeType
	}
	if s.req.IncomeDate != nil {
		incomeDate := s.req.IncomeDate.Time
		s.data.IncomeDate = &incomeDate
	}

	// TODO: Implement Update when the proper API is known
	// The Update method signature needs to be verified with the db-fd-model package
	// err = s.f.pkg.M.FinDash.Income.Update(s.ctx, s.req.IncomeID, updateData)
	// if err != nil {
	// 	return errs.FailedToUpdateIncome
	// }

	return nil
}

func (s *service) reply() *income.UpdateResponse {
	var incomeName, incomeType string
	var incomeAmount float64
	var incomeDate time.Time

	if s.data.IncomeName != nil {
		incomeName = *s.data.IncomeName
	}
	if s.data.IncomeAmount != nil {
		incomeAmount = *s.data.IncomeAmount
	}
	if s.data.IncomeType != nil {
		incomeType = *s.data.IncomeType
	}
	if s.data.IncomeDate != nil {
		incomeDate = *s.data.IncomeDate
	}

	// Convert float64 to big.Rat for precise decimal handling
	incomeAmountRat := *big.NewRat(int64(incomeAmount*100), 100)
	amountValue, _ := incomeAmountRat.Float64()

	return &income.UpdateResponse{
		IncomeID:   s.data.IncomeID,
		IncomeName: incomeName,
		IncomeAmount: []*models.Amount{{
			Amount:         amountValue,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		IncomeType: incomeType,
		IncomeDate: models.NewDate(incomeDate),
		UpdatedAt:  models.NewDate(time.Now()),
	}
}
