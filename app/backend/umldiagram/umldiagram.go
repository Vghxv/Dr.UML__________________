package umldiagram

import (
	"fmt"
	"slices"
	"time"

	"Dr.uml/backend/command"
	"Dr.uml/backend/component/attribute"

	"Dr.uml/backend/component"
	"Dr.uml/backend/components"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type DiagramType int

const (
	ClassDiagram = 1 << iota // 0x01
	UseCaseDiagram
	SequenceDiagram
	supportedType = ClassDiagram
)

var AllDiagramTypes = []struct {
	Value  DiagramType
	TSName string
}{
	{ClassDiagram, "ClassDiagram"},
}

// Other methods
func validateDiagramType(input DiagramType) duerror.DUError {
	if !(input&supportedType == input && input != 0) {
		return duerror.NewInvalidArgumentError("Invalid diagram type")
	}
	return nil
}

type UMLDiagram struct {
	name            string
	diagramType     DiagramType // e.g., "Class", "UseCase", "Sequence"
	startPoint      utils.Point // for dragging and linking ass
	backgroundColor string

	cmdManager          *command.Manager
	componentsContainer components.Container
	componentsSelected  map[component.Component]bool
	associations        map[*component.Gadget][2][]*component.Association

	lastSave time.Time // for saving and loading

	updateParentDraw func() duerror.DUError
	drawData         drawdata.Diagram
}

// Constructor
func CreateEmptyUMLDiagram(name string, dt DiagramType) (*UMLDiagram, duerror.DUError) {
	// TODO: also check the file is exist or not
	if err := utils.ValidateFilePath(name); err != nil {
		return nil, err
	}
	if err := validateDiagramType(dt); err != nil {
		return nil, err
	}
	return &UMLDiagram{
		name:                name,
		diagramType:         dt,
		startPoint:          utils.Point{X: 0, Y: 0},
		backgroundColor:     drawdata.DefaultDiagramColor, // Default white background
		cmdManager:          command.NewManager(time.Now()),
		componentsContainer: components.NewContainerMap(),
		componentsSelected:  make(map[component.Component]bool),
		associations:        make(map[*component.Gadget][2][]*component.Association),
		drawData: drawdata.Diagram{
			Margin:    drawdata.Margin,
			LineWidth: drawdata.LineWidth,
			Color:     drawdata.DefaultDiagramColor,
		},
	}, nil
}

func LoadExistUMLDiagram(filename string, file utils.SavedDiagram) (*UMLDiagram, duerror.DUError) {
	dia, err := CreateEmptyUMLDiagram(filename, DiagramType(file.Filetype)) // Shift right to remove the filetype bit
	if err != nil {
		return nil, err
	}

	dp, err := dia.loadGadgets(file.Gadgets)
	if err != nil {
		return nil, duerror.NewCorruptedFile(fmt.Sprintf(err.Error()+"from %s", filename))
	}

	if err = dia.loadAsses(file.Associations, dp); err != nil {
		return nil, err
	}

	if err = dia.updateDrawData(); err != nil {
		return nil, err
	}

	dia.name = filename

	return dia, nil
}

// Getters
func (ud *UMLDiagram) GetName() string {
	return ud.name
}

func (ud *UMLDiagram) GetDiagramType() DiagramType {
	return ud.diagramType
}

func (ud *UMLDiagram) GetLastModified() time.Time {
	return ud.cmdManager.GetLastModified()
}

// Setters
func (ud *UMLDiagram) SetPointComponent(point utils.Point) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	g, ok := c.(*component.Gadget)
	if !ok {
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
	cmd := &moveGadgetCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		gadget:   g,
		newPoint: point,
		oldPoint: g.GetPoint(),
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) SetLayerComponent(layer int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	oldLayer := c.GetLayer()
	cmd := &setterCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		component: c,
		execute:   func() duerror.DUError { return c.SetLayer(layer) },
		unexecute: func() duerror.DUError { return c.SetLayer(oldLayer) },
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil

}

func (ud *UMLDiagram) SetColorComponent(colorHexStr string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	g, ok := c.(*component.Gadget)
	if !ok {
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
	oldColor := g.GetColor()
	cmd := &setterCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		component: g,
		execute:   func() duerror.DUError { return g.SetColor(colorHexStr) },
		unexecute: func() duerror.DUError { return g.SetColor(oldColor) },
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil

}

func (ud *UMLDiagram) SetAttrContentComponent(section int, index int, content string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	var setContent func(content string) duerror.DUError
	var oldContent string
	switch c := c.(type) {
	case *component.Gadget:
		att, err := c.GetAttribute(section, index)
		if err != nil {
			return nil
		}
		oldContent = att.GetContent()
		setContent = func(content string) duerror.DUError {
			return c.SetAttrContent(section, index, content)
		}
	case *component.Association:
		att, err := c.GetAttribute(index)
		if err != nil {
			return nil
		}
		oldContent = att.GetContent()
		setContent = func(content string) duerror.DUError {
			return c.SetAttrContent(index, content)
		}
	default:
		return duerror.NewInvalidArgumentError("invalid selected component")
	}

	cmd := &setterCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		component: c,
		execute:   func() duerror.DUError { return setContent(content) },
		unexecute: func() duerror.DUError { return setContent(oldContent) },
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) SetAttrSizeComponent(section int, index int, size int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	var setSize func(size int) duerror.DUError
	var oldSize int
	switch c := c.(type) {
	case *component.Gadget:
		att, err := c.GetAttribute(section, index)
		if err != nil {
			return nil
		}
		oldSize = att.GetSize()
		setSize = func(size int) duerror.DUError {
			return c.SetAttrSize(section, index, size)
		}
	case *component.Association:
		att, err := c.GetAttribute(index)
		if err != nil {
			return nil
		}
		oldSize = att.GetSize()
		setSize = func(size int) duerror.DUError {
			return c.SetAttrSize(index, size)
		}
	default:
		return duerror.NewInvalidArgumentError("invalid selected component")
	}

	cmd := &setterCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		component: c,
		execute:   func() duerror.DUError { return setSize(size) },
		unexecute: func() duerror.DUError { return setSize(oldSize) },
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) SetAttrStyleComponent(section int, index int, style int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	var setStyle func(style int) duerror.DUError
	var oldStyle int
	switch c := c.(type) {
	case *component.Gadget:
		att, err := c.GetAttribute(section, index)
		if err != nil {
			return nil
		}
		oldStyle = int(att.GetStyle())
		setStyle = func(style int) duerror.DUError {
			return c.SetAttrStyle(section, index, style)
		}
	case *component.Association:
		att, err := c.GetAttribute(index)
		if err != nil {
			return nil
		}
		oldStyle = int(att.GetStyle())
		setStyle = func(style int) duerror.DUError {
			return c.SetAttrStyle(index, style)
		}
	default:
		return duerror.NewInvalidArgumentError("invalid selected component")
	}

	cmd := &setterCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		component: c,
		execute:   func() duerror.DUError { return setStyle(style) },
		unexecute: func() duerror.DUError { return setStyle(oldStyle) },
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) SetAttrFontComponent(section int, index int, fontFile string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	var setFont func(font string) duerror.DUError
	var oldFontFile string
	switch c := c.(type) {
	case *component.Gadget:
		att, err := c.GetAttribute(section, index)
		if err != nil {
			return nil
		}
		oldFontFile = att.GetFontFile()
		setFont = func(font string) duerror.DUError {
			return c.SetAttrFontFile(section, index, font)
		}
	case *component.Association:
		att, err := c.GetAttribute(index)
		if err != nil {
			return nil
		}
		oldFontFile = att.GetFontFile()
		setFont = func(font string) duerror.DUError {
			return c.SetAttrFontFile(index, font)
		}
	default:
		return duerror.NewInvalidArgumentError("invalid selected component")
	}

	cmd := &setterCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		component: c,
		execute:   func() duerror.DUError { return setFont(fontFile) },
		unexecute: func() duerror.DUError { return setFont(oldFontFile) },
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) SetAttrRatioComponent(section int, index int, ratio float64) duerror.DUError {
	// section arg is not used, just to keep similar signature
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	switch c := c.(type) {
	case *component.Association:
		return c.SetAttrRatio(index, ratio)
	default:
		return duerror.NewInvalidArgumentError("selected component is not an association")
	}
}

func (ud *UMLDiagram) SetParentStartComponent(point utils.Point) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	a, ok := c.(*component.Association)
	if !ok {
		return duerror.NewInvalidArgumentError("selected component is not an association")
	}

	c, err = ud.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	stNew, ok := c.(*component.Gadget)
	if !ok {
		return duerror.NewInvalidArgumentError("component at point is not a gadget")
	}
	stRatioNew, err := component.CalAssociationPointRatio(stNew, point)
	if err != nil {
		return err
	}

	cmd := &setParentStartCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		association: a,
		stNew:       stNew,
		stOld:       a.GetParentStart(),
		stRatioNew:  stRatioNew,
		stRatioOld:  a.GetStartRatio(),
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) SetParentEndComponent(point utils.Point) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	a, ok := c.(*component.Association)
	if !ok {
		return duerror.NewInvalidArgumentError("selected component is not an association")
	}

	c, err = ud.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	enNew, ok := c.(*component.Gadget)
	if !ok {
		return duerror.NewInvalidArgumentError("component at point is not a gadget")
	}
	enRatioNew, err := component.CalAssociationPointRatio(enNew, point)
	if err != nil {
		return err
	}

	cmd := &setParentEndCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		association: a,
		enNew:       enNew,
		enOld:       a.GetParentEnd(),
		enRatioNew:  enRatioNew,
		enRatioOld:  a.GetEndRatio(),
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) SetAssociationType(value component.AssociationType) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	a, ok := c.(*component.Association)
	if !ok {
		return duerror.NewInvalidArgumentError("selected component is not an association")
	}
	oldAssType := a.GetAssType() // Capture the original value before change
	cmd := &setterCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		component: c,
		execute:   func() duerror.DUError { return a.SetAssType(value) },
		unexecute: func() duerror.DUError { return a.SetAssType(oldAssType) },
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

