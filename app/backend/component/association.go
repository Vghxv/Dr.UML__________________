package component

import (
	"Dr.UML/app/backend/duerror"
)

type Association struct {
	parents [2]*Gadget
}

func (a *Association) GetParentStart() (*Gadget, duerror.DuError) {
	return a.parents[0], nil
}

func (a *Association) GetParentEnd() (*Gadget, duerror.DuError) {
	return a.parents[1], nil
}

func (a *Association) SetParentStart(gadget *Gadget) duerror.DuError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	a.parents[0] = gadget
	return nil
}

func (a *Association) SetParentEnd(gadget *Gadget) duerror.DuError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	a.parents[1] = gadget
	return nil
}
