package Component

import "Dr.uml/backend/Utils"

type Component interface {
	SetupProperty() Utils.DUError
	CreatePropertyTree() (PropertyTree, Utils.DUError)
	Copy() (Component, Utils.DUError)
}
