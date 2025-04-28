package component

import (
	"Dr.uml/backend/utils/duerror"
)

type Association struct {
	parents [2]*Gadget
}

func NewAssociation(parents [2]*Gadget) *Association {
	if parents[0] == nil || parents[1] == nil {
		return nil
	}
	return &Association {
		parents: [2]*Gadget{parents[0], parents[1]},
	}
}

func (a *Association) GetParentStart() (*Gadget, duerror.DUError) {
	return a.parents[0], nil
}

func (a *Association) GetParentEnd() (*Gadget, duerror.DUError) {
	return a.parents[1], nil
}

func (a *Association) SetParentStart(gadget *Gadget) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	a.parents[0] = gadget
	return nil
}

func (a *Association) SetParentEnd(gadget *Gadget) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	a.parents[1] = gadget
	return nil
}
