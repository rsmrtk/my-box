package anomalies

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToCalculateAverages *err.HTTPError
	FailedToFindAnomalies     *err.HTTPError
	FailedToScanData          *err.HTTPError
}{
	FailedToCalculateAverages: err.NewHTTPError(http.StatusInternalServerError, "Failed to calculate category averages."),
	FailedToFindAnomalies:     err.NewHTTPError(http.StatusInternalServerError, "Failed to find expense anomalies."),
	FailedToScanData:          err.NewHTTPError(http.StatusInternalServerError, "Failed to scan anomaly data."),
}
