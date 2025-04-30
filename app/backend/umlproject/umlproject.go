package umlproject

import (
	"context"
	"fmt"
	"time"

	"Dr.uml/backend/component"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type UMLProject struct {
	ctx              context.Context
	name             string
	lastModified     time.Time
	currentDiagram   *umldiagram.UMLDiagram            // The currently selected diagram
	diagrams         map[string]*umldiagram.UMLDiagram // Use a map to store diagrams, keyed by their ID
	openedDiagrams   map[string]*umldiagram.UMLDiagram // Keep track of opened diagrams
	activeDiagrams   map[string]*umldiagram.UMLDiagram // Keep track of active diagrams
	notifyDrawUpdate func(string) duerror.DUError
	// notifyDrawUpdate func() duerror.DUError TODO
}

func (p *UMLProject) Startup(ctx context.Context) {
	p.ctx = ctx
	p.AddNewDiagram(umldiagram.ClassDiagram, "new class diagram")
	p.SelectDiagram("new class diagram")
}

// func (p *UMLProject) ProcessWithCallback(callbackID string) {
// 	// Simulate some processing
// 	// result := number * 2

// 	// Call the JavaScript callback function with the result
// 	runtime.EventsEmit(p.ctx, callbackID, 654)
// }

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

// OpenProject returns the active diagrams and available diagrams in the project
func (p *UMLProject) OpenProject() ([]string, []string, duerror.DUError) {
	return p.GetLastOpenedDiagrams(), p.GetAvailableDiagrams(), nil
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
) duerror.DUError {

	err := p.currentDiagram.AddGadget(gadgetType, point)
	if err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil

}

func (p *UMLProject) StartAddAssociation(point utils.Point) duerror.DUError {
	return p.currentDiagram.StartAddAssociation(point)
}

func (p *UMLProject) EndAddAssociation(associationType component.AssociationType, point utils.Point) duerror.DUError {
	return p.currentDiagram.EndAddAssociation(associationType, [2]utils.Point{point, point})
}

// Add diagram
func (p *UMLProject) AddNewDiagram(
	diagramType umldiagram.DiagramType,
	name string,
) duerror.DUError {
	if _, exists := p.diagrams[name]; exists {
		return duerror.NewInvalidArgumentError("Diagram name already exists")
	}

	diagram, err := umldiagram.NewUMLDiagram(name, diagramType)
	if err != nil {
		return err
	}
	diagram.RegisterNotifyDrawUpdate(p.InvalidateCanvas)

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

func (p *UMLProject) InvalidateCanvas() duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	// p.notifyDrawUpdate(p.currentDiagram.GetName())
	dd, err := p.currentDiagram.GetDrawData()
	if err != nil {
		return err
	}
	fmt.Println("InvalidateCanvas", dd)
	runtime.EventsEmit(p.ctx, "backend-event", dd)

	return nil
}

// GetUserData returns a struct with user information
func (a *UMLProject) GetUserData() map[string]interface{} {
	return map[string]interface{}{
		"id":       1,
		"username": "wailsuser",
		"email":    "user@example.com",
		"roles":    []string{"admin", "user"},
		"active":   true,
	}
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

func (p *UMLProject) GetCurrentDiagram() *umldiagram.UMLDiagram {
	if p.currentDiagram == nil {
		return nil
	}
	return p.currentDiagram
}

func (p *UMLProject) GetCurrentDiagramName() string {
	if p.currentDiagram == nil {
		return ""
	}
	return p.currentDiagram.GetName()
}

func (p *UMLProject) ProcessWithCallback(callbackID string) duerror.DUError {
	// Call the JavaScript callback function with the result

	runtime.EventsEmit(p.ctx, callbackID, 123)
	return nil
}
