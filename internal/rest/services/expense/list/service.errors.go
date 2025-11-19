package list

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToListExpenses *err.HTTPError
}{
	FailedToListExpenses: err.NewHTTPError(http.StatusInternalServerError, "Failed to list expenses."),
}
