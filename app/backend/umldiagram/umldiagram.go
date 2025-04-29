package umldiagram

import (
	"time"

	"github.com/google/uuid"
)

type DiagramType string

const (
	ClassDiagram DiagramType = "ClassDiagram"
)

// Diagram represents a UML diagram
type UMLDiagram struct {
	id          uuid.UUID
	name        string
	diagramType DiagramType // e.g., "Class", "UseCase", "Sequence"
	lastOpened  time.Time
	// Add other relevant diagram properties here
}

// NewUMLDiagram creates a new UMLDiagram instance
func NewUMLDiagram(name string, dt DiagramType) *UMLDiagram {
	id := uuid.New()
	return &UMLDiagram{
		id:          id,
		name:        name,
		diagramType: dt,
		lastOpened:  time.Now(),
	}
}

func NewUMLDiagramWithPath(path string) (*UMLDiagram, error) {
    /* TODO */
	return nil, nil
}

func (ud *UMLDiagram) GetId() uuid.UUID {
	return ud.id
}
func (ud *UMLDiagram) GetName() string {
	return ud.name
}

func (ud *UMLDiagram) AddGadget(gadgetType string) error {
	// Add a gadget to the diagram
/* TODO */
	return nil
}
