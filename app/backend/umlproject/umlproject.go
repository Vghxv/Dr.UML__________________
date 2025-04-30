package umlproject

import (
	"time"

	"github.com/google/uuid"

	"Dr.uml/backend/component"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/utils/duerror"
)

type UMLProject struct {
	name           string
	lastModified   time.Time
	currentDiagram *umldiagram.UMLDiagram               // The currently selected diagram
	diagrams       map[uuid.UUID]*umldiagram.UMLDiagram // Use a map to store diagrams, keyed by their ID
	openedDiagrams map[uuid.UUID]*umldiagram.UMLDiagram // Keep track of opened diagrams
}

// NewUMLProject creates a new UMLProject instance
func NewUMLProject(name string) *UMLProject {

	return &UMLProject{
		name:           name,
		lastModified:   time.Now(),
		diagrams:       make(map[uuid.UUID]*umldiagram.UMLDiagram),
		openedDiagrams: make(map[uuid.UUID]*umldiagram.UMLDiagram),
	}
}

// GetName returns the name of the UMLProject
func (p *UMLProject) GetName() string {
	return p.name
}

// GetLastModified returns the last modified time of the UMLProject
func (p *UMLProject) OpenProject() ([]*umldiagram.UMLDiagram, []string, []uuid.UUID) {
	openedDiagrams := make([]*umldiagram.UMLDiagram, 0, len(p.openedDiagrams))
	uuidList := make([]uuid.UUID, 0, len(p.openedDiagrams))
	for _, diagram := range p.openedDiagrams {
		openedDiagrams = append(openedDiagrams, diagram)
		uuidList = append(uuidList, diagram.GetId())
	}

	diagramList := make([]string, 0, len(p.diagrams))
	for _, diagram := range p.diagrams {
		diagramList = append(diagramList, diagram.GetName())
	}

	return openedDiagrams, diagramList, uuidList
}

// GetAvailableDiagrams returns a list of the names of all available diagrams in the project
func (p *UMLProject) GetAvailableDiagrams() []string {
	diagramList := make([]string, 0, len(p.diagrams))
	for _, diagram := range p.diagrams {
		diagramList = append(diagramList, diagram.GetName())
	}
	return diagramList
}

// GetLastOpenedDiagrams returns a list of the names of the last opened diagrams
func (p *UMLProject) GetLastOpenedDiagrams() []string {
	openedDiagramList := make([]string, 0, len(p.openedDiagrams))
	for _, diagram := range p.openedDiagrams {
		openedDiagramList = append(openedDiagramList, diagram.GetName())
	}
	return openedDiagramList
}

// SelectDiagram sets the current diagram to the one with the given ID
func (p *UMLProject) SelectDiagram(diagramID uuid.UUID) duerror.DUError {
	if diagram, ok := p.diagrams[diagramID]; ok {
		p.currentDiagram = diagram
		return nil
	}
	return duerror.NewInvalidArgumentError("Diagram not found")
}

// AddGadget
func (p *UMLProject) AddGadget(
	gadgetType component.GadgetType,
	diagramID uuid.UUID,
) duerror.DUError {
	if diagram, ok := p.diagrams[diagramID]; ok {
		// Add the gadget to the diagram
		diagram.AddGadget(gadgetType)
		p.lastModified = time.Now()
		return nil
	}
	return duerror.NewInvalidArgumentError("Diagram not found")
}

// Add diagram
func (p *UMLProject) AddNewDiagram(
	diagramType umldiagram.DiagramType,
	name string,
) duerror.DUError {
	id := uuid.New()
	diagram, err := umldiagram.NewUMLDiagram(name, diagramType)
	if err != nil {
		return err
	}
	p.diagrams[id] = diagram
	p.currentDiagram = diagram
	p.openedDiagrams[id] = diagram
	p.lastModified = time.Now()
	return nil

}

// CreateDiagram(path) creates a new instance of the diagram and load the diagram info at path
func (p *UMLProject) createDiagram(path string) duerror.DUError {
	diagram, err := umldiagram.NewUMLDiagramWithPath(path)
	if err != nil {
		return err
	}
	p.diagrams[diagram.GetId()] = diagram
	p.currentDiagram = diagram
	p.openedDiagrams[diagram.GetId()] = diagram
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) DrawDigram() drawdata.Diagram {
	if p.currentDiagram == nil {
		return drawdata.Diagram{}
	}
	data, err := p.currentDiagram.GetDrawData()
	if err != nil {
		return drawdata.Diagram{}
	}
	return data
}
