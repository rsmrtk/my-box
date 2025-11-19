package delete

import (
	"context"

	"github.com/google/uuid"
	"github.com/rsmrtk/db-fd-model/m_expense"
	"github.com/rsmrtk/mybox/internal/rest/domain/expense"
)

type service struct {
	ctx context.Context
	req *expense.DeleteRequest
	f   *Facade
}

func (s *service) delete() error {
	_, err := uuid.Parse(s.req.ExpenseID)
	if err != nil {
		return errs.InvalidExpenseID
	}

	pk := m_expense.PrimaryKey{
		ExpenseID: s.req.ExpenseID,
	}

	err = s.f.pkg.M.FinDash.Expense.Delete(s.ctx, pk)
	if err != nil {
		return errs.FailedToDeleteExpense
	}

	return nil
}

func (s *service) reply() *expense.DeleteResponse {
	return &expense.DeleteResponse{
		Success: true,
		Message: "Expense deleted successfully",
	}
}
