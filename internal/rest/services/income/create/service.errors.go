package create

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToCreateIncome *err.HTTPError
}{
	FailedToCreateIncome: err.NewHTTPError(http.StatusNotFound, "Failed to create income."),
}
