package get

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	ExpenseNotFound  *err.HTTPError
	InvalidExpenseID *err.HTTPError
}{
	ExpenseNotFound:  err.NewHTTPError(http.StatusNotFound, "Expense not found."),
	InvalidExpenseID: err.NewHTTPError(http.StatusBadRequest, "Invalid expense ID format."),
}
