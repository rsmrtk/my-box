package create

import "github.com/rsmrtk/mybox/pkg"

type Facade struct {
	pkg *pkg.Facade
}

func New(pkg *pkg.Facade) *Facade {
	return &Facade{pkg: pkg}
}
