package growth

import (
	"net/http"

	err "github.com/rsmrtk/fd-er"
)

var errs = struct {
	FailedToCalculateGrowth *err.HTTPError
}{
	FailedToCalculateGrowth: err.NewHTTPError(http.StatusInternalServerError, "Failed to calculate income growth."),
}
