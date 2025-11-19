package delete

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	InvalidIncomeID      *err.HTTPError
	FailedToDeleteIncome *err.HTTPError
}{
	InvalidIncomeID:      err.NewHTTPError(http.StatusBadRequest, "Invalid income ID format."),
	FailedToDeleteIncome: err.NewHTTPError(http.StatusInternalServerError, "Failed to delete income."),
}
