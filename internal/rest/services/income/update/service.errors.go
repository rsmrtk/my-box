package update

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	IncomeNotFound       *err.HTTPError
	InvalidIncomeID      *err.HTTPError
	FailedToUpdateIncome *err.HTTPError
}{
	IncomeNotFound:       err.NewHTTPError(http.StatusNotFound, "Income not found."),
	InvalidIncomeID:      err.NewHTTPError(http.StatusBadRequest, "Invalid income ID format."),
	FailedToUpdateIncome: err.NewHTTPError(http.StatusInternalServerError, "Failed to update income."),
}
