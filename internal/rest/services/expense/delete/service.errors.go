package delete

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	InvalidExpenseID      *err.HTTPError
	FailedToDeleteExpense *err.HTTPError
}{
	InvalidExpenseID:      err.NewHTTPError(http.StatusBadRequest, "Invalid expense ID format."),
	FailedToDeleteExpense: err.NewHTTPError(http.StatusInternalServerError, "Failed to delete expense."),
}
