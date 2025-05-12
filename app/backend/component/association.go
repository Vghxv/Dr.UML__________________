package component

import (
	"math"

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

	startPointRatio [2]float64
	endPointRatio   [2]float64
}

// Constructor
func NewAssociation(parents [2]*Gadget, assType AssociationType, stPoint utils.Point, enPoint utils.Point) (*Association, duerror.DUError) {
	if assType&supportedAssociationType != assType || assType == 0 {
		return nil, duerror.NewInvalidArgumentError("unsupported association type")
	}
	if parents[0] == nil || parents[1] == nil {
		return nil, duerror.NewInvalidArgumentError("parents are nil")
	}
	stGdd := parents[0].GetDrawData().(*drawdata.Gadget)
	enGdd := parents[1].GetDrawData().(*drawdata.Gadget)
	a := &Association{
		assType: assType,
		parents: [2]*Gadget{parents[0], parents[1]},
		startPointRatio: [2]float64{
			float64(stPoint.X-stGdd.X) / float64(stGdd.Width),
			float64(stPoint.Y-stGdd.Y) / float64(stGdd.Height)},
		endPointRatio: [2]float64{
			float64(enPoint.X-enGdd.X) / float64(enGdd.Width),
			float64(enPoint.Y-enGdd.Y) / float64(enGdd.Height)},
	}
	a.updateDrawData()
	return a, nil
}

// other function
func snapToEdge(rec utils.Point, width int, height int, ratio [2]float64) utils.Point {
	leftDist := ratio[0]
	rightDist := 1 - ratio[0]
	topDist := ratio[1]
	bottomDist := 1 - ratio[1]

	minDist := math.Min(math.Min(leftDist, rightDist), math.Min(topDist, bottomDist))
	switch minDist {
	case leftDist:
		return utils.Point{X: rec.X, Y: rec.Y + int(ratio[1]*float64(height))}
	case rightDist:
		return utils.Point{X: rec.X + width, Y: rec.Y + int(ratio[1]*float64(height))}
	case topDist:
		return utils.Point{X: rec.X + int(ratio[0]*float64(width)), Y: rec.Y}
	default:
		return utils.Point{X: rec.X + int(ratio[0]*float64(width)), Y: rec.Y + height}
	}
}

func dist(st utils.Point, en utils.Point, p utils.Point) float64 {
	stX, stY := float64(st.X), float64(st.Y)
	enX, enY := float64(en.X), float64(en.Y)
	x, y := float64(p.X), float64(p.Y)

	dx := enX - stX
	dy := enY - stY
	if dx == 0 && dy == 0 {
		return math.Hypot(x-stX, y-stY)
	}

	t := ((x-stX)*dx + (y-stY)*dy) / (dx*dx + dy*dy)
	t = math.Max(0, math.Min(1, t))

	xx := stX + t*dx
	yy := stY + t*dy
	return math.Hypot(x-xx, y-yy)
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

func (this *Association) GetStartRatio() [2]float64 {
	return this.startPointRatio
}

func (this *Association) GetEndRatio() [2]float64 {
	return this.endPointRatio
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

func (this *Association) SetParentStart(gadget *Gadget, point utils.Point) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	this.parents[0] = gadget
	return this.SetStartPoint(point)
}

func (this *Association) SetParentEnd(gadget *Gadget, point utils.Point) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	this.parents[1] = gadget
	return this.SetEndPoint(point)
}

func (this *Association) SetStartPoint(point utils.Point) duerror.DUError {
	if this.parents[0] == nil {
		return duerror.NewInvalidArgumentError("parent is nil")
	}
	gdd := this.parents[0].GetDrawData().(*drawdata.Gadget)
	if point.X < gdd.X || point.X > gdd.X+gdd.Width || point.Y < gdd.Y || point.Y > gdd.Y+gdd.Height {
		return duerror.NewInvalidArgumentError("point is out of range")
	}
	this.startPointRatio[0] = float64(point.X-gdd.X) / float64(gdd.Width)
	this.startPointRatio[1] = float64(point.Y-gdd.Y) / float64(gdd.Height)
	return this.updateDrawData()
}

func (this *Association) SetEndPoint(point utils.Point) duerror.DUError {
	if this.parents[1] == nil {
		return duerror.NewInvalidArgumentError("parent is nil")
	}
	gdd := this.parents[1].GetDrawData().(*drawdata.Gadget)
	if point.X < gdd.X || point.X > gdd.X+gdd.Width || point.Y < gdd.Y || point.Y > gdd.Y+gdd.Height {
		return duerror.NewInvalidArgumentError("point is out of range")
	}
	this.endPointRatio[0] = float64(point.X-gdd.X) / float64(gdd.Width)
	this.endPointRatio[1] = float64(point.Y-gdd.Y) / float64(gdd.Height)
	return this.updateDrawData()
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
	if this.parents[0] == nil || this.parents[1] == nil {
		return false, duerror.NewInvalidArgumentError("parents are nil")
	}

	st := utils.Point{this.drawdata.StartX, this.drawdata.StartY}
	en := utils.Point{this.drawdata.EndX, this.drawdata.EndY}
	delta := utils.Point{this.drawdata.DeltaX, this.drawdata.DeltaY}
	stDelta := utils.AddPoints(st, delta)
	enDelta := utils.AddPoints(en, delta)

	// TODO: correct threshold
	threshold := float64(5)
	return dist(stDelta, enDelta, p) <= threshold ||
		dist(st, stDelta, p) <= threshold ||
		dist(en, enDelta, p) <= threshold, nil
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

	if this.parents[0] != this.parents[1] {
		stGdd := this.parents[0].GetDrawData().(*drawdata.Gadget)
		startPoint := snapToEdge(utils.Point{stGdd.X, stGdd.Y}, stGdd.Width, stGdd.Height, this.startPointRatio)
		enGdd := this.parents[1].GetDrawData().(*drawdata.Gadget)
		endPoint := snapToEdge(utils.Point{enGdd.X, enGdd.Y}, enGdd.Width, enGdd.Height, this.endPointRatio)

		this.drawdata.StartX = startPoint.X
		this.drawdata.StartY = startPoint.Y
		this.drawdata.EndX = endPoint.X
		this.drawdata.EndY = endPoint.Y
		this.drawdata.DeltaX = 0
		this.drawdata.DeltaY = 0
	} else {
		// TODO: parents are same
	}

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
