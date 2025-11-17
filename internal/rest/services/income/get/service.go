package get

import (
	"context"
	"math/big"
	"time"

	"github.com/rsmrtk/db-fd-model/m_income"
	di "github.com/rsmrtk/mybox/internal/rest/domain/income"
)

type service struct {
	ctx context.Context
	req *di.GetRequest
	f   *Facade

	incomeID     string
	incomeName   string
	incomeAmount big.Rat
	incomeType   []string
	incomeDate   time.Time
}

func (s *service) find() error {
	rows, err := s.f.pkg.M.FinDash.Income.Get(s.ctx,
		[]m_income.QueryParam{
			{
				Field:    m_income.IncomeID,
				Operator: "=",
				Value:    s.req.IncomeID,
			},
		},
		[]m_income.Field{
			m_income.IncomeID,
			m_income.IncomeName,
			m_income.IncomeAmount,
			m_income.IncomeType,
			m_income.IncomeDate,
			m_income.CreatedAt,
		},
	)
	if err != nil {
		return errs.FailedToFindIncome
	} else if len(rows) == 0 {
		return errs.FailedToFindIncome
	}

	s.incomeID = rows[0].IncomeID
	s.incomeName = rows[0].IncomeName.StringVal
	s.incomeAmount = rows[0].IncomeAmount.Numeric
	s.incomeType = m_income.EnumType(rows[0].IncomeType)

	return nil
}
