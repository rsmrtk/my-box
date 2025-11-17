package create

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/rsmrtk/db-fd-model/m_income"
	"github.com/rsmrtk/mybox/internal/rest/domain/income"
)

//create
//reply

type service struct {
	ctx context.Context
	req *income.CreateRequest
	f   *Facade

	incomeID     string
	incomeName   string
	incomeAmount big.Rat
	incomeType   string
	incomeDate   time.Time
	createdAt    time.Time
}

func (s *service) create() error {
	createIncome, err := s.f.pkg.M.FinDash.Income.Create(s.ctx, &m_income.Data{
		IncomeID:     uuid.NewString(),
		IncomeName:   s.req.IncomeName,
		IncomeAmount: s.req.IncomeAmount,
		IncomeType:   s.req.IncomeType,
		IncomeDate:   s.req.IncomeDate,
		CreatedAt:    time.Now().UTC(),
	})
	if err != nil {
		return errs.FailedToCreateIncome
	}

	return nil
}

func (s *service) reply() (*di.CreateRe error) {}
