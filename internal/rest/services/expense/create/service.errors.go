package create

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToCreateExpense *err.HTTPError
}{
	FailedToCreateExpense: err.NewHTTPError(http.StatusInternalServerError, "Failed to create expense."),
}
