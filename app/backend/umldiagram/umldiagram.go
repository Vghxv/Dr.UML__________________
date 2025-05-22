package umldiagram

import (
	"slices"
	"time"

	"Dr.uml/backend/component"
	"Dr.uml/backend/components"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type DiagramType int

const (
	ClassDiagram = 1 << iota // 0x01
	UseCaseDiagram
	SequenceDiagram
	supportedType = ClassDiagram
)

var AllDiagramTypes = []struct {
	Value  DiagramType
	TSName string
}{
	{ClassDiagram, "ClassDiagram"},
}

// Other methods
func validateDiagramType(input DiagramType) duerror.DUError {
	if !(input&supportedType == input && input != 0) {
		return duerror.NewInvalidArgumentError("Invalid diagram type")
	}
	return nil
}

type UMLDiagram struct {
	name            string
	diagramType     DiagramType // e.g., "Class", "UseCase", "Sequence"
	lastModified    time.Time
	startPoint      utils.Point // for dragging and linking ass
	backgroundColor string

	componentsContainer components.Container
	componentsSelected  map[component.Component]bool
	associations        map[*component.Gadget]([2][]*component.Association)

	updateParentDraw func() duerror.DUError
	drawData         drawdata.Diagram
}

// Constructor
func CreateEmptyUMLDiagram(name string, dt DiagramType) (*UMLDiagram, duerror.DUError) {
	// TODO: also check the file is exist or not
	if err := utils.ValidateFilePath(name); err != nil {
		return nil, err
	}
	if err := validateDiagramType(dt); err != nil {
		return nil, err
	}
	return &UMLDiagram{
		name:                name,
		diagramType:         dt,
		lastModified:        time.Now(),
		startPoint:          utils.Point{X: 0, Y: 0},
		backgroundColor:     drawdata.DefaultDiagramColor, // Default white background
		componentsContainer: components.NewContainerMap(),
		associations:        make(map[*component.Gadget][2][]*component.Association),
		componentsSelected:  make(map[component.Component]bool),
		drawData: drawdata.Diagram{
			Margin:    drawdata.Margin,
			LineWidth: drawdata.LineWidth,
			Color:     drawdata.DefaultDiagramColor,
		},
	}, nil
}

func LoadExistUMLDiagram(name string) (*UMLDiagram, duerror.DUError) {
	// TODO
	return CreateEmptyUMLDiagram(name, ClassDiagram)
}

// Getters
func (ud *UMLDiagram) GetName() string {
	return ud.name
}

func (ud *UMLDiagram) GetDiagramType() DiagramType {
	return ud.diagramType
}

func (ud *UMLDiagram) GetLastModified() time.Time {
	return ud.lastModified
}

// Setters
func (ud *UMLDiagram) SetPointGadget(point utils.Point) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.SetPoint(point)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

func (ud *UMLDiagram) SetSetLayerGadget(layer int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.SetLayer(layer)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

func (ud *UMLDiagram) SetColorGadget(colorHexStr string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.SetColor(colorHexStr)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

func (ud *UMLDiagram) SetAttrContentGadget(section int, index int, content string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.SetAttrContent(section, index, content)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}
func (ud *UMLDiagram) SetAttrSizeGadget(section int, index int, size int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.SetAttrSize(section, index, size)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}
func (ud *UMLDiagram) SetAttrStyleGadget(section int, index int, style int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.SetAttrStyle(section, index, style)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

// Methods
func (ud *UMLDiagram) AddGadget(gadgetType component.GadgetType, point utils.Point, layer int, colorHexStr string, header string) duerror.DUError {
	g, err := component.NewGadget(gadgetType, point, layer, colorHexStr, header)
	if err != nil {
		return err
	}
	if err = g.RegisterUpdateParentDraw(ud.updateDrawData); err != nil {
		return err
	}
	if err = ud.componentsContainer.Insert(g); err != nil {
		return err
	}
	ud.associations[g] = [2][]*component.Association{{}, {}}
	return ud.updateDrawData()
}

func (ud *UMLDiagram) StartAddAssociation(point utils.Point) duerror.DUError {
	if err := ud.validatePoint(point); err != nil {
		return err
	}
	ud.startPoint = point
	return nil
}

func (ud *UMLDiagram) EndAddAssociation(assType component.AssociationType, endPoint utils.Point) duerror.DUError {
	stPoint := ud.startPoint
	ud.startPoint = utils.Point{X: 0, Y: 0}
	if err := ud.validatePoint(endPoint); err != nil {
		return err
	}

	// search parents
	stGad, err := ud.componentsContainer.SearchGadget(stPoint)
	if err != nil {
		return err
	}
	if stGad == nil {
		return duerror.NewInvalidArgumentError("start point does not contain a gadget")
	}
	enGad, err := ud.componentsContainer.SearchGadget(endPoint)
	if err != nil {
		return err
	}
	if enGad == nil {
		return duerror.NewInvalidArgumentError("end point does not contain a gadget")
	}

	// create association
	parents := [2]*component.Gadget{stGad, enGad}
	a, err := component.NewAssociation(parents, component.AssociationType(assType), stPoint, endPoint)
	if err != nil {
		return err
	}
	if err = a.RegisterUpdateParentDraw(ud.updateDrawData); err != nil {
		return err
	}
	if err = ud.componentsContainer.Insert(a); err != nil {
		return err
	}

	// record it, cant modify the slice, being a value of the map, directly
	tmp := ud.associations[stGad]
	tmp[0] = append(tmp[0], a)
	ud.associations[stGad] = tmp

	tmp = ud.associations[enGad]
	tmp[1] = append(tmp[1], a)
	ud.associations[enGad] = tmp

	return ud.updateDrawData()
}

func (ud *UMLDiagram) RemoveSelectedComponents() duerror.DUError {
	for c := range ud.componentsSelected {
		switch c := c.(type) {
		case *component.Gadget:
			if err := ud.removeGadget(c); err != nil {
				return err
			}
		case *component.Association:
			if err := ud.removeAssociation(c); err != nil {
				return err
			}
		}
	}
	return ud.updateDrawData()
}

func (ud *UMLDiagram) SelectComponent(point utils.Point) duerror.DUError {
	c, err := ud.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	if c == nil {
		return nil
	}
	// if is in componentsSelected remove it, else add it
	if _, ok := ud.componentsSelected[c]; ok {
		delete(ud.componentsSelected, c)
	} else {
		ud.componentsSelected[c] = true
	}
	//ud.componentsSelected[c] = true
	return ud.updateDrawData()
}

func (ud *UMLDiagram) UnselectComponent(point utils.Point) duerror.DUError {
	c, err := ud.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	if c == nil {
		return nil
	}
	delete(ud.componentsSelected, c)
	return ud.updateDrawData()
}

func (ud *UMLDiagram) UnselectAllComponents() duerror.DUError {
	ud.componentsSelected = make(map[component.Component]bool)
	return nil
}

func (ud *UMLDiagram) AddAttributeToGadget(section int, content string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.AddAttribute(section, content)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

func (ud *UMLDiagram) RemoveAttributeFromGadget(section int, index int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.RemoveAttribute(section, index)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

// Private methods
func (ud *UMLDiagram) getSelectedComponent() (component.Component, duerror.DUError) {
	if len(ud.componentsSelected) != 1 {
		return nil, duerror.NewInvalidArgumentError("can only operate on one component")
	}
	for c := range ud.componentsSelected {
		return c, nil
	}
	return nil, duerror.NewInvalidArgumentError("no component selected")
}

func (ud *UMLDiagram) removeGadget(gad *component.Gadget) duerror.DUError {
	if _, ok := ud.associations[gad]; ok {
		for _, a := range ud.associations[gad][0] {
			if err := ud.removeAssociation(a); err != nil {
				return err
			}
		}
		for _, a := range ud.associations[gad][1] {
			if err := ud.removeAssociation(a); err != nil {
				return err
			}
		}
		delete(ud.associations, gad)
	}
	delete(ud.componentsSelected, gad)
	return ud.componentsContainer.Remove(gad)
}

func (ud *UMLDiagram) removeAssociation(a *component.Association) duerror.DUError {
	st := a.GetParentStart()
	en := a.GetParentEnd()
	if _, ok := ud.associations[st]; ok {
		stList := ud.associations[st][0]
		index := slices.Index(stList, a)
		if index >= 0 {
			stList = slices.Delete(stList, index, index+1)
		}
		ud.associations[st] = [2][]*component.Association{stList, ud.associations[st][1]}
	}
	if _, ok := ud.associations[en]; ok {
		enList := ud.associations[en][1]
		index := slices.Index(enList, a)
		if index >= 0 {
			enList = slices.Delete(enList, index, index+1)
		}
		ud.associations[en] = [2][]*component.Association{ud.associations[en][0], enList}
	}
	delete(ud.componentsSelected, a)
	return ud.componentsContainer.Remove(a)
}

func (ud *UMLDiagram) validatePoint(point utils.Point) duerror.DUError {
	if point.X < 0 || point.Y < 0 {
		return duerror.NewInvalidArgumentError("point coordinates must be non-negative")
	}
	return nil
}

// draw
func (ud *UMLDiagram) GetDrawData() drawdata.Diagram {
	return ud.drawData
}

func (ud *UMLDiagram) RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError {
	if update == nil {
		return duerror.NewInvalidArgumentError("update function cannot be nil")
	}
	ud.updateParentDraw = update
	return nil

}

func (ud *UMLDiagram) updateDrawData() duerror.DUError {
	gs := make([]drawdata.Gadget, 0, len(ud.componentsSelected))
	as := make([]drawdata.Association, 0, len(ud.componentsSelected))
	for _, c := range ud.componentsContainer.GetAll() {
		cDrawData := c.GetDrawData()
		if cDrawData == nil {
			continue
		}
		switch c.(type) {
		case *component.Gadget:
			// check if this gadget is in componentsSelected
			if _, ok := ud.componentsSelected[c]; ok {
				gadget := cDrawData.(drawdata.Gadget)
				gadget.IsSelected = true
				gs = append(gs, gadget)
			}
			gs = append(gs, cDrawData.(drawdata.Gadget))
		case *component.Association:
			as = append(as, cDrawData.(drawdata.Association))
		}
	}
	ud.drawData.Gadgets = gs
	ud.drawData.Associations = as
	if ud.updateParentDraw == nil {
		return nil
	}
	return ud.updateParentDraw()
}
