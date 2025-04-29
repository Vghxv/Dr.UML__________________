package component

import (
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type Association struct {
	parents [2]*Gadget
	layer int
}


/*
component interface
*/

func (a *Association) Cover(p utils.Point) (bool, duerror.DUError) {
	return false, nil
}

func (a *Association) GetLayer() (int, duerror.DUError) {
	return a.layer, nil
}

func (a *Association) SetLayer(layer int) duerror.DUError {
	a.layer = layer
	return nil
}

func (a *Association) GetDrawData() (any, duerror.DUError) {
	return nil, nil
}

func (g *Association) updateDrawData() duerror.DUError {
	return nil
}


/*
associaiton func
*/

func NewAssociation(parents [2]*Gadget) (*Association, duerror.DUError) {
	if parents[0] == nil || parents[1] == nil {
		return nil, duerror.NewInvalidArgumentError("parents are nil")
	}
	return &Association {
		parents: [2]*Gadget{parents[0], parents[1]},
	}, nil
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
