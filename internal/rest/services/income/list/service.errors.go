package list

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToListIncomes *err.HTTPError
}{
	FailedToListIncomes: err.NewHTTPError(http.StatusInternalServerError, "Failed to list incomes."),
}
