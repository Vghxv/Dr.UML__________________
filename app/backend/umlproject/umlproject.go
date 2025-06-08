package umlproject

import (
	"Dr.uml/backend/component"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
	"context"
	"encoding/json"
	"fmt"
	"github.com/titanous/json5"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"maps"
	"os"
	"slices"
	"time"
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
func (p *UMLProject) SetPointComponent(point utils.Point) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetPointComponent(point); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetLayerComponent(layer int) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetLayerComponent(layer); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetColorComponent(colorHexStr string) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetColorComponent(colorHexStr); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetAttrContentComponent(section int, index int, content string) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetAttrContentComponent(section, index, content); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetAttrSizeComponent(section int, index int, size int) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetAttrSizeComponent(section, index, size); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetAttrStyleComponent(section int, index int, style int) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetAttrStyleComponent(section, index, style); err != nil {
		return err
	}
	p.lastModified = time.Now()
	return nil
}

func (p *UMLProject) SetAttrFontComponent(section int, index int, font string) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	if err := p.currentDiagram.SetAttrFontComponent(section, index, font); err != nil {
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
		err := p.OpenDiagram(diagramName)
		if err != nil {
			return err
		}
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

func (p *UMLProject) OpenDiagram(filename string) duerror.DUError {
	err := utils.ValidateFilePath(filename)
	if err != nil {
		return err
	}

	if p.availableDiagrams[filename] {
		return p.SelectDiagram(filename)
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return duerror.NewFileIOError(fmt.Sprintf("Failed to open file %s.\n Error: %s", filename, err.Error()))
	}
	defer file.Close()

	decoder := json5.NewDecoder(file)
	var savedFileData utils.SavedFile

	if err := decoder.Decode(&savedFileData); err != nil {
		return duerror.NewInvalidArgumentError(fmt.Sprintf("Failed to decode file %s.\n Error: %s", filename, err.Error()))
	}
	if savedFileData.Filetype&utils.SupportedFiletypes == 0 {
		return duerror.NewInvalidArgumentError(fmt.Sprintf("Unsupported file type %d in file %s", savedFileData.Filetype, filename))
	}
	savedFileData.Filetype >>= 1 // Remove the first bit, which is used to indicate if the file is a diagram or submodule
	switch savedFileData.Filetype {
	case utils.FiletypeDiagram:
		dia, err := umldiagram.LoadExistUMLDiagram(filename, savedFileData)
		if err != nil {
			return err
		}
		p.availableDiagrams[filename] = true
		p.activeDiagrams[filename] = dia
		p.lastModified = time.Now()
		p.currentDiagram = dia

		break
	case utils.FiletypeSubmodule:

		// TODO
		break
	default:
		return duerror.NewInvalidArgumentError(fmt.Sprintf("Unknown filetype %d", savedFileData.Filetype))
	}

	return nil
}

func (p *UMLProject) SaveDiagram(filename string) duerror.DUError {
	if p.currentDiagram == nil {
		return duerror.NewInvalidArgumentError("No current diagram selected")
	}
	originalFilename := p.currentDiagram.GetName()
	savedFileData, err := p.currentDiagram.SaveToFile(filename)
	if err != nil {
		return duerror.NewParsingError(fmt.Sprintf("Failed to export diagram %s.\n Error: %s", filename, err.Error()))
	}
	if originalFilename != filename {
		delete(p.availableDiagrams, originalFilename)
		delete(p.activeDiagrams, originalFilename)
		p.availableDiagrams[filename] = true
		p.activeDiagrams[filename] = p.currentDiagram
	}

	file, err := os.Create(filename)
	if err != nil {
		return duerror.NewFileIOError(fmt.Sprintf("Failed to create file %s.\n Error: %s", p.name, err.Error()))
	}

	data, err := json.MarshalIndent(savedFileData, "", "  ")
	if err != nil {
		return duerror.NewFileIOError(fmt.Sprintf("Failed to marshal data to JSON for file %s.\n Error: %s", p.name, err.Error()))
	}
	if _, err := file.Write(data); err != nil {
		return duerror.NewFileIOError(fmt.Sprintf("Failed to write data to file %s.\n Error: %s", p.name, err.Error()))
	}
	if err := file.Close(); err != nil {
		return duerror.NewFileIOError(fmt.Sprintf("Failed to close file %s.\n Error: %s", p.name, err.Error()))
	}

	p.lastModified = time.Now()
	return nil
}
