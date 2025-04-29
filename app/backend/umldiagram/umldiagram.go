package umldiagram

import (
	"time"

	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
	"github.com/google/uuid"
)

type DiagramType int

const (
	ClassDiagram    = 1 << iota // 0x01
	UseCaseDiagram  = 1 << iota // 0x02
	SequenceDiagram = 1 << iota // 0x04
	supportedType   = ClassDiagram | UseCaseDiagram | SequenceDiagram
)

func check(input DiagramType) bool {
	return input&supportedType == input && input != 0
}

// Diagram represents a UML diagram
type UMLDiagram struct {
	id           uuid.UUID
	name         string
	diagramType  DiagramType // e.g., "Class", "UseCase", "Sequence"
	lastModified time.Time
	startPoint   utils.Point // for dragging and linking ass
	/* TODO */
	// add background color

}

// NewUMLDiagram creates a new UMLDiagram instance
func NewUMLDiagram(name string, dt DiagramType) (*UMLDiagram, duerror.DUError) {
	id := uuid.New()

	if !utils.IsValidFilePath(name) {
		return nil, duerror.NewInvalidArgumentError("Invalid diagram name")
	}

	if !check(dt) {
		return nil, duerror.NewInvalidArgumentError("Invalid diagram type")
	}

	return &UMLDiagram{
		id:           id,
		name:         name,
		diagramType:  dt,
		lastModified: time.Now(),
		startPoint:   utils.Point{X: 0, Y: 0},
	}, nil
}

func (ud *UMLDiagram) GetId() uuid.UUID {
	return ud.id
}

func (ud *UMLDiagram) GetName() string {
	return ud.name
}

func NewUMLDiagramWithPath(path string) (*UMLDiagram, error) {
	if !utils.IsValidFilePath(path) {
		return nil, duerror.NewInvalidArgumentError("Invalid diagram name")
	}
	return &UMLDiagram{
		id:           uuid.New(),
		name:         path,
		diagramType:  ClassDiagram,
		lastModified: time.Now(),
		startPoint:   utils.Point{X: 0, Y: 0},
	}, nil
}

func (ud *UMLDiagram) AddGadget(gadgetType string) error {
	// Add a gadget to the diagram
	/* TODO */
	return nil
}
