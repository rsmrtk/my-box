package trends

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToGetTrends *err.HTTPError
	FailedToScanData  *err.HTTPError
}{
	FailedToGetTrends: err.NewHTTPError(http.StatusInternalServerError, "Failed to get expense trends."),
	FailedToScanData:  err.NewHTTPError(http.StatusInternalServerError, "Failed to scan trend data."),
}
