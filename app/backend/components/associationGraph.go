package components

import (
	"Dr.uml/backend/component"
	"Dr.uml/backend/utils/duerror"
)

type AssociationGraph interface {
	FindStartEnd(st *component.Gadget, en *component.Gadget) ([]*component.Association, duerror.DUError)
	FindStart(st *component.Gadget) ([]*component.Association, duerror.DUError)
	FindEnd(en *component.Gadget) ([]*component.Association, duerror.DUError)
	FindEither(g *component.Gadget) ([]*component.Association, duerror.DUError)
	
	Update(a *component.Association, oldSt *component.Gadget, oldEn *component.Gadget) duerror.DUError
	Insert(a *component.Association) duerror.DUError
	Remove(a *component.Association) duerror.DUError
	RemoveGadget(g *component.Gadget) ([]*component.Association, duerror.DUError)
}
