package component

import (
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type Component interface {
	// SetupProperty() duerror.DUError, UG 4/29: comment it for now because no implementation
	// CreatePropertyTree() (PropertyTree, duerror.DUError)
	// Copy() (Component, duerror.DUError)
	Cover(p utils.Point) (bool, duerror.DUError)
	GetLayer() int
	SetLayer(layer int) duerror.DUError
	SetIsSelected(isSelected bool) duerror.DUError
	GetDrawData() any
	RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError
}
