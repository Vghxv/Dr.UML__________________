package component

import (
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type Component interface {
	SetupProperty() duerror.DUError
	CreatePropertyTree() (PropertyTree, duerror.DUError)
	Copy() (Component, duerror.DUError)
	Cover(p utils.Point) (bool, duerror.DUError)
	GetLayer() (int, duerror.DUError)
}
