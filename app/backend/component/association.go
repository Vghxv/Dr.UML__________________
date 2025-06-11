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
	if err := a.UpdateDrawData(); err != nil {
		return nil, err
	}
	if err := a.RegisterAsObserver(); err != nil {
		return nil, err
	}
	return a, nil
}

func FromSavedAssociation(saved utils.SavedAss, parents [2]*Gadget) (*Association, duerror.DUError) {
	if parents[0] == nil || parents[1] == nil {
		return nil, duerror.NewInvalidArgumentError("At least one of the parent is nil")
	}
	ass := &Association{
		assType:         AssociationType(saved.AssType),
		layer:           saved.Layer,
		parents:         parents,
		startPointRatio: saved.StartPointRatio,
		endPointRatio:   saved.EndPointRatio,
	}

	if err := ass.UpdateDrawData(); err != nil {
		return nil, err
	}

	// Register this association as an observer of its parent gadgets
	if err := ass.RegisterAsObserver(); err != nil {
		return nil, err
	}

	return ass, nil
}

func (ass *Association) ToSavedAssociation(parents [2]int) utils.SavedAss {
	savedAss := utils.SavedAss{
		AssType:         int(ass.assType),
		Layer:           ass.layer,
		StartPointRatio: ass.startPointRatio,
		EndPointRatio:   ass.endPointRatio,
		Attributes:      make([]utils.SavedAtt, 0, len(ass.attributes)),
	}
	savedAss.Parents = []int{parents[0], parents[1]}

	for _, att := range ass.attributes {
		savedAss.Attributes = append(savedAss.Attributes, att.ToSavedAssAttribute())
	}

	return savedAss
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

func validateRatio(ratio *[2]float64) duerror.DUError {
	if ratio[0] < 0 || ratio[0] > 1 || ratio[1] < 0 || ratio[1] > 1 {
		return duerror.NewInvalidArgumentError("invalid ratio")
	}
	return nil
}

func CalAssociationPointRatio(g *Gadget, point utils.Point) ([2]float64, duerror.DUError) {
	gdd := g.GetDrawData().(drawdata.Gadget)
	if point.X < gdd.X || point.X > gdd.X+gdd.Width || point.Y < gdd.Y || point.Y > gdd.Y+gdd.Height {
		return [2]float64{}, duerror.NewInvalidArgumentError("point is out of range")
	}
	return [2]float64{
		float64(point.X-gdd.X) / float64(gdd.Width),
		float64(point.Y-gdd.Y) / float64(gdd.Height),
	}, nil
}

// Getters
func (ass *Association) GetAssType() AssociationType {
	return ass.assType
}

func (ass *Association) GetAttributes() []*attribute.AssAttribute {
	return ass.attributes
}

func (ass *Association) GetAttributesLen() int {
	return len(ass.attributes)
}

func (ass *Association) GetAttribute(index int) (*attribute.AssAttribute, duerror.DUError) {
	// to implement command pattern, require "original" data before dong setter
	// using this function as attribute's getter, no changes to the attribute
	// beside expose attribute, this function also make some error checking duplicate, may need to refactor
	if err := ass.validateIndex(index); err != nil {
		return nil, err
	}
	return ass.attributes[index], nil
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

func (ass *Association) GetIsSelected() bool {
	return ass.isSelected
}

// Setters
func (ass *Association) SetIsSelected(value bool) duerror.DUError {
	ass.isSelected = value
	ass.drawdata.IsSelected = value
	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return ass.updateParentDraw()
}

func (ass *Association) SetAssType(assType AssociationType) duerror.DUError {
	if assType&supportedAssociationType != assType || assType == 0 {
		return duerror.NewInvalidArgumentError("unsupported association type")
	}
	ass.assType = assType
	ass.drawdata.AssType = int(assType)
	if ass.updateParentDraw == nil {
		return nil
	}
	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return ass.updateParentDraw()
}

func (ass *Association) SetLayer(layer int) duerror.DUError {
	ass.layer = layer
	ass.drawdata.Layer = layer
	if ass.updateParentDraw == nil {
		return nil
	}
	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return ass.updateParentDraw()

}

func (ass *Association) SetParentStart(gadget *Gadget, ratio [2]float64) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	if err := validateRatio(&ratio); err != nil {
		return err
	}

	// Unregister from old parent if different
	// This ensures the observer does not receive updates from a gadget it is no longer associated with.
	oldParent := ass.parents[0]
	if oldParent != nil && oldParent != gadget {
		oldParent.RemoveObserver(ass)
	}

	ass.parents[0] = gadget
	ass.startPointRatio[0] = ratio[0]
	ass.startPointRatio[1] = ratio[1]

	// Register with new parent
	if err := gadget.AddObserver(ass, ass.UpdateDrawData); err != nil {
		return err
	}

	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ass *Association) SetParentEnd(gadget *Gadget, ratio [2]float64) duerror.DUError {
	if gadget == nil {
		return duerror.NewInvalidArgumentError("gadget is nil")
	}
	if err := validateRatio(&ratio); err != nil {
		return err
	}

	// Unregister from old parent if different
	oldParent := ass.parents[1]
	if oldParent != nil && oldParent != gadget {
		oldParent.RemoveObserver(ass)
	}

	ass.parents[1] = gadget
	ass.endPointRatio[0] = ratio[0]
	ass.endPointRatio[1] = ratio[1]

	if err := gadget.AddObserver(ass, ass.UpdateDrawData); err != nil {
		return err
	}

	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ass *Association) SetAttrContent(index int, content string) duerror.DUError {
	if err := ass.validateIndex(index); err != nil {
		return err
	}
	if err := ass.attributes[index].SetContent(content); err != nil {
		return err
	}
	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ass *Association) SetAttrSize(index int, size int) duerror.DUError {
	if err := ass.validateIndex(index); err != nil {
		return err
	}
	if err := ass.attributes[index].SetSize(size); err != nil {
		return err
	}
	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ass *Association) SetAttrStyle(index int, style int) duerror.DUError {
	if err := ass.validateIndex(index); err != nil {
		return err
	}
	if err := ass.attributes[index].SetStyle(attribute.Textstyle(style)); err != nil {
		return err
	}
	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ass *Association) SetAttrFontFile(index int, fontFile string) duerror.DUError {
	if err := ass.validateIndex(index); err != nil {
		return err
	}
	if err := ass.attributes[index].SetFontFile(fontFile); err != nil {
		return err
	}

	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ass *Association) SetAttrRatio(index int, ratio float64) duerror.DUError {
	if err := ass.validateIndex(index); err != nil {
		return err
	}
	if err := ass.attributes[index].SetRatio(ratio); err != nil {
		return err
	}

	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
}

// Other methods
func (ass *Association) Cover(p utils.Point) (bool, duerror.DUError) {
	if ass.parents[0] == nil || ass.parents[1] == nil {
		return false, duerror.NewInvalidArgumentError("parents are nil")
	}

	st := utils.Point{X: ass.drawdata.StartX, Y: ass.drawdata.StartY}
	en := utils.Point{X: ass.drawdata.EndX, Y: ass.drawdata.EndY}
	delta := utils.Point{X: ass.drawdata.DeltaX, Y: ass.drawdata.DeltaY}
	stDelta := utils.AddPoints(st, delta)
	enDelta := utils.AddPoints(en, delta)

	threshold := float64(4)
	return dist(stDelta, enDelta, p) <= threshold ||
		dist(st, stDelta, p) <= threshold ||
		dist(en, enDelta, p) <= threshold, nil
}

func (ass *Association) AddAttribute(index int, ratio float64, content string) duerror.DUError {
	if index < -1 || index > len(ass.attributes) {
		return duerror.NewInvalidArgumentError("index not allow")
	}

	att, err := attribute.NewAssAttribute(ratio, content)
	if err != nil {
		return err
	}
    if err := att.RegisterUpdateParentDraw(ass.UpdateDrawData); err != nil {
		return err
    }
	ass.attributes = append(ass.attributes, att)

	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ass *Association) AddLoadedAttribute(att *attribute.AssAttribute) duerror.DUError {
	att.RegisterUpdateParentDraw(ass.UpdateDrawData)
	ass.attributes = append(ass.attributes, att)

	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
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

	if err := ass.UpdateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ass *Association) UpdateDrawData() duerror.DUError {
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
		ass.drawdata.Attributes[i] = att.GetDrawData()
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

// Observer pattern methods for associations
func (ass *Association) RegisterAsObserver() duerror.DUError {
	// Register this association as an observer of its parent gadgets
	if ass.parents[0] != nil {
		if err := ass.parents[0].AddObserver(ass, ass.UpdateDrawData); err != nil {
			return err
		}
	}
	if ass.parents[1] != nil && ass.parents[1] != ass.parents[0] {
		if err := ass.parents[1].AddObserver(ass, ass.UpdateDrawData); err != nil {
			return err
		}
	}
	return nil
}

func (ass *Association) UnregisterAsObserver() duerror.DUError {
	// Unregister this association as an observer of its parent gadgets
	if ass.parents[0] != nil {
		if err := ass.parents[0].RemoveObserver(ass); err != nil {
			// Don't return error if observer wasn't found, just continue
		}
	}
	if ass.parents[1] != nil && ass.parents[1] != ass.parents[0] {
		if err := ass.parents[1].RemoveObserver(ass); err != nil {
			// Don't return error if observer wasn't found, just continue
		}
	}
	return nil
}

func (ass *Association) validateIndex(index int) duerror.DUError {
	if index < 0 || index >= len(ass.attributes) {
		return duerror.NewInvalidArgumentError("index out of range")
	}
	return nil
}
