package Component

type Component interface {
	SetupProperty() Utils.DUError
	CreatePropertyTree() (PropertyTree, Utils.DUError)
	Copy() (Component, Utils.DUError)
}
