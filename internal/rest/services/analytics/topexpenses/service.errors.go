package topexpenses

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToGetTopExpenses *err.HTTPError
	FailedToScanData       *err.HTTPError
}{
	FailedToGetTopExpenses: err.NewHTTPError(http.StatusInternalServerError, "Failed to get top expense categories."),
	FailedToScanData:       err.NewHTTPError(http.StatusInternalServerError, "Failed to scan category data."),
}
