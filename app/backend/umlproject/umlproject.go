package umlproject

import (
	"time"

	"Dr.uml/backend/component"
	"Dr.uml/backend/component/drawdata"
	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type UMLProject struct {
	name           string
	lastModified   time.Time
	currentDiagram *umldiagram.UMLDiagram            // The currently selected diagram
	diagrams       map[string]*umldiagram.UMLDiagram // Use a map to store diagrams, keyed by their ID
	openedDiagrams map[string]*umldiagram.UMLDiagram // Keep track of opened diagrams
	activeDiagrams map[string]*umldiagram.UMLDiagram // Keep track of active diagrams
}

// NewUMLProject creates a new UMLProject instance
func NewUMLProject(name string) *UMLProject {

	return &UMLProject{
		name:           name,
		lastModified:   time.Now(),
		diagrams:       make(map[string]*umldiagram.UMLDiagram),
		openedDiagrams: make(map[string]*umldiagram.UMLDiagram),
		activeDiagrams: make(map[string]*umldiagram.UMLDiagram),
	}
}

// GetName returns the name of the UMLProject
func (p *UMLProject) GetName() string {
	return p.name
}

// GetLastModified returns the last modified time of the UMLProject
func (p *UMLProject) OpenProject() ([]*umldiagram.UMLDiagram, []string, duerror.DUError) {
	activeDiagrams := make([]*umldiagram.UMLDiagram, 0, len(p.activeDiagrams))
	for _, diagram := range p.openedDiagrams {
		d, err := umldiagram.NewUMLDiagram(diagram.GetName(), diagram.GetDiagramType())
		if err != nil {
			return nil, nil, duerror.NewInvalidArgumentError("Failed to create diagram")
		}
		p.activeDiagrams[d.GetName()] = d
	}

	return activeDiagrams, p.GetAvailableDiagrams(), nil
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
func (p *UMLProject) SelectDiagram(diagramName string) duerror.DUError {
	if diagram, ok := p.diagrams[diagramName]; ok {
		p.currentDiagram = diagram
		return nil
	}
	return duerror.NewInvalidArgumentError("Diagram not found")
}

// AddGadget
func (p *UMLProject) AddGadget(
	gadgetType component.GadgetType,
	point utils.Point,
) (drawdata.Gadget, duerror.DUError) {

	dd, err := p.currentDiagram.AddGadget(gadgetType, point)
	if err != nil {
		return drawdata.Gadget{}, err
	}
	p.lastModified = time.Now()
	if dd.GadgetType != 1 {
		p.activeDiagrams[p.currentDiagram.GetName()] = p.currentDiagram
	}
	return dd, nil

}

// Add diagram
func (p *UMLProject) AddNewDiagram(
	diagramType umldiagram.DiagramType,
	name string,
) duerror.DUError {
	for _, diagram := range p.diagrams {
		if diagram.GetName() == name {
			return duerror.NewInvalidArgumentError("Diagram name already exists")
		}
	}

	diagram, err := umldiagram.NewUMLDiagram(name, diagramType)
	if err != nil {
		return err
	}

	p.diagrams[name] = diagram
	p.currentDiagram = diagram
	p.openedDiagrams[name] = diagram
	p.lastModified = time.Now()
	return nil

}

// CreateDiagram(path) creates a new instance of the diagram and load the diagram info at path
func (p *UMLProject) createDiagram(path string) duerror.DUError {
	diagram, err := umldiagram.NewUMLDiagramWithPath(path)
	if err != nil {
		return err
	}
	p.diagrams[diagram.GetName()] = diagram
	p.currentDiagram = diagram
	p.openedDiagrams[diagram.GetName()] = diagram
	p.lastModified = time.Now()
	return nil
}
