package create

import (
	"context"
	"database/sql"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/rsmrtk/db-fd-model/m_expense"
	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx  context.Context
	req  *expense.CreateRequest
	f    *Facade
	data *m_expense.Data
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

	// ExpenseDate is always valid since it's required in the request
	expenseDateNull := sql.NullTime{Time: s.req.ExpenseDate.Time, Valid: true}

	s.data = &m_expense.Data{
		ExpenseID:     expenseID,
		ExpenseName:   s.req.ExpenseName,
		ExpenseAmount: sql.NullFloat64{Float64: expenseAmount, Valid: true},
		ExpenseType:   s.req.ExpenseType,
		ExpenseDate:   expenseDateNull,
		CreatedAt:     sql.NullTime{Time: createdAt, Valid: true},
	}

	err := s.f.pkg.M.FinDash.Expense.Create(s.ctx, s.data)
	if err != nil {
		return errs.FailedToCreateExpense
	}

	return nil
}

func (s *service) reply() *expense.CreateResponse {
	var expenseName, expenseType string
	var expenseAmount float64
	var expenseDate time.Time

	// Handle interface{} types for ExpenseName and ExpenseType
	if s.data.ExpenseName != nil {
		if name, ok := s.data.ExpenseName.(string); ok {
			expenseName = name
		}
	}

	if s.data.ExpenseAmount.Valid {
		expenseAmount = s.data.ExpenseAmount.Float64
	}

	if s.data.ExpenseType != nil {
		if typ, ok := s.data.ExpenseType.(string); ok {
			expenseType = typ
		}
	}

	if s.data.ExpenseDate.Valid {
		expenseDate = s.data.ExpenseDate.Time
	}

	// Convert float64 to big.Rat for precise decimal handling
	expenseAmountRat := *big.NewRat(int64(expenseAmount*100), 100)
	amountValue, _ := expenseAmountRat.Float64()

	// Handle ExpenseID conversion
	expenseIDStr := ""
	if id, ok := s.data.ExpenseID.(uuid.UUID); ok {
		expenseIDStr = id.String()
	}

	return &expense.CreateResponse{
		ExpenseID:   expenseIDStr,
		ExpenseName: expenseName,
		ExpenseAmount: []*models.Amount{{
			Amount:         amountValue,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		ExpenseType: expenseType,
		ExpenseDate: models.NewDate(expenseDate),
		CreatedAt:   models.NewDate(time.Now()),
	}
}
