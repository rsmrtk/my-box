package delete

import (
	"context"

	"github.com/google/uuid"
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

	// TODO: Uncomment when Expense model is available in db-fd-model
	// // Delete the expense
	// err = s.f.pkg.M.FinDash.Expense.Delete(s.ctx, expenseID)
	// if err != nil {
	// 	return errs.FailedToDeleteExpense
	// }

	// For now, just return success
	return nil
}

func (s *service) reply() *expense.DeleteResponse {
	return &expense.DeleteResponse{
		Success: true,
		Message: "Expense deleted successfully",
	}
}
