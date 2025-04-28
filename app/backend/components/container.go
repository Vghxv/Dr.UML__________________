package components

import (
	"Dr.uml/backend/component"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type ComponentsContainer interface {
	Insert(c component.Component) duerror.DUError
	Remove(c component.Component) duerror.DUError
	Search(p utils.Point) (component.Component, duerror.DUError)
	Len() (int, duerror.DUError)
}
