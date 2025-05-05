package component

import (
	"Dr.uml/backend/component/attribute"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type AssociationType int

const (
	Extension                = 1 << iota // 0x01
	Implementation           = 1 << iota // 0x02
	Composition              = 1 << iota // 0x04
	Dependency               = 1 << iota // 0x08
	supportedAssociationType = Extension | Implementation | Composition | Dependency
)

type Association struct {
	assType          AssociationType
	layer            int
	attributes       []*attribute.AssAttribute
	parents          [2]*Gadget
	drawdata         drawdata.Association
	updateParentDraw func() duerror.DUError
}

// Constructor
func NewAssociation(parents [2]*Gadget, assType AssociationType) (*Association, duerror.DUError) {
	if assType&supportedAssociationType != assType || assType == 0 {
		return nil, duerror.NewInvalidArgumentError("unsupported association type")
	}
	if parents[0] == nil || parents[1] == nil {
		return nil, duerror.NewInvalidArgumentError("parents are nil")
	}
	a := &Association{
		parents: [2]*Gadget{parents[0], parents[1]},
	}
	a.updateDrawData()
	return a, nil
}

// Getters
func (this *Association) GetAssType() AssociationType {
	return this.assType
}

func (this *Association) GetAttributes() ([]*attribute.AssAttribute, duerror.DUError) {
	// TODO: should not do this
	if len(this.attributes) == 0 {
		return nil, duerror.NewInvalidArgumentError("no attributes found")
	}
	return this.attributes, nil
}

func (this *Association) GetDrawData() any {
	return this.drawdata
}

func (this *Association) GetLayer() int {
	return this.layer
}

func (this *Association) GetParentEnd() *Gadget {
	return this.parents[1]
}

func (this *Association) GetParentStart() *Gadget {
	return this.parents[0]
}

// Setters
func (this *Association) SetAssType(assType AssociationType) duerror.DUError {
	if assType&supportedAssociationType != assType || assType == 0 {
		return duerror.NewInvalidArgumentError("unsupported association type")
	}
	this.assType = assType
	this.drawdata.AssType = int(assType)
	if this.updateParentDraw == nil {
		return nil
	}
	return this.updateParentDraw()
}

func (this *Association) SetLayer(layer int) duerror.DUError {
	this.layer = layer
	this.drawdata.Layer = layer
	if this.updateParentDraw == nil {
		return nil
	}
	return this.updateParentDraw()
}

func (this *Association) SetParentStart(gadget *Gadget) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	// TODO: cal the real point
	this.parents[0] = gadget
	this.drawdata.StartX = gadget.GetPoint().X
	this.drawdata.StartY = gadget.GetPoint().Y
	if this.updateParentDraw == nil {
		return nil
	}
	return this.updateParentDraw()
}

func (this *Association) SetParentEnd(gadget *Gadget) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	// TODO: cal the real point
	this.parents[1] = gadget
	this.drawdata.EndX = gadget.GetPoint().X
	this.drawdata.EndY = gadget.GetPoint().Y
	if this.updateParentDraw == nil {
		return nil
	}
	return this.updateParentDraw()
}

// Other methods
func (this *Association) AddAttribute(attribute *attribute.AssAttribute) duerror.DUError {
	if attribute == nil {
		return duerror.NewInvalidArgumentError("attribute is nil")
	}
	attribute.RegisterUpdateParentDraw(this.updateDrawData)
	this.attributes = append(this.attributes, attribute)
	// cuz this is the heaviest part
	return this.updateDrawData()
}

func (this *Association) Cover(p utils.Point) (bool, duerror.DUError) {
	/*TODO*/
	return false, nil
}

func (this *Association) MoveAttribute(index int, ratio float64) duerror.DUError {
	if index < 0 || index >= len(this.attributes) {
		return duerror.NewInvalidArgumentError("index out of range")
	}
	return this.attributes[index].SetRatio(ratio)
}

func (this *Association) RemoveAttribute(index int) duerror.DUError {
	if index < 0 || index >= len(this.attributes) {
		return duerror.NewInvalidArgumentError("index out of range")
	}
	this.attributes = append(this.attributes[:index], this.attributes[index+1:]...)
	return this.updateDrawData()
}

func (this *Association) updateDrawData() duerror.DUError {
	if this == nil || this.parents[0] == nil || this.parents[1] == nil {
		return duerror.NewInvalidArgumentError("association or parents are nil")
	}
	start := this.parents[0].GetPoint()
	end := this.parents[1].GetPoint()

	this.drawdata.StartX = start.X
	this.drawdata.StartY = start.Y
	this.drawdata.EndX = end.X
	this.drawdata.EndY = end.Y
	this.drawdata.AssType = int(this.assType)
	this.drawdata.Attributes = make([]drawdata.AssAttribute, len(this.attributes))

	for i, att := range this.attributes {
		if att == nil {
			return duerror.NewInvalidArgumentError("attribute is nil")
		}
		this.drawdata.Attributes[i] = att.GetAssDD()
	}
	if this.updateParentDraw == nil {
		return nil
	}
	return this.updateParentDraw()
}

func (this *Association) RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError {
	if update == nil {
		return duerror.NewInvalidArgumentError("update function is nil")
	}
	this.updateParentDraw = update
	return nil
}
