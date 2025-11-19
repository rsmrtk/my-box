package dashboard

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToCalculateIncome  *err.HTTPError
	FailedToCalculateExpense *err.HTTPError
	FailedToGetCategories    *err.HTTPError
}{
	FailedToCalculateIncome:  err.NewHTTPError(http.StatusInternalServerError, "Failed to calculate income metrics."),
	FailedToCalculateExpense: err.NewHTTPError(http.StatusInternalServerError, "Failed to calculate expense metrics."),
	FailedToGetCategories:    err.NewHTTPError(http.StatusInternalServerError, "Failed to get expense categories."),
}
