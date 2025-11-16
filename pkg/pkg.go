package pkg

import (
	"github.com/rsmrtk/mybox/pkg/pkg_model"
	lg "github.com/rsmrtk/smartlg"
)

type Facade struct {
	Log *lg.Logger
	M   *pkg_model.Models
}
