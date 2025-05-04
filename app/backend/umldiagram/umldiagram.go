package umldiagram

import (
	"time"

	"Dr.uml/backend/component"
	"Dr.uml/backend/components"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type DiagramType int

// TODO: ug want to refactor all the type things
const (
	ClassDiagram    = 1 << iota // 0x01
	UseCaseDiagram  = 1 << iota // 0x02
	SequenceDiagram = 1 << iota // 0x04
	supportedType   = ClassDiagram | UseCaseDiagram | SequenceDiagram
)

var AllDiagramTypes = []struct {
	Value  DiagramType
	Number int
}{
	{ClassDiagram, 1},
	{UseCaseDiagram, 2},
	{SequenceDiagram, 4},
}

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
	backgroundColor utils.Color

	componentsContainer components.Container
	componentsGraph     components.Graph
	componentsSelected  map[component.Component]bool

	notifyDrawUpdate func() duerror.DUError
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
		backgroundColor:     utils.FromHex(drawdata.DefaultDiagramColor), // Default white background
		componentsContainer: components.NewContainerMap(),
		componentsGraph:     components.NewGraphMap(),
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

// Other methods
func (ud *UMLDiagram) AddGadget(gadgetType component.GadgetType, point utils.Point) duerror.DUError {
	g, err := component.NewGadget(gadgetType, point)
	if err != nil {
		return err
	}
	if err = g.RegisterUpdateParentDraw(ud.updateDrawData); err != nil {
		return err
	}
	if err = ud.componentsContainer.Insert(g); err != nil {
		return err
	}
	return ud.updateDrawData()
}

func (ud *UMLDiagram) removeGadget(gad *component.Gadget) duerror.DUError {
	// TODO
	return nil
}

func (ud *UMLDiagram) validatePoint(point utils.Point) duerror.DUError {
	if point.X < 0 || point.Y < 0 {
		return duerror.NewInvalidArgumentError("point coordinates must be non-negative")
	}
	return nil
}

func (ud *UMLDiagram) StartAddAssociation(point utils.Point) duerror.DUError {
	if err := ud.validatePoint(point); err != nil {
		return err
	}
	ud.startPoint = point
	return nil
}

func (ud *UMLDiagram) EndAddAssociation(assType component.AssociationType, endPoint utils.Point) duerror.DUError {
	if err := ud.validatePoint(endPoint); err != nil {
		return err
	}
	// TODO: serach parents
	// st := ud.startPoint
	// en := endPoint
	parents := [2]*component.Gadget{}
	a, err := component.NewAssociation(parents, component.AssociationType(assType))
	if err != nil {
		return err
	}
	if err = a.RegisterUpdateParentDraw(ud.updateDrawData); err != nil {
		return err
	}
	if err = ud.componentsContainer.Insert(a); err != nil {
		return err
	}
	if err = ud.componentsGraph.Insert(a); err != nil {
		return err
	}
	return ud.updateDrawData()
}

func (ud *UMLDiagram) removeAssociation(a *component.Association) duerror.DUError {
	// TODO
	return nil
}

func (ud *UMLDiagram) RemoveSelectedComponents() duerror.DUError {
	// TODO
	return nil
}

func (ud *UMLDiagram) SelectComponent(point utils.Point) duerror.DUError {
	c, err := ud.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	if c == nil {
		return nil
	}
	ud.componentsSelected[c] = true
	return nil
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
	return nil
}

func (ud *UMLDiagram) UnselectAllComponents() duerror.DUError {
	ud.componentsSelected = make(map[component.Component]bool)
	return nil
}

// Draw
func (ud *UMLDiagram) GetDrawData() drawdata.Diagram {
	return ud.drawData
}

func (ud *UMLDiagram) RegisterNotifyDrawUpdate(update func() duerror.DUError) duerror.DUError {
	if update == nil {
		return duerror.NewInvalidArgumentError("update function cannot be nil")
	}
	ud.notifyDrawUpdate = update
	return nil

}

func (ud *UMLDiagram) updateDrawData() duerror.DUError {
	gs := make([]drawdata.Gadget, 0, len(ud.componentsSelected))
	// TODO
	// as := make([]drawdata.Association, 0, len(ud.componentsSelected))
	for _, c := range ud.componentsContainer.GetAll() {
		cDrawData := c.GetDrawData()
		if cDrawData == nil {
			continue
		}
		switch c.(type) {
		case *component.Gadget:
			gs = append(gs, cDrawData.(drawdata.Gadget))
		case *component.Association:
			continue
		}
	}
	ud.drawData.Gadgets = gs
	// ud.drawData.Associations = as
	if ud.notifyDrawUpdate == nil {
		return nil
	}
	return ud.notifyDrawUpdate()
}
