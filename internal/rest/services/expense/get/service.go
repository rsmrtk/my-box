package get

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/rsmrtk/db-fd-model/m_expense"
	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
	"github.com/rsmrtk/mybox/internal/rest/domain/models"
)

type service struct {
	ctx  context.Context
	req  *expense.GetRequest
	f    *Facade
	data *m_expense.Data
}

func (s *service) find() error {
	_, err := uuid.Parse(s.req.ExpenseID)
	if err != nil {
		return errs.InvalidExpenseID
	}

	pk := m_expense.PrimaryKey{
		ExpenseID: s.req.ExpenseID,
	}

	fields := []m_expense.Field{
		m_expense.ExpenseID,
		m_expense.ExpenseName,
		m_expense.ExpenseAmount,
		m_expense.ExpenseType,
		m_expense.ExpenseDate,
		m_expense.CreatedAt,
	}

	s.data, err = s.f.pkg.M.FinDash.Expense.Find(s.ctx, pk, fields)
	if err != nil {
		return errs.ExpenseNotFound
	}

	return nil
}

func (s *service) reply() *expense.GetResponse {
	var expenseName, expenseType string
	var expenseAmount float64
	var expenseDate, createdAt time.Time

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

	if s.data.CreatedAt.Valid {
		createdAt = s.data.CreatedAt.Time
	}

	// Convert float64 to big.Rat for precise decimal handling
	expenseAmountRat := *big.NewRat(int64(expenseAmount*100), 100)
	amountValue, _ := expenseAmountRat.Float64()

	// Handle ExpenseID conversion
	expenseIDStr := ""
	if id, ok := s.data.ExpenseID.(uuid.UUID); ok {
		expenseIDStr = id.String()
	}

	return &expense.GetResponse{
		ExpenseID:   expenseIDStr,
		ExpenseName: expenseName,
		ExpenseAmount: []*models.Amount{{
			Amount:         amountValue,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		ExpenseType: expenseType,
		ExpenseDate: expenseDate,
		CreatedAt:   createdAt,
	}
}
