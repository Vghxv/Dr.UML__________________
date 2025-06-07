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

var AllAssociationTypes = []struct {
	Value  AssociationType
	TSName string
}{
	{Extension, "Extension"},
	{Implementation, "Implementation"},
	{Composition, "Composition"},
	{Dependency, "Dependency"},
}

type Association struct {
	assType          AssociationType
	layer            int
	attributes       []*attribute.AssAttribute
	parents          [2]*Gadget
	drawdata         drawdata.Association
	isSelected       bool
	updateParentDraw func() duerror.DUError
	IsSelected       bool
	startPointRatio  [2]float64
	endPointRatio    [2]float64
}

// Constructor
func NewAssociation(parents [2]*Gadget, assType AssociationType, stPoint utils.Point, enPoint utils.Point) (*Association, duerror.DUError) {
	if assType&supportedAssociationType != assType || assType == 0 {
		return nil, duerror.NewInvalidArgumentError("unsupported association type")
	}
	if parents[0] == nil || parents[1] == nil {
		return nil, duerror.NewInvalidArgumentError("parents are nil")
	}
	stGdd := parents[0].GetDrawData().(drawdata.Gadget)
	enGdd := parents[1].GetDrawData().(drawdata.Gadget)
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
	if err := a.updateDrawData(); err != nil {
		return nil, err
	}
	return a, nil
}

// other function
func snapToEdge(rec utils.Point, width int, height int, ratio [2]float64) utils.Point {
	// snap a point onto the edge of a rectangle, the point is float {xRatio, yRatio}
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
func (ass *Association) GetAssType() AssociationType {
	return ass.assType
}

func (ass *Association) GetAttributes() ([]*attribute.AssAttribute, duerror.DUError) {
	// TODO: should not do ass
	if len(ass.attributes) == 0 {
		return nil, duerror.NewInvalidArgumentError("no attributes found")
	}
	return ass.attributes, nil
}

func (ass *Association) GetDrawData() any {
	return ass.drawdata
}

func (ass *Association) GetLayer() int {
	return ass.layer
}

func (ass *Association) GetParentEnd() *Gadget {
	return ass.parents[1]
}

func (ass *Association) GetParentStart() *Gadget {
	return ass.parents[0]
}

func (ass *Association) GetStartRatio() [2]float64 {
	return ass.startPointRatio
}

func (ass *Association) GetEndRatio() [2]float64 {
	return ass.endPointRatio
}

func (this *Association) GetIsSelected() bool {
	return this.IsSelected
}

// Setters
func (ass *Association) SetAssType(assType AssociationType) duerror.DUError {
	if assType&supportedAssociationType != assType || assType == 0 {
		return duerror.NewInvalidArgumentError("unsupported association type")
	}
	ass.assType = assType
	ass.drawdata.AssType = int(assType)
	if ass.updateParentDraw == nil {
		return nil
	}
	return ass.updateParentDraw()
}

func (ass *Association) SetLayer(layer int) duerror.DUError {
	ass.layer = layer
	ass.drawdata.Layer = layer
	if ass.updateParentDraw == nil {
		return nil
	}
	return ass.updateParentDraw()
}

func (ass *Association) SetParentStart(gadget *Gadget, point utils.Point) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	ass.parents[0] = gadget

	gdd := ass.parents[0].GetDrawData().(drawdata.Gadget)
	if point.X < gdd.X || point.X > gdd.X+gdd.Width || point.Y < gdd.Y || point.Y > gdd.Y+gdd.Height {
		return duerror.NewInvalidArgumentError("point is out of range")
	}
	ass.startPointRatio[0] = float64(point.X-gdd.X) / float64(gdd.Width)
	ass.startPointRatio[1] = float64(point.Y-gdd.Y) / float64(gdd.Height)
	return ass.updateDrawData()
}

func (ass *Association) SetParentEnd(gadget *Gadget, point utils.Point) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	ass.parents[1] = gadget

	gdd := ass.parents[1].GetDrawData().(drawdata.Gadget)
	if point.X < gdd.X || point.X > gdd.X+gdd.Width || point.Y < gdd.Y || point.Y > gdd.Y+gdd.Height {
		return duerror.NewInvalidArgumentError("point is out of range")
	}
	ass.endPointRatio[0] = float64(point.X-gdd.X) / float64(gdd.Width)
	ass.endPointRatio[1] = float64(point.Y-gdd.Y) / float64(gdd.Height)
	return ass.updateDrawData()
}

func (ass *Association) SetIsSelected(value bool) duerror.DUError {
	ass.isSelected = value
	ass.drawdata.IsSelected = value
	return ass.updateParentDraw()
}

// Other methods
func (ass *Association) AddAttribute(attribute *attribute.AssAttribute) duerror.DUError {
	if attribute == nil {
		return duerror.NewInvalidArgumentError("attribute is nil")
	}
	attribute.RegisterUpdateParentDraw(ass.updateDrawData)
	ass.attributes = append(ass.attributes, attribute)
	// cuz ass is the heaviest part
	return ass.updateDrawData()
}

