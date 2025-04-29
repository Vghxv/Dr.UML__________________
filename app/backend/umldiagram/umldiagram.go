package umldiagram

import (
	"time"

	"github.com/google/uuid"
)

// Diagram represents a UML diagram
type UMLDiagram struct {
	id          uuid.UUID
	name        string
	diagramType string // e.g., "Class", "UseCase", "Sequence"
	lastOpened  time.Time
	// Add other relevant diagram properties here
}

// NewUMLDiagram creates a new UMLDiagram instance
func NewUMLDiagram(name, diagramType string) *UMLDiagram {
	id := uuid.New()
	return &UMLDiagram{
		id:          id,
		name:        name,
		diagramType: diagramType,
		lastOpened:  time.Now(),
	}
}

func NewUMLDiagramWithPath(path string) (*UMLDiagram, error) {
	// read from disk
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
	// This is a placeholder implementation
	return nil
}
