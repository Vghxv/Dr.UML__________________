package component

import (
	"Dr.uml/backend/component/attribute"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type GadgetType int

const (
	Class               GadgetType = 1 << iota // 0x01
	supportedGadgetType            = Class
)

var AllGadgetTypes = []struct {
	Value  GadgetType
	Number int
}{
	Class: {Class, 1},
}

type Gadget struct {
	gadgetType       GadgetType
	point            utils.Point
	layer            int
	attributes       [][]attribute.Attribute // Gadget have multiple sections, each section have multiple attributes
	color            utils.Color
	drawData         drawdata.Gadget
	updateParentDraw func() duerror.DUError
}

/*
component interface
*/

func (g *Gadget) Cover(p utils.Point) (bool, duerror.DUError) {
	tl := g.point                                                                          // top-left
	br := utils.AddPoints(g.point, utils.Point{X: g.drawData.Width, Y: g.drawData.Height}) // bottom-right
	return p.X >= tl.X && p.X <= br.X && p.Y >= tl.Y && p.Y <= br.Y, nil
}

func (g *Gadget) GetLayer() (int, duerror.DUError) {
	return g.layer, nil
}

func (g *Gadget) SetLayer(layer int) duerror.DUError {
	g.layer = layer
	g.drawData.Layer = layer
	if g.updateParentDraw == nil {
		return nil
	}
	return g.updateParentDraw()
}

func (g *Gadget) GetDrawData() (any, duerror.DUError) {
	return g.drawData, nil
}

func (g *Gadget) updateDrawData() duerror.DUError {
	height := drawdata.LineWidth + drawdata.Margin
	maxAttWidth := 0
	atts := make([][]drawdata.Attribute, len(g.attributes))
	for i, attsRow := range g.attributes {
		atts[i] = make([]drawdata.Attribute, 0, len(attsRow))
		for _, att := range attsRow {
			attDrawData, err := att.GetDrawData()
			if err != nil {
				return err
			}
			atts[i] = append(atts[i], attDrawData)
			if attDrawData.Width > maxAttWidth {
				maxAttWidth = attDrawData.Width
			}
			height += attDrawData.Height + drawdata.Margin
		}
		height += drawdata.LineWidth
	}
	width := maxAttWidth + drawdata.Margin*2 + drawdata.LineWidth*2

	g.drawData.GadgetType = int(g.gadgetType)
	g.drawData.X = g.point.X
	g.drawData.Y = g.point.Y
	g.drawData.Layer = g.layer
	g.drawData.Height = height
	g.drawData.Width = width
	g.drawData.Color = g.color.ToHex()
	g.drawData.Attributes = atts

	if g.updateParentDraw == nil {
		return nil
	}
	return g.updateParentDraw()
}

func (g *Gadget) RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError {
	g.updateParentDraw = update
	if err := update(); err != nil {
		return err
	}
	return nil
}

/*
gadget func
*/

// point getter
func (g *Gadget) GetPoint() utils.Point {
	return g.point
}

// point setter
func (g *Gadget) SetPoint(point utils.Point) duerror.DUError {
	g.point = point
	g.drawData.X = point.X
	g.drawData.Y = point.Y
	if g.updateParentDraw == nil {
		return nil
	}
	return g.updateParentDraw()
}

func NewGadget(gadgetType GadgetType, point utils.Point) (*Gadget, duerror.DUError) {
	if gadgetType&supportedGadgetType == 0 {
		return nil, duerror.NewInvalidArgumentError("gadget type is not supported")
	}
	if gadgetType == 0 {
		return nil, duerror.NewInvalidArgumentError("gadget type is 0")
	}
	g := Gadget{
		gadgetType: gadgetType,
		point:      point,
		layer:      0,
		color:      utils.FromHex(drawdata.DefaultGadgetColor),
	}
	err := g.updateDrawData()
	if err != nil {
		return nil, err
	}
	return &g, nil
}
