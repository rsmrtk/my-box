package update

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
}

type service struct {
	ctx  context.Context
	req  *expense.UpdateRequest
	f    *Facade
	data *expenseData // TODO: Change to *m_expense.Data when available
}

func (s *service) update() error {
	expenseID, err := uuid.Parse(s.req.ExpenseID)
	if err != nil {
		return errs.InvalidExpenseID
	}

	// TODO: Uncomment when Expense model is available in db-fd-model
	// // First, fetch the existing expense
	// fields := []string{
	// 	"expense_id",
	// 	"expense_name",
	// 	"expense_amount",
	// 	"expense_type",
	// 	"expense_date",
	// }
	//
	// s.data, err = s.f.pkg.M.FinDash.Expense.Find(s.ctx, expenseID, fields)
	// if err != nil {
	// 	return errs.ExpenseNotFound
	// }

	// Temporary mock data for testing
	mockName := "Test Expense"
	mockAmount := 100.00
	mockType := "Test"
	mockDate := time.Now()

	s.data = &expenseData{
		ExpenseID:     expenseID,
		ExpenseName:   &mockName,
		ExpenseAmount: &mockAmount,
		ExpenseType:   &mockType,
		ExpenseDate:   &mockDate,
	}

	// Update only the provided fields
	if s.req.ExpenseName != "" {
		s.data.ExpenseName = &s.req.ExpenseName
	}

	if len(s.req.ExpenseAmount) > 0 {
		amount := s.req.ExpenseAmount[0].Amount
		s.data.ExpenseAmount = &amount
	}

	if s.req.ExpenseType != "" {
		s.data.ExpenseType = &s.req.ExpenseType
	}

	if s.req.ExpenseDate != nil {
		s.data.ExpenseDate = s.req.ExpenseDate
	}

	// TODO: Uncomment when Expense model is available in db-fd-model
	// // Update the expense
	// err = s.f.pkg.M.FinDash.Expense.Update(s.ctx, s.data)
	// if err != nil {
	// 	return errs.FailedToUpdateExpense
	// }

	return nil
}

func (s *service) reply() *expense.UpdateResponse {
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

	return &expense.UpdateResponse{
		ExpenseID:   s.data.ExpenseID.String(),
		ExpenseName: expenseName,
		ExpenseAmount: []*models.Amount{{
			Amount:         amountValue,
			CurrencyCode:   "USD",
			CurrencySymbol: "$",
		}},
		ExpenseType: expenseType,
		ExpenseDate: expenseDate,
		UpdatedAt:   time.Now(),
	}
}
