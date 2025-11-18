package delete

import (
	"context"

	"github.com/google/uuid"
	"github.com/rsmrtk/mybox/internal/rest/domain/income"
)

type service struct {
	ctx context.Context
	req *income.DeleteRequest
	f   *Facade
}

func (s *service) delete() error {
	// Validate the income ID format
	_, err := uuid.Parse(s.req.IncomeID)
	if err != nil {
		return errs.InvalidIncomeID
	}

	// Delete the income
	err = s.f.pkg.M.FinDash.Income.Delete(s.ctx, s.req.IncomeID)
	if err != nil {
		return errs.FailedToDeleteIncome
	}

	return nil
}

func (s *service) reply() *income.DeleteResponse {
	return &income.DeleteResponse{
		Success: true,
		Message: "Income deleted successfully",
	}
}
