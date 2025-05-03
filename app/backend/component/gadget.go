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

var gadgetDefaultAtts = map[GadgetType]([][]string){
	Class: [][]string{
		{"Name"},
		{"Attributes"},
		{"Methods"},
	},
}

type Gadget struct {
	gadgetType       GadgetType
	point            utils.Point
	layer            int
	attributes       [][]*attribute.Attribute // Gadget have multiple sections, each section have multiple attributes
	color            utils.Color
	drawData         drawdata.Gadget
	updateParentDraw func() duerror.DUError
}

func NewGadget(gadgetType GadgetType, point utils.Point) (*Gadget, duerror.DUError) {
	if err := validateGadgetType(gadgetType); err != nil {
		return nil, err
	}
	g := Gadget{
		gadgetType: gadgetType,
		point:      point,
		layer:      0,
		color:      utils.FromHex(drawdata.DefaultGadgetColor),
	}

	// Init default attributes
	g.attributes = make([][]*attribute.Attribute, len(gadgetDefaultAtts[gadgetType]))
	for i, contents := range gadgetDefaultAtts[gadgetType] {
		g.attributes[i] = make([]*attribute.Attribute, 0, len(contents))
		for _, content := range contents {
			if err := g.AddAttribute(content, i); err != nil {
				return nil, err
			}
		}
	}

	if err := g.updateDrawData(); err != nil {
		return nil, err
	}
	return &g, nil
}

// functions
func validateGadgetType(input GadgetType) duerror.DUError {
	if !(input&supportedGadgetType == input && input != 0) {
		return duerror.NewInvalidArgumentError("gadget type is not supported")
	}
	return nil
}

// Getter
func (g *Gadget) GetPoint() utils.Point {
	return g.point
}

func (g *Gadget) GetLayer() int {
	return g.layer
}

func (g *Gadget) GetColor() utils.Color {
	return g.color
}

func (g *Gadget) GetGadgetType() GadgetType {
	return g.gadgetType
}

// Setter
func (g *Gadget) SetPoint(point utils.Point) duerror.DUError {
	g.point = point
	g.drawData.X = point.X
	g.drawData.Y = point.Y
	if g.updateParentDraw == nil {
		return nil
	}
	return g.updateParentDraw()
}

func (g *Gadget) SetLayer(layer int) duerror.DUError {
	g.layer = layer
	g.drawData.Layer = layer
	if g.updateParentDraw == nil {
		return nil
	}
	return g.updateParentDraw()
}

func (g *Gadget) SetColor(color utils.Color) duerror.DUError {
	g.color = color
	g.drawData.Color = color.ToHex()
	if g.updateParentDraw == nil {
		return nil
	}
	return g.updateParentDraw()
}

// Methods
func (g *Gadget) Cover(p utils.Point) (bool, duerror.DUError) {
	tl := g.point                                                                          // top-left
	br := utils.AddPoints(g.point, utils.Point{X: g.drawData.Width, Y: g.drawData.Height}) // bottom-right
	return p.X >= tl.X && p.X <= br.X && p.Y >= tl.Y && p.Y <= br.Y, nil
}

func (g *Gadget) AddAttribute(content string, section int) duerror.DUError {
	if section < 0 || section >= len(g.attributes) {
		return duerror.NewInvalidArgumentError("section out of range")
	}
	att, err := attribute.NewAttribute(content)
	if err != nil {
		return err
	}
	if err = att.RegisterUpdateParentDraw(g.updateDrawData); err != nil {
		return err
	}
	g.attributes[section] = append(g.attributes[section], att)
	return g.updateDrawData()
}

func (g *Gadget) RemoveAttribute(index int, section int) duerror.DUError {
	if section < 0 || section >= len(g.attributes) {
		return duerror.NewInvalidArgumentError("section out of range")
	}
	if index < 0 || index >= len(g.attributes[section]) {
		return duerror.NewInvalidArgumentError("index out of range")
	}
	g.attributes[section] = append(g.attributes[section][:index], g.attributes[section][index+1:]...)
	return g.updateDrawData()
}

// Draw
func (g *Gadget) GetDrawData() any {
	return g.drawData
}

func (g *Gadget) updateDrawData() duerror.DUError {
	height := drawdata.LineWidth + drawdata.Margin
	maxAttWidth := 0
	atts := make([][]drawdata.Attribute, len(g.attributes))
	for i, attsRow := range g.attributes {
		atts[i] = make([]drawdata.Attribute, 0, len(attsRow))
		for _, att := range attsRow {
			attDrawData := att.GetDrawData()
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
	if update == nil {
		return duerror.NewInvalidArgumentError("update function is nil")
	}
	g.updateParentDraw = update
	return nil
}
