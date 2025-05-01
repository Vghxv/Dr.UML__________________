package umldiagram

import (
	"time"

	"Dr.uml/backend/component"
	"Dr.uml/backend/components"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

// type Weekday string

// const (
//     Sunday    Weekday = "Sunday"
//     Monday    Weekday = "Monday"
//     Tuesday   Weekday = "Tuesday"
//     Wednesday Weekday = "Wednesday"
//     Thursday  Weekday = "Thursday"
//     Friday    Weekday = "Friday"
//     Saturday  Weekday = "Saturday"
// )

// var AllWeekdays = []struct {
//     Value  Weekday
//     TSName string
// }{
//     {Sunday, "SUNDAY"},
//     {Monday, "MONDAY"},
//     {Tuesday, "TUESDAY"},
//     {Wednesday, "WEDNESDAY"},
//     {Thursday, "THURSDAY"},
//     {Friday, "FRIDAY"},
//     {Saturday, "SATURDAY"},
// }

type DiagramType int

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

func isValidDiagramType(input DiagramType) bool {
	return input&supportedType == input && input != 0
}

// Diagram represents a UML diagram
type UMLDiagram struct {
	name             string
	diagramType      DiagramType // e.g., "Class", "UseCase", "Sequence"
	lastModified     time.Time
	startPoint       utils.Point // for dragging and linking ass
	backgroundColor  utils.Color
	components       *components.Components
	drawData         drawdata.Diagram
	notifyDrawUpdate func() duerror.DUError
}

// NewUMLDiagram creates a new UMLDiagram instance
func NewUMLDiagram(name string, dt DiagramType) (*UMLDiagram, duerror.DUError) {

	if !utils.IsValidFilePath(name) {
		return nil, duerror.NewInvalidArgumentError("Invalid diagram name")
	}

	if !isValidDiagramType(dt) {
		return nil, duerror.NewInvalidArgumentError("Invalid diagram type")
	}

	dg := UMLDiagram{
		name:            name,
		diagramType:     dt,
		lastModified:    time.Now(),
		startPoint:      utils.Point{X: 0, Y: 0},
		backgroundColor: utils.FromHex(drawdata.DefaultDiagramColor), // Default white background
		components:      components.NewComponents(),
		drawData:        drawdata.Diagram{Color: drawdata.DefaultDiagramColor},
	}
	dg.components.RegisterUpdateParentDraw(dg.updateDrawData)
	return &dg, nil
}

func (ud *UMLDiagram) StartAddAssociation(point utils.Point) duerror.DUError {
	ud.startPoint = point
	return nil
}

func (ud *UMLDiagram) GetName() string {
	return ud.name
}

func (ud *UMLDiagram) GetDiagramType() DiagramType {
	return ud.diagramType
}

func NewUMLDiagramWithPath(path string) (*UMLDiagram, error) {
	if !utils.IsValidFilePath(path) {
		return nil, duerror.NewInvalidArgumentError("Invalid diagram name")
	}
	return &UMLDiagram{
		name:         path,
		diagramType:  ClassDiagram,
		lastModified: time.Now(),
		startPoint:   utils.Point{X: 0, Y: 0},
	}, nil
}

func (ud *UMLDiagram) AddGadget(gadgetType component.GadgetType, point utils.Point) duerror.DUError {
	err := ud.components.AddGadget(gadgetType, point)
	if err != nil {
		return err
	}
	ud.updateDrawData()
	return nil

}

func (ud *UMLDiagram) EndAddAssociation(assType component.AssociationType, point [2]utils.Point) duerror.DUError {
	err := ud.components.AddAssociation(assType, point)
	if err != nil {
		return err
	}
	return ud.updateDrawData()
}

func (ud *UMLDiagram) GetDrawData() (drawdata.Diagram, duerror.DUError) {
	return ud.drawData, nil
}

func (ud *UMLDiagram) RegisterNotifyDrawUpdate(update func() duerror.DUError) duerror.DUError {
	ud.notifyDrawUpdate = update
	return nil

}

func (ud *UMLDiagram) updateDrawData() duerror.DUError {
	csdd, err := ud.components.GetDrawData()
	if err != nil {
		return err
	}
	if ud.notifyDrawUpdate == nil {
		return nil
	}
	ud.drawData.Color = ud.backgroundColor.ToHex()
	ud.drawData.Components = csdd
	return ud.notifyDrawUpdate()
}
