package umlproject

import (
	"context"
	"maps"
	"slices"
	"time"

	"Dr.uml/backend/component"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type UMLProject struct {
	ctx               context.Context
	name              string
	lastModified      time.Time
	currentDiagram    *umldiagram.UMLDiagram            // The currently selected diagram
	availableDiagrams map[string]bool                   // Use a map to store diagrams, keyed by their ID
	activeDiagrams    map[string]*umldiagram.UMLDiagram // Keep track of active diagrams
	runFrontend       bool
}

// Constructor
func CreateEmptyUMLProject(fileName string) (*UMLProject, duerror.DUError) {
	// TODO: also check the file is exist or not
	if err := utils.ValidateFilePath(fileName); err != nil {
		return nil, err
	}
	return &UMLProject{
		name:              fileName,
		lastModified:      time.Now(),
		availableDiagrams: make(map[string]bool),
		activeDiagrams:    make(map[string]*umldiagram.UMLDiagram),
	}, nil
}

func LoadExistUMLProject(fileName string) (*UMLProject, duerror.DUError) {
	// TODO
	return nil, nil
}

// Getter
func (p *UMLProject) GetName() string {
	return p.name
}

func (p *UMLProject) GetLastModified() time.Time {
	return p.lastModified
}

func (p *UMLProject) GetCurrentDiagramName() string {
	if p.currentDiagram == nil {
		return ""
	}
	return p.currentDiagram.GetName()
}

func (p *UMLProject) GetAvailableDiagramsNames() []string {
	return slices.Collect(maps.Keys(p.availableDiagrams))
}

func (p *UMLProject) GetActiveDiagramsNames() []string {
	activeNames := make([]string, 0, len(p.activeDiagrams))
	for _, d := range p.activeDiagrams {
		activeNames = append(activeNames, d.GetName())
	}
	return activeNames
}

// Setter
func (p *UMLProject) SetPointComp(point utils.Point) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetPointComp(point); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetSetLayerComp(layer int) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetSetLayerComp(layer); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetColorComp(colorHexStr string) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetColorComp(colorHexStr); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetAttrContentComp(section int, index int, content string) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetAttrContentComp(section, index, content); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetAttrSizeComp(section int, index int, size int) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetAttrSizeComp(section, index, size); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetAttrStyleComp(section int, index int, style int) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetAttrStyleComp(section, index, style); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

// methods
func (p *UMLProject) Startup(ctx context.Context) {
	p.ctx = ctx
	// TODO: Remove this bcz can't handle error here
	p.runFrontend = true
	p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "new class diagram")
	p.SelectDiagram("new class diagram")
}

func (p *UMLProject) SelectDiagram(diagramName string) duerror.DUError {
	if _, ok := p.availableDiagrams[diagramName]; !ok {
		return duerror.NewInvalidArgumentError("Diagram not found")
	}
	if _, ok := p.activeDiagrams[diagramName]; !ok {
		diagram, err := umldiagram.LoadExistUMLDiagram(diagramName)
		if err != nil {
			return err
		}
		p.activeDiagrams[diagramName] = diagram
	}
	p.currentDiagram = p.activeDiagrams[diagramName]
	// TODO: when multiple diagrams exists, unregister the old one

	return p.currentDiagram.RegisterUpdateParentDraw(p.InvalidateCanvas)
}

func (p *UMLProject) CreateEmptyUMLDiagram(diagramType umldiagram.DiagramType, diagramName string) duerror.DUError {
	if _, ok := p.availableDiagrams[diagramName]; ok {
		return duerror.NewInvalidArgumentError("Diagram name already exists")
	}
	d, err := umldiagram.CreateEmptyUMLDiagram(diagramName, diagramType)
	if err != nil {
		return err
	}
	p.availableDiagrams[diagramName] = true
	p.activeDiagrams[diagramName] = d
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) CloseDiagram(diagramName string) duerror.DUError {
	// TODO: save file?
	if _, ok := p.activeDiagrams[diagramName]; !ok {
		return duerror.NewInvalidArgumentError("Diagram not loaded")
	}
	if p.currentDiagram != nil && p.currentDiagram.GetName() == diagramName {
		p.currentDiagram = nil
	}
	delete(p.activeDiagrams, diagramName)
	return nil
}

func (p *UMLProject) DeleteDiagram(diagramName string) duerror.DUError {
	// TODO: remove the file
	return nil
}

func (p *UMLProject) AddGadget(gadgetType component.GadgetType, point utils.Point, layer int, colorHexStr string, header string) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.AddGadget(gadgetType, point, layer, colorHexStr, header); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) StartAddAssociation(point utils.Point) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	return p.currentDiagram.StartAddAssociation(point)
}

func (p *UMLProject) EndAddAssociation(associationType component.AssociationType, point utils.Point) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.EndAddAssociation(associationType, point); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) RemoveSelectedComponents() duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.RemoveSelectedComponents(); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) AddAttributeToGadget(section int, content string) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.AddAttributeToGadget(section, content); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) RemoveAttributeFromGadget(section int, index int) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.RemoveAttributeFromGadget(section, index); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SelectComponent(point utils.Point) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SelectComponent(point); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

// draw
func (p *UMLProject) GetDrawData() drawdata.Diagram {
	if p.currentDiagram == nil {
		return drawdata.Diagram{}
	}
	return p.currentDiagram.GetDrawData()
}

func (p *UMLProject) InvalidateCanvas() duerror.DUError {
	if !p.runFrontend {
		return nil
	}

	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	runtime.EventsEmit(p.ctx, "backend-event", p.GetDrawData())
	return nil
}
