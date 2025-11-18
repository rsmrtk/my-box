package get

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToFindIncome *err.HTTPError
}{
	FailedToFindIncome: err.NewHTTPError(http.StatusNotFound, "Failed to find income."),
}
