package list

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
	ctx   context.Context
	req   *expense.ListRequest
	f     *Facade
	items []*expenseData // TODO: Change to []*m_expense.Data when available
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

	// TODO: Uncomment when Expense model is available in db-fd-model
	// Define fields to fetch
	// fields := []m_expense.Field{
	// 	m_expense.ExpenseID,
	// 	m_expense.ExpenseName,
	// 	m_expense.ExpenseAmount,
	// 	m_expense.ExpenseType,
	// 	m_expense.ExpenseDate,
	// 	m_expense.CreatedAt,
	// }
	//
	// // Fetch all expenses with pagination
	// var err error
	// s.items, s.total, err = s.f.pkg.M.FinDash.Expense.List(
	// 	s.ctx,
	// 	fields,
	// 	s.req.Limit,
	// 	s.req.Offset,
	// 	s.req.SortBy,
	// 	s.req.Order,
	// )
	// if err != nil {
	// 	return errs.FailedToListExpenses
	// }

	// Temporary mock data for testing
	s.items = make([]*expenseData, 0)

	// Create some mock expenses
	for i := 0; i < 5; i++ {
		mockName := "Test Expense " + string(rune('A'+i))
		mockAmount := float64((i + 1) * 100)
		mockType := "Category " + string(rune('A'+i))
		mockDate := time.Now().AddDate(0, 0, -i)
		mockCreated := time.Now().AddDate(0, 0, -i)

		s.items = append(s.items, &expenseData{
			ExpenseID:     uuid.New(),
			ExpenseName:   &mockName,
			ExpenseAmount: &mockAmount,
			ExpenseType:   &mockType,
			ExpenseDate:   &mockDate,
			CreatedAt:     &mockCreated,
		})
	}

	s.total = len(s.items)

	// Apply pagination to mock data
	if s.req.Offset < len(s.items) {
		end := s.req.Offset + s.req.Limit
		if end > len(s.items) {
			end = len(s.items)
		}
		s.items = s.items[s.req.Offset:end]
	} else {
		s.items = []*expenseData{}
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
		if data.ExpenseName != nil {
			expenseName = *data.ExpenseName
		}
		if data.ExpenseAmount != nil {
			expenseAmount = *data.ExpenseAmount
		}
		if data.ExpenseType != nil {
			expenseType = *data.ExpenseType
		}
		if data.ExpenseDate != nil {
			expenseDate = *data.ExpenseDate
		}
		if data.CreatedAt != nil {
			createdAt = *data.CreatedAt
		}

		// Convert float64 to big.Rat for precise decimal handling
		expenseAmountRat := *big.NewRat(int64(expenseAmount*100), 100)
		amountValue, _ := expenseAmountRat.Float64()

		item := &expense.ListItem{
			ExpenseID:   data.ExpenseID.String(),
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
