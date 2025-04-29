package component

import (
	"Dr.uml/backend/component/attribute"
	"Dr.uml/backend/component/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type Gadget struct {
	gadgetType string
	point utils.Point
	layer int
	attributes []attribute.Attribute
	drawData drawdata.Gadget
	updateParentDraw func() duerror.DUError
}


/*
component interface
*/

func (g *Gadget) Cover(p utils.Point) (bool, duerror.DUError) {
	tl, br, err := g.getBounds()
	if err != nil {
		return false, err
	}
	return p.X >= tl.X && p.X <= br.X && p.Y >= tl.Y && p.Y <= br.Y, nil
}

func (g *Gadget) GetLayer() (int, duerror.DUError) {
	return g.layer, nil
}

func (g *Gadget) SetLayer(layer int) duerror.DUError {
	g.layer = layer
	return nil
}

func (g *Gadget) GetDrawData() (any, duerror.DUError) {
	return g.drawData, nil
}

func (g *Gadget) updateDrawData() duerror.DUError {
	g.drawData = drawdata.Gadget{
		GadgetType: g.gadgetType,
		X: g.point.X,
		Y: g.point.Y,
	}
	if g.updateParentDraw == nil {
		return nil
	}
	return g.updateParentDraw()
}

func (g *Gadget) RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError {
	g.updateParentDraw = update
	return nil
}


/*
gadget func
*/

func (g *Gadget) getBounds() (utils.Point, utils.Point, duerror.DUError) {
	//TODO: calculate the Bottom-Right point (maybe store it?)
	size := 5
	return g.point, utils.AddPoints(g.point, utils.Point{X: size, Y: size}), nil
}
