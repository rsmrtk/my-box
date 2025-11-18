package list

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
	ctx   context.Context
	req   *expense.ListRequest
	f     *Facade
	items []*m_expense.Data
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

	// Note: Current API doesn't support pagination, sorting, and total count directly
	// This would need to be implemented in the model or handled differently
	queryParams := []m_expense.QueryParam{}

	// Add any filter parameters if needed
	// For now, get all records

	var err error
	s.items, err = s.f.pkg.M.FinDash.Expense.List(s.ctx, queryParams)
	if err != nil {
		return errs.FailedToListExpenses
	}

	// Manual pagination since API doesn't support it
	s.total = len(s.items)
	if s.req.Offset < len(s.items) {
		end := s.req.Offset + s.req.Limit
		if end > len(s.items) {
			end = len(s.items)
		}
		s.items = s.items[s.req.Offset:end]
	} else {
		s.items = []*m_expense.Data{}
	}

	return nil
}

func (s *service) reply() *expense.ListResponse {
	items := make([]*expense.ListItem, 0, len(s.items))

	for _, data := range s.items {
		var expenseName, expenseType string
		var expenseAmount float64
		var expenseDate, createdAt time.Time

		// Extract nullable fields
		// Handle interface{} types for ExpenseName and ExpenseType
		if data.ExpenseName != nil {
			if name, ok := data.ExpenseName.(string); ok {
				expenseName = name
			}
		}

		if data.ExpenseAmount.Valid {
			expenseAmount = data.ExpenseAmount.Float64
		}

		if data.ExpenseType != nil {
			if typ, ok := data.ExpenseType.(string); ok {
				expenseType = typ
			}
		}

		if data.ExpenseDate.Valid {
			expenseDate = data.ExpenseDate.Time
		}

		if data.CreatedAt.Valid {
			createdAt = data.CreatedAt.Time
		}

		// Convert float64 to big.Rat for precise decimal handling
		expenseAmountRat := *big.NewRat(int64(expenseAmount*100), 100)
		amountValue, _ := expenseAmountRat.Float64()

		// Handle ExpenseID conversion
		expenseIDStr := ""
		if id, ok := data.ExpenseID.(uuid.UUID); ok {
			expenseIDStr = id.String()
		}

		item := &expense.ListItem{
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

		items = append(items, item)
	}

	return &expense.ListResponse{
		Items:      items,
		TotalCount: s.total,
		Limit:      s.req.Limit,
		Offset:     s.req.Offset,
	}
}
