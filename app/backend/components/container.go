package components

import (
	"Dr.uml/backend/component"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type Container interface {
	Insert(c component.Component) duerror.DUError
	Remove(c component.Component) duerror.DUError
	Search(p utils.Point) (component.Component, duerror.DUError)
	SearchGadget(p utils.Point) (*component.Gadget, duerror.DUError)
	GetAll() []component.Component
	Len() (int, duerror.DUError)
}
