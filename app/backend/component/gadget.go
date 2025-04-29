package component

import (
	"Dr.uml/backend/component/attribute"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type GadgetType string

const (
	Class GadgetType = "Class"
	Note  GadgetType = "Note"
)

type Gadget struct {
	point      utils.Point
	GadgetType GadgetType

	layer     int
	attribute []attribute.Attribute
}

func (g *Gadget) getBounds() (utils.Point, utils.Point, duerror.DUError) {
	//TODO: calculate the Bottom-Right point (maybe store it?)
	size := 5
	return g.point, utils.AddPoints(g.point, utils.Point{X: size, Y: size}), nil
}

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
