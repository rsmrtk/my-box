package create

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/rsmrtk/db-fd-model/m_income"
	di "github.com/rsmrtk/mybox/internal/rest/domain/income"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
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
	var incomeAmountFloat float64
	if len(s.req.IncomeAmount) > 0 && s.req.IncomeAmount[0] != nil {
		incomeAmountFloat = s.req.IncomeAmount[0].Amount
	}

	// Generate new ID and timestamp
	incomeID := uuid.New().String()
	createdAt := time.Now().UTC()

	// Create income in database using new pgx models
	// Convert values to pointers for nullable fields
	incomeName := s.req.IncomeName
	incomeType := s.req.IncomeType
	incomeDate := s.req.IncomeDate

	err := s.f.pkg.M.FinDash.Income.Create(s.ctx, &m_income.Data{
		IncomeID:     incomeID,
		IncomeName:   &incomeName,
		IncomeAmount: &incomeAmountFloat,
		IncomeType:   &incomeType,
		IncomeDate:   &incomeDate,
		CreatedAt:    &createdAt,
	})
	if err != nil {
		return errs.FailedToCreateIncome
	}

	// Store data in service for reply
	s.incomeID = incomeID
	s.incomeName = s.req.IncomeName
	s.incomeAmount = *big.NewRat(int64(incomeAmountFloat*100), 100) // Convert to big.Rat
	s.incomeType = s.req.IncomeType
	s.incomeDate = s.req.IncomeDate
	s.createdAt = createdAt

	return nil
}

func (s *service) reply() *di.CreateResponse {
	// Convert big.Rat to float64
	amountValue, _ := s.incomeAmount.Float64()

	// Create amount structure
	amountObj := &models.Amount{
		Amount:         amountValue,
		CurrencyCode:   "USD", // Default, adjust as needed
		CurrencySymbol: "$",   // Default, adjust as needed
	}

	return &di.CreateResponse{
		IncomeID:     s.incomeID,
		IncomeName:   s.incomeName,
		IncomeAmount: []*models.Amount{amountObj},
		IncomeType:   s.incomeType,
		IncomeDate:   s.incomeDate,
		CreatedAt:    s.createdAt,
	}
}