func (ass *Association) Cover(p utils.Point) (bool, duerror.DUError) {
	if ass.parents[0] == nil || ass.parents[1] == nil {
		return false, duerror.NewInvalidArgumentError("parents are nil")
	}

	st := utils.Point{ass.drawdata.StartX, ass.drawdata.StartY}
	en := utils.Point{ass.drawdata.EndX, ass.drawdata.EndY}
	delta := utils.Point{ass.drawdata.DeltaX, ass.drawdata.DeltaY}
	stDelta := utils.AddPoints(st, delta)
	enDelta := utils.AddPoints(en, delta)

	threshold := float64(4)
	return dist(stDelta, enDelta, p) <= threshold ||
		dist(st, stDelta, p) <= threshold ||
		dist(en, enDelta, p) <= threshold, nil
}

func (ass *Association) MoveAttribute(index int, ratio float64) duerror.DUError {
	if index < 0 || index >= len(ass.attributes) {
		return duerror.NewInvalidArgumentError("index out of range")
	}
	return ass.attributes[index].SetRatio(ratio)
}

func (ass *Association) RemoveAttribute(index int) duerror.DUError {
	if index < 0 || index >= len(ass.attributes) {
		return duerror.NewInvalidArgumentError("index out of range")
	}
	ass.attributes = append(ass.attributes[:index], ass.attributes[index+1:]...)
	return ass.updateDrawData()
}

func (ass *Association) updateDrawData() duerror.DUError {
	if ass == nil || ass.parents[0] == nil || ass.parents[1] == nil {
		return duerror.NewInvalidArgumentError("association or parents are nil")
	}

	ass.drawdata.DeltaX = 0
	ass.drawdata.DeltaY = 0
	var startPoint, endPoint utils.Point
	if ass.parents[0] != ass.parents[1] {
		// diff parents: start and end both snap to edges of their parents
		stGdd := ass.parents[0].GetDrawData().(drawdata.Gadget)
		startPoint = snapToEdge(utils.Point{X: stGdd.X, Y: stGdd.Y}, stGdd.Width, stGdd.Height, ass.startPointRatio)
		enGdd := ass.parents[1].GetDrawData().(drawdata.Gadget)
		endPoint = snapToEdge(utils.Point{X: enGdd.X, Y: enGdd.Y}, enGdd.Width, enGdd.Height, ass.endPointRatio)
	} else {
		// same parents: choose a side closest to the start point, and calculate delta
		gdd := ass.parents[0].GetDrawData().(drawdata.Gadget)
		startPoint = snapToEdge(utils.Point{X: gdd.X, Y: gdd.Y}, gdd.Width, gdd.Height, ass.startPointRatio)
		endPoint.X, endPoint.Y = startPoint.X, startPoint.Y
		if startPoint.X == gdd.X || startPoint.X == gdd.X+gdd.Width {
			// left / right
			endPoint.Y = gdd.Y + int(float64(gdd.Height)*ass.endPointRatio[1])
			ass.drawdata.DeltaX = utils.AbsInt(startPoint.Y-endPoint.Y) / 2
			if startPoint.X == gdd.X {
				// left, deltaX is negative
				ass.drawdata.DeltaX = -ass.drawdata.DeltaX
			}
		} else if startPoint.Y == gdd.Y || startPoint.Y == gdd.Y+gdd.Height {
			endPoint.X = gdd.X + int(float64(gdd.Width)*ass.endPointRatio[0])
			// up / bottom
			ass.drawdata.DeltaY = utils.AbsInt(startPoint.X-endPoint.X) / 2
			if startPoint.Y == gdd.Y {
				// up, deltaY is negative
				ass.drawdata.DeltaY = -ass.drawdata.DeltaY
			}
		}
	}

	if utils.EqualPoints(startPoint, endPoint) {
		return duerror.NewInvalidArgumentError("start and end points are the same")
	}

	ass.drawdata.StartX = startPoint.X
	ass.drawdata.StartY = startPoint.Y
	ass.drawdata.EndX = endPoint.X
	ass.drawdata.EndY = endPoint.Y
	ass.drawdata.IsSelected = ass.isSelected
	ass.drawdata.AssType = int(ass.assType)
	ass.drawdata.Attributes = make([]drawdata.AssAttribute, len(ass.attributes))

	for i, att := range ass.attributes {
		if att == nil {
			return duerror.NewInvalidArgumentError("attribute is nil")
		}
		ass.drawdata.Attributes[i] = att.GetAssDD()
	}
	if ass.updateParentDraw == nil {
		return nil
	}
	return ass.updateParentDraw()
}

func (ass *Association) RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError {
	if update == nil {
		return duerror.NewInvalidArgumentError("update function is nil")
	}
	ass.updateParentDraw = update
	return nil
}
