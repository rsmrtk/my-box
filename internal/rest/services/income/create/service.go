package create

import (
	"context"
	"math/big"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
	"github.com/rsmrtk/db-fd-model/m_income"
	di "github.com/rsmrtk/mybox/internal/rest/domain/income"
	amount "github.com/rsmrtk/mybox/internal/rest/domain/models"
)

//create
//reply

type service struct {
	ctx context.Context
	req *di.CreateRequest
	f   *Facade

	incomeID     string
	incomeName   string
	incomeAmount big.Rat
	incomeType   string
	incomeDate   time.Time
	createdAt    time.Time
}

func (s *service) create() error {
	// Convert amount from request (assuming first amount in array)
	var incomeAmount big.Rat
	if len(s.req.IncomeAmount) > 0 && s.req.IncomeAmount[0] != nil {
		incomeAmount.SetFloat64(s.req.IncomeAmount[0].Amount)
	}

	// Generate new ID and timestamp
	incomeID := uuid.NewString()
	createdAt := time.Now().UTC()

	// Create income in database
	err := s.f.pkg.M.FinDash.Income.Create(s.ctx, &m_income.Data{
		IncomeID:     incomeID,
		IncomeName:   spanner.NullString{StringVal: s.req.IncomeName, Valid: true},
		IncomeAmount: spanner.NullNumeric{Numeric: incomeAmount, Valid: true},
		IncomeType:   spanner.NullString{StringVal: s.req.IncomeType, Valid: true},
		IncomeDate:   spanner.NullTime{Time: s.req.IncomeDate, Valid: true},
		CreatedAt:    spanner.NullTime{Time: createdAt, Valid: true},
	})
	if err != nil {
		return errs.FailedToCreateIncome
	}

	// Store data in service for reply
	s.incomeID = incomeID
	s.incomeName = s.req.IncomeName
	s.incomeAmount = incomeAmount
	s.incomeType = s.req.IncomeType
	s.incomeDate = s.req.IncomeDate
	s.createdAt = createdAt

	return nil
}

func (s *service) reply() *di.CreateResponse {
	// Convert big.Rat to float64
	amountValue, _ := s.incomeAmount.Float64()

	// Create amount structure
	amountObj := &amount.Amount{
		Amount:          amountValue,
		AmountFormatted: s.incomeAmount.String(),
		CurrencyCode:    "USD", // Default, adjust as needed
		CurrencySymbol:  "$",   // Default, adjust as needed
	}

	return &di.CreateResponse{
		IncomeID:     s.incomeID,
		IncomeName:   s.incomeName,
		IncomeAmount: []*amount.Amount{amountObj},
		IncomeType:   s.incomeType,
		IncomeDate:   s.incomeDate,
		CreatedAt:    s.createdAt,
	}
}
