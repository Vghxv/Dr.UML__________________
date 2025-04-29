package umldiagram

import (
	"time"

	"Dr.uml/backend/utils"
	"github.com/google/uuid"
)

type DiagramType string

func NewDiagramType(dt string) DiagramType {
	switch dt {
	case "ClassDiagram", "UseCaseDiagram", "SequenceDiagram":
		return DiagramType(dt)
	}
	panic("invalid diagramType")
}

// Diagram represents a UML diagram
type UMLDiagram struct {
	id           uuid.UUID
	name         string
	diagramType  DiagramType // e.g., "Class", "UseCase", "Sequence"
	lastModified time.Time
	startPoint   utils.Point // for dragging and linking ass
	color        utils.Color
}

// NewUMLDiagram creates a new UMLDiagram instance
func NewUMLDiagram(name string, dt DiagramType) *UMLDiagram {
	id := uuid.New()
	return &UMLDiagram{
		id:           id,
		name:         name,
		diagramType:  dt,
		lastModified: time.Now(),
	}
}

func (ud *UMLDiagram) GetId() uuid.UUID {
	return ud.id
}
func (ud *UMLDiagram) GetName() string {
	return ud.name
}

func NewUMLDiagramWithPath(path string) (*UMLDiagram, error) {
	/* TODO */
	return nil, nil
}

func (ud *UMLDiagram) AddGadget(gadgetType string) error {
	// Add a gadget to the diagram
	/* TODO */
	return nil
}