// Methods
func (ud *UMLDiagram) Undo() duerror.DUError {
	if err := ud.cmdManager.Undo(); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) Redo() duerror.DUError {
	if err := ud.cmdManager.Redo(); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) AddGadget(gadgetType component.GadgetType, point utils.Point, layer int, colorHexStr string, header string) duerror.DUError {
	g, err := component.NewGadget(gadgetType, point, layer, colorHexStr, header)
	if err != nil {
		return err
	}
	if err = g.RegisterUpdateParentDraw(ud.updateDrawData); err != nil {
		return err
	}
	cmd := &addComponentCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		component: g,
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) StartAddAssociation(point utils.Point) duerror.DUError {
	if err := ud.validatePoint(point); err != nil {
		return err
	}
	ud.startPoint = point
	return nil
}

func (ud *UMLDiagram) EndAddAssociation(assType component.AssociationType, endPoint utils.Point) duerror.DUError {
	stPoint := ud.startPoint
	ud.startPoint = utils.Point{X: 0, Y: 0}
	if err := ud.validatePoint(endPoint); err != nil {
		return err
	}

	// search parents
	stGad, err := ud.componentsContainer.SearchGadget(stPoint)
	if err != nil {
		return err
	}
	if stGad == nil {
		return duerror.NewInvalidArgumentError("start point does not contain a gadget")
	}
	enGad, err := ud.componentsContainer.SearchGadget(endPoint)
	if err != nil {
		return err
	}
	if enGad == nil {
		return duerror.NewInvalidArgumentError("end point does not contain a gadget")
	}

	// create association
	parents := [2]*component.Gadget{stGad, enGad}
	a, err := component.NewAssociation(parents, assType, stPoint, endPoint)
	if err != nil {
		return err
	}
	if err = a.RegisterUpdateParentDraw(ud.updateDrawData); err != nil {
		return err
	}

	cmd := &addComponentCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		component: a,
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) RemoveSelectedComponents() duerror.DUError {
	comps := map[component.Component]bool{}
	for c := range ud.componentsSelected {
		comps[c] = true
		switch g := c.(type) {
		case *component.Gadget:
			for a := range ud.getAllAssociationsInGadget(g) {
				comps[a] = true
			}
		}
	}
	cmd := &removeSelectedComponentCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		components: comps,
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) RemoveComponentAtPoint(point utils.Point) duerror.DUError {
	// Find the component at the specified point
	c, err := ud.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	if c == nil {
		return duerror.NewInvalidArgumentError("No component found at the specified point")
	}

	// Create a map of components to remove (similar to RemoveSelectedComponents)
	comps := map[component.Component]bool{c: true}

	// If it's a gadget, also include all associated associations
	switch g := c.(type) {
	case *component.Gadget:
		for a := range ud.getAllAssociationsInGadget(g) {
			comps[a] = true
		}
	}

	// Execute the removal using the command pattern
	cmd := &removeSelectedComponentCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		components: comps,
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) SelectComponent(point utils.Point) duerror.DUError {
	c, err := ud.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	if c != nil && c.GetIsSelected() {
		// if comp is already selected, do nothing
		return nil
	}

	var affectedComps map[component.Component]bool
	var newValue bool
	if c == nil {
		// if click on nothing, unselect all
		affectedComps = ud.componentsSelected
		newValue = false
	} else {
		// if click on comp, select it
		affectedComps = map[component.Component]bool{c: true}
		newValue = true
	}
	cmd := &selectAllCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		components: affectedComps,
		newValue:   newValue,
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) AddAttributeToGadget(section int, content string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	g, ok := c.(*component.Gadget)
	if !ok {
		duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
	if section < 0 || section >= len(g.GetAttributesLen()) {
		duerror.NewInvalidArgumentError("invalid section")
	}

	cmd := &addAttributeGadgetCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		gadget:  g,
		content: content,
		section: section,
		index:   g.GetAttributesLen()[section],
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) RemoveAttributeFromGadget(section int, index int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	g, ok := c.(*component.Gadget)
	if !ok {
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
	att, err := g.GetAttribute(section, index)
	if err != nil {
		return err
	}

	cmd := &removeAttributeGadgetCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		gadget:  g,
		content: att.GetContent(),
		section: section,
		index:   index,
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) AddAttributeToAssociation(ratio float64, content string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	a, ok := c.(*component.Association)
	if !ok {
		duerror.NewInvalidArgumentError("selected component is not an association")
	}

	cmd := &addAttributeAssociationCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		association: a,
		content:     content,
		ratio:       ratio,
		index:       a.GetAttributesLen(),
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) RemoveAttributeFromAssociation(index int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	a, ok := c.(*component.Association)
	if !ok {
		duerror.NewInvalidArgumentError("selected component is not an association")
	}
	att, err := a.GetAttribute(index)
	if err != nil {
		return err
	}

	cmd := &removeAttributeAssociationCommand{
		baseCommand: baseCommand{
			diagram: ud,
			before:  ud.GetLastModified(),
			after:   time.Now(),
		},
		association: a,
		content:     att.GetContent(),
		ratio:       att.GetRatio(),
		index:       index,
	}
	if err := ud.cmdManager.Execute(cmd); err != nil {
		return err
	}
	return nil
}

// Private methods
func (ud *UMLDiagram) getSelectedComponent() (component.Component, duerror.DUError) {
	if len(ud.componentsSelected) != 1 {
		return nil, duerror.NewInvalidArgumentError("can only operate on one component")
	}
	for c := range ud.componentsSelected {
		return c, nil
	}
	return nil, duerror.NewInvalidArgumentError("no component selected")
}

func (ud *UMLDiagram) validatePoint(point utils.Point) duerror.DUError {
	if point.X < 0 || point.Y < 0 {
		return duerror.NewInvalidArgumentError("point coordinates must be non-negative")
	}
	return nil
}

func (ud *UMLDiagram) loadGadgetAttributes(gadget *component.Gadget, attributes []utils.SavedAtt) (duerror.DUError, int) {
	const SectionBound = 0.3 // A gadget has 3 sections: header, attributes, methods. Saving sections in SavedAtt.Ratio
	if gadget == nil {
		return duerror.NewInvalidArgumentError("Cannot load attributes to a nil gadget"), 0
	}
	for index, savedAtt := range attributes {
		newAtt, err := attribute.FromSavedAttribute(savedAtt)
		if err != nil {
			return err, index
		}

		if err = gadget.AddBuiltAttribute(int(savedAtt.Ratio/SectionBound), newAtt); err != nil {
			return err, index
		}
	}
	return nil, 0
}

func (ud *UMLDiagram) loadGadgets(gadgets []utils.SavedGad) (map[int]*component.Gadget, duerror.DUError) {
	dp := make(map[int]*component.Gadget)

	// Load Gadgets
	for index, savedGadget := range gadgets {
		gadget, err := component.FromSavedGadget(savedGadget)
		if err != nil {
			return nil, err
		}

		if err, errIndex := ud.loadGadgetAttributes(gadget, savedGadget.Attributes); err != nil {
			return nil, duerror.NewCorruptedFile(fmt.Sprintf(
				"Error on parsing %d-th attribute of %d-th gadget.  Detail: %s",
				errIndex, index, err.Error()),
			)
		}

		// Code below suffers from #89. If anything is being added, put 'em above.
		// Stuff done by UD.AddGadget.
		// Replace these shitty code with it when it won't cause a series of useless overhead.
		if err = gadget.RegisterUpdateParentDraw(ud.updateDrawData); err != nil {
			return nil, err
		}
		if err = ud.componentsContainer.Insert(gadget); err != nil {
			return nil, err
		}
		ud.associations[gadget] = [2][]*component.Association{{}, {}}

		dp[index] = gadget
	}

	return dp, nil
}

func (ud *UMLDiagram) loadAssAttributes(ass *component.Association, attributes []utils.SavedAtt) (duerror.DUError, int) {
	if ass == nil {
		return duerror.NewInvalidArgumentError("UR loading attributes to a nil ass"), 0
	}
	for index, savedAtt := range attributes {
		newAtt, err := attribute.FromSavedAssAttribute(savedAtt)
		if err != nil {
			return err, index
		}
		if err = ass.AddLoadedAttribute(newAtt); err != nil {
			return err, index
		}
	}
	return nil, 0
}

func (ud *UMLDiagram) loadAsses(asses []utils.SavedAss, dp map[int]*component.Gadget) duerror.DUError {
	for index, ass := range asses {
		parents := [2]*component.Gadget{dp[ass.Parents[0]], dp[ass.Parents[1]]}
		newAss, err := component.FromSavedAssociation(ass, parents)
		if err != nil {
			return duerror.NewCorruptedFile(fmt.Sprintf("Error on creating %d-th association: %s", index, err.Error()))
		}
		if err = ud.componentsContainer.Insert(newAss); err != nil {
			return err
		}
		if err, attIndex := ud.loadAssAttributes(newAss, ass.Attributes); err != nil {
			return duerror.NewCorruptedFile(fmt.Sprintf("Error on adding %d-th attribute for %d-th association: %s", attIndex, index, err.Error()))
		}

		// I'm a thief
		tmp := ud.associations[parents[0]]
		tmp[0] = append(tmp[0], newAss)
		ud.associations[parents[0]] = tmp

		tmp = ud.associations[parents[1]]
		tmp[1] = append(tmp[1], newAss)
		ud.associations[parents[1]] = tmp

		if err = newAss.RegisterUpdateParentDraw(ud.updateDrawData); err != nil {
			return err
		}
	}

	return nil
}
func (ud *UMLDiagram) collectGadgets(res *utils.SavedDiagram) (map[*component.Gadget]int, duerror.DUError) {
	dp := make(map[*component.Gadget]int, ud.componentsContainer.Len())
	cnt := 0
	for _, comp := range ud.componentsContainer.GetAll() {
		switch comp.(type) {
		case *component.Gadget:
			if _, ok := dp[comp.(*component.Gadget)]; !ok {
				dp[comp.(*component.Gadget)] = cnt
				res.Gadgets = append(res.Gadgets, comp.(*component.Gadget).ToSavedGadget())
				cnt++
			}
		default:
			continue
		}
	}
	return dp, nil
}

func (ud *UMLDiagram) collectAssociations(dp map[*component.Gadget]int, res *utils.SavedDiagram) duerror.DUError {
	for comp, index := range dp {
		for _, ass := range ud.associations[comp][0] {
			milkBuyer := ass.GetParentEnd()
			milkBuyerIndex, ok := dp[milkBuyer]
			if !ok {
				return duerror.NewParsingError("SecondParent not found")
			}
			res.Associations = append(res.Associations, ass.ToSavedAssociation(
				[2]int{
					index, milkBuyerIndex,
				}))
		}
	}
	return nil
}

func (ud *UMLDiagram) SaveToFile(filename string) (*utils.SavedDiagram, duerror.DUError) {
	if filename != ud.name {
		ud.name = filename
	}

	res := &utils.SavedDiagram{
		Filetype:     utils.FiletypeDiagram | int(ud.diagramType)<<1,
		LastEdit:     "",
		Gadgets:      nil,
		Associations: nil,
	}

	dp, err := ud.collectGadgets(res)
	if err != nil {
		return nil, err
	}

	if err := ud.collectAssociations(dp, res); err != nil {
		return nil, err
	}
	ud.lastSave = time.Now()
	res.LastEdit = ud.lastSave.Format(time.RFC3339)

	return res, nil
}

func (ud *UMLDiagram) HasUnsavedChanges() bool {
	return ud.GetLastModified().After(ud.lastSave)
}

// draw
func (ud *UMLDiagram) GetDrawData() drawdata.Diagram {
	return ud.drawData
}

func (ud *UMLDiagram) RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError {
	if update == nil {
		return duerror.NewInvalidArgumentError("update function cannot be nil")
	}
	ud.updateParentDraw = update
	return nil

}

func (ud *UMLDiagram) updateDrawData() duerror.DUError {
	gs := make([]drawdata.Gadget, 0, len(ud.componentsSelected))
	as := make([]drawdata.Association, 0, len(ud.componentsSelected))
	for _, c := range ud.componentsContainer.GetAll() {
		cDrawData := c.GetDrawData()
		if cDrawData == nil {
			continue
		}
		switch c.(type) {
		case *component.Gadget:
			gs = append(gs, cDrawData.(drawdata.Gadget))
		case *component.Association:
			as = append(as, cDrawData.(drawdata.Association))
		}
	}
	ud.drawData.Gadgets = gs
	ud.drawData.Associations = as
	if ud.updateParentDraw == nil {
		return nil
	}
	return ud.updateParentDraw()
}

func (ud *UMLDiagram) addComponent(c component.Component) duerror.DUError {
	switch c := c.(type) {
	case *component.Gadget:
		return ud.addGadget(c)
	case *component.Association:
		return ud.addAssociation(c)
	default:
		return duerror.NewInvalidArgumentError("unsupported component type")
	}
}

func (ud *UMLDiagram) addGadget(g *component.Gadget) duerror.DUError {
	if err := ud.componentsContainer.Insert(g); err != nil {
		return err
	}
	if _, ok := ud.associations[g]; !ok {
		ud.associations[g] = [2][]*component.Association{{}, {}}
	}
	if g.GetIsSelected() {
		ud.componentsSelected[g] = true
	}
	if err := ud.updateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) addAssociation(a *component.Association) duerror.DUError {
	if err := ud.componentsContainer.Insert(a); err != nil {
		return err
	}
	// record it, cant modify the slice, being a value of the map, directly
	stGad := a.GetParentStart()
	enGad := a.GetParentEnd()
	if _, ok := ud.associations[stGad]; !ok {
		ud.associations[stGad] = [2][]*component.Association{{}, {}}
	}
	if _, ok := ud.associations[enGad]; !ok {
		ud.associations[enGad] = [2][]*component.Association{{}, {}}
	}

	tmp := ud.associations[stGad]
	tmp[0] = append(tmp[0], a)
	ud.associations[stGad] = tmp

	tmp = ud.associations[enGad]
	tmp[1] = append(tmp[1], a)
	ud.associations[enGad] = tmp

	if a.GetIsSelected() {
		ud.componentsSelected[a] = true
	}
	if err := ud.updateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) removeComponent(c component.Component) duerror.DUError {
	switch c := c.(type) {
	case *component.Gadget:
		return ud.removeGadget(c)
	case *component.Association:
		return ud.removeAssociation(c)
	default:
		return duerror.NewInvalidArgumentError("unsupported component type")
	}
}
func (ud *UMLDiagram) removeGadget(gad *component.Gadget) duerror.DUError {
	// Remove all associations related to this gadget and remove the gadget itself
	if _, ok := ud.associations[gad]; ok {
		for _, a := range ud.associations[gad][0] {
			if err := ud.removeAssociation(a); err != nil {
				return err
			}
		}
		for _, a := range ud.associations[gad][1] {
			if err := ud.removeAssociation(a); err != nil {
				return err
			}
		}
		delete(ud.associations, gad)
	}
	delete(ud.componentsSelected, gad)
	if err := ud.componentsContainer.Remove(gad); err != nil {
		return err
	}
	if err := ud.updateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) removeAssociation(a *component.Association) duerror.DUError {
	// Unregister the association as an observer of its parent gadgets
	if err := a.UnregisterAsObserver(); err != nil {
		// Log error but continue with removal
		// In a production system, you might want to log this error
	}

	st := a.GetParentStart()
	en := a.GetParentEnd()
	if _, ok := ud.associations[st]; ok {
		stList := ud.associations[st][0]
		index := slices.Index(stList, a)
		if index >= 0 {
			stList = slices.Delete(stList, index, index+1)
		}
		ud.associations[st] = [2][]*component.Association{stList, ud.associations[st][1]}
	}
	if _, ok := ud.associations[en]; ok {
		enList := ud.associations[en][1]
		index := slices.Index(enList, a)
		if index >= 0 {
			enList = slices.Delete(enList, index, index+1)
		}
		ud.associations[en] = [2][]*component.Association{ud.associations[en][0], enList}
	}
	delete(ud.componentsSelected, a)
	if err := ud.componentsContainer.Remove(a); err != nil {
		return err
	}
	if err := ud.updateDrawData(); err != nil {
		return err
	}
	return nil
}

func (ud *UMLDiagram) getAllAssociationsInGadget(g *component.Gadget) map[*component.Association]bool {
	asses := map[*component.Association]bool{}
	if g == nil {
		return asses
	}
	if _, ok := ud.associations[g]; ok {
		for _, a := range ud.associations[g][0] {
			asses[a] = true
		}
		for _, a := range ud.associations[g][1] {
			asses[a] = true
		}
	}
	return asses
}

func (ud *UMLDiagram) removeComponents(comps map[component.Component]bool) duerror.DUError {
	for c := range comps {
		if err := ud.removeComponent(c); err != nil {
			return err
		}
	}
	// updateDD in the removeComponent
	return nil
}

func (ud *UMLDiagram) addComponents(comps map[component.Component]bool) duerror.DUError {
	for c := range comps {
		if err := ud.addComponent(c); err != nil {
			return err
		}
	}
	// updateDD in the addComponent
	return nil
}

func (ud *UMLDiagram) selectAll(comps map[component.Component]bool, newValue bool) duerror.DUError {
	for c := range comps {
		if err := c.SetIsSelected(newValue); err != nil {
			return err
		}
		if newValue {
			ud.componentsSelected[c] = true
		} else {
			delete(ud.componentsSelected, c)
		}
	}
	// updateDD in the SetIsSelected
	return nil
}

func (ud *UMLDiagram) moveGadget(g *component.Gadget, point utils.Point) duerror.DUError {
	if err := g.SetPoint(point); err != nil {
		return err
	}
	if _, ok := ud.associations[g]; ok {
		for _, a := range ud.associations[g][0] {
			if err := a.UpdateDrawData(); err != nil {
				return err
			}
		}
		for _, a := range ud.associations[g][1] {
			if err := a.UpdateDrawData(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ud *UMLDiagram) setParentStartAssociation(a *component.Association, stNew *component.Gadget, stRatio [2]float64) duerror.DUError {
	stOld := a.GetParentStart()
	if err := a.SetParentStart(stNew, stRatio); err != nil {
		return err
	}
	// update ud.associations
	if _, ok := ud.associations[stOld]; ok {
		list := ud.associations[stOld][0]
		index := slices.Index(list, a)
		if index >= 0 {
			list = slices.Delete(list, index, index+1)
		}
		ud.associations[stOld] = [2][]*component.Association{list, ud.associations[stOld][1]}
	}
	if _, ok := ud.associations[stNew]; ok {
		list := ud.associations[stNew][0]
		list = append(list, a)
		ud.associations[stNew] = [2][]*component.Association{list, ud.associations[stNew][1]}
	}
	return nil
}

func (ud *UMLDiagram) setParentEndAssociation(a *component.Association, enNew *component.Gadget, enRatio [2]float64) duerror.DUError {
	enOld := a.GetParentEnd()
	if err := a.SetParentEnd(enNew, enRatio); err != nil {
		return err
	}
	// update ud.associations
	if _, ok := ud.associations[enOld]; ok {
		list := ud.associations[enOld][1]
		index := slices.Index(list, a)
		if index >= 0 {
			list = slices.Delete(list, index, index+1)
		}
		ud.associations[enOld] = [2][]*component.Association{ud.associations[enOld][0], list}
	}
	if _, ok := ud.associations[enNew]; ok {
		list := ud.associations[enNew][1]
		list = append(list, a)
		ud.associations[enNew] = [2][]*component.Association{ud.associations[enNew][0], list}
	}
	return nil
}

func (ud *UMLDiagram) addAttributeGadget(g *component.Gadget, section, index int, content string) duerror.DUError {
	return g.AddAttribute(section, index, content)
}

func (ud *UMLDiagram) removeAttributeGadget(g *component.Gadget, section, index int) duerror.DUError {
	return g.RemoveAttribute(section, index)
}

func (ud *UMLDiagram) addAttributeAssociation(a *component.Association, index int, ratio float64, content string) duerror.DUError {
	return a.AddAttribute(index, ratio, content)
}

func (ud *UMLDiagram) removeAttributeAssociation(a *component.Association, index int) duerror.DUError {
	return a.RemoveAttribute(index)
}
