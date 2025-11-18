package create

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	// TODO: Uncomment when m_expense is available in db-fd-model
	// m_expense "github.com/rsmrtk/db-fd-model/m_expense"
	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

// TODO: Replace with m_expense.Data when available
type expenseData struct {
	ExpenseID     uuid.UUID
	ExpenseName   *string
	ExpenseAmount *float64
	ExpenseType   *string
	ExpenseDate   *time.Time
	CreatedAt     *time.Time
}

type service struct {
	ctx  context.Context
	req  *expense.CreateRequest
	f    *Facade
	data *expenseData // TODO: Change to *m_expense.Data when available
}

func (s *service) create() error {
	// Convert amount array to float64
	var expenseAmount float64
	if len(s.req.ExpenseAmount) > 0 {
		expenseAmount = s.req.ExpenseAmount[0].Amount
	}

	// Generate new UUID for expense
	expenseID := uuid.New()

	// Create the expense record
	createdAt := time.Now()
	s.data = &expenseData{
		ExpenseID:     expenseID,
		ExpenseName:   &s.req.ExpenseName,
		ExpenseAmount: &expenseAmount,
		ExpenseType:   &s.req.ExpenseType,
		ExpenseDate:   &s.req.ExpenseDate,
		CreatedAt:     &createdAt,
	}

	// TODO: Uncomment when Expense model is available in db-fd-model
	// err := s.f.pkg.M.FinDash.Expense.Create(s.ctx, s.data)
	// if err != nil {
	// 	return errs.FailedToCreateExpense
	// }

	return nil
}

func (s *service) reply() *expense.CreateResponse {
	var expenseName, expenseType string
	var expenseAmount float64
	var expenseDate time.Time

	if s.data.ExpenseName != nil {
		expenseName = *s.data.ExpenseName
	}
	if s.data.ExpenseAmount != nil {
		expenseAmount = *s.data.ExpenseAmount
	}
	if s.data.ExpenseType != nil {
		expenseType = *s.data.ExpenseType
	}
	if s.data.ExpenseDate != nil {
		expenseDate = *s.data.ExpenseDate
	}

	// Convert float64 to big.Rat for precise decimal handling
	expenseAmountRat := *big.NewRat(int64(expenseAmount*100), 100)
	amountValue, _ := expenseAmountRat.Float64()

	return &expense.CreateResponse{
		ExpenseID:   s.data.ExpenseID.String(),
		ExpenseName: expenseName,
		ExpenseAmount: []*models.Amount{{
			Amount:         amountValue,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		ExpenseType: expenseType,
		ExpenseDate: expenseDate,
		CreatedAt:   time.Now(),
	}
}
