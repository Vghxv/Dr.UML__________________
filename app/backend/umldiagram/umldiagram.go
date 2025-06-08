package umldiagram

import (
	"Dr.uml/backend/component/attribute"
	"fmt"
	"slices"
	"time"

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
	lastModified    time.Time
	startPoint      utils.Point // for dragging and linking ass
	backgroundColor string

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
		lastModified:        time.Now(),
		startPoint:          utils.Point{X: 0, Y: 0},
		backgroundColor:     drawdata.DefaultDiagramColor, // Default white background
		componentsContainer: components.NewContainerMap(),
		associations:        make(map[*component.Gadget][2][]*component.Association),
		componentsSelected:  make(map[component.Component]bool),
		drawData: drawdata.Diagram{
			Margin:    drawdata.Margin,
			LineWidth: drawdata.LineWidth,
			Color:     drawdata.DefaultDiagramColor,
		},
	}, nil
}

func LoadExistUMLDiagram(filename string, file utils.SavedFile) (*UMLDiagram, duerror.DUError) {
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
	return ud.lastModified
}

// Setters
func (ud *UMLDiagram) SetPointComponent(point utils.Point) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.SetPoint(point)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

func (ud *UMLDiagram) SetLayerComponent(layer int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.SetLayer(layer)
	case *component.Association:
		return g.SetLayer(layer)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

func (ud *UMLDiagram) SetColorComponent(colorHexStr string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.SetColor(colorHexStr)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

func (ud *UMLDiagram) SetAttrContentComponent(section int, index int, content string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch c := c.(type) {
	case *component.Gadget:
		return c.SetAttrContent(section, index, content)
	case *component.Association:
		return c.SetAttrContent(index, content)
	default:
		return duerror.NewInvalidArgumentError("invalid selected component")
	}
}

func (ud *UMLDiagram) SetAttrSizeComponent(section int, index int, size int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch c := c.(type) {
	case *component.Gadget:
		return c.SetAttrSize(section, index, size)
	case *component.Association:
		return c.SetAttrSize(index, size)
	default:
		return duerror.NewInvalidArgumentError("invalid selected component")
	}
}

func (ud *UMLDiagram) SetAttrStyleComponent(section int, index int, style int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch c := c.(type) {
	case *component.Gadget:
		return c.SetAttrStyle(section, index, style)
	case *component.Association:
		return c.SetAttrStyle(index, style)
	default:
		return duerror.NewInvalidArgumentError("invalid selected component")
	}
}

func (ud *UMLDiagram) SetAttrFontComponent(section int, index int, fontFile string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}
	switch c := c.(type) {
	case *component.Gadget:
		return c.SetAttrFontFile(section, index, fontFile)
	case *component.Association:
		return c.SetAttrFontFile(index, fontFile)
	default:
		return duerror.NewInvalidArgumentError("invalid selected component")
	}
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

	switch a := c.(type) {
	case *component.Association:
		stNew, err := ud.componentsContainer.SearchGadget(point)
		if err != nil {
			return err
		}
		if stNew == nil {
			return duerror.NewInvalidArgumentError("point does not contain a gadget")
		}

		// update association
		stOld := a.GetParentStart()
		if err = a.SetParentStart(stNew, point); err != nil {
			return err
		}

		// update ud.associations map
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

	default:
		return duerror.NewInvalidArgumentError("selected component is not an association")
	}
}

func (ud *UMLDiagram) SetParentEndComponent(point utils.Point) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch a := c.(type) {
	case *component.Association:
		enNew, err := ud.componentsContainer.SearchGadget(point)
		if err != nil {
			return err
		}
		if enNew == nil {
			return duerror.NewInvalidArgumentError("point does not contain a gadget")
		}

		// update association
		enOld := a.GetParentEnd()
		if err = a.SetParentEnd(enNew, point); err != nil {
			return err
		}

		// update ud.associations map
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

	default:
		return duerror.NewInvalidArgumentError("selected component is not an association")
	}
}

// Methods
func (ud *UMLDiagram) AddGadget(gadgetType component.GadgetType, point utils.Point, layer int, colorHexStr string, header string) duerror.DUError {
	g, err := component.NewGadget(gadgetType, point, layer, colorHexStr, header)
	if err != nil {
		return err
	}
	if err = g.RegisterUpdateParentDraw(ud.updateDrawData); err != nil {
		return err
	}
	if err = ud.componentsContainer.Insert(g); err != nil {
		return err
	}
	ud.associations[g] = [2][]*component.Association{{}, {}}
	return ud.updateDrawData()
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
	if err = ud.componentsContainer.Insert(a); err != nil {
		return err
	}

	// record it, cant modify the slice, being a value of the map, directly
	tmp := ud.associations[stGad]
	tmp[0] = append(tmp[0], a)
	ud.associations[stGad] = tmp

	tmp = ud.associations[enGad]
	tmp[1] = append(tmp[1], a)
	ud.associations[enGad] = tmp

	return ud.updateDrawData()
}

func (ud *UMLDiagram) RemoveSelectedComponents() duerror.DUError {
	for c := range ud.componentsSelected {
		switch c := c.(type) {
		case *component.Gadget:
			if err := ud.removeGadget(c); err != nil {
				return err
			}
		case *component.Association:
			if err := ud.removeAssociation(c); err != nil {
				return err
			}
		}
	}
	return ud.updateDrawData()
}

func (ud *UMLDiagram) SelectComponent(point utils.Point) duerror.DUError {
	c, err := ud.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	if c == nil {
		ud.UnselectAllComponents()
		return ud.updateDrawData()
	}
	if _, ok := ud.componentsSelected[c]; !ok {
		switch c := c.(type) {
		case *component.Gadget:
			err := c.SetIsSelected(true)
			if err != nil {
				return err
			}
			ud.componentsSelected[c] = true
		case *component.Association:
			err := c.SetIsSelected(true)
			if err != nil {
				return err
			}
			ud.componentsSelected[c] = true
		}
	}
	// ud.componentsSelected[c] = true
	return ud.updateDrawData()
}

func (ud *UMLDiagram) UnselectComponent(point utils.Point) duerror.DUError {
	c, err := ud.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	if c == nil {
		return nil
	}
	delete(ud.componentsSelected, c)
	return ud.updateDrawData()
}

func (ud *UMLDiagram) UnselectAllComponents() duerror.DUError {
	for c := range ud.componentsSelected {
		c.SetIsSelected(false)
	}
	// Clear the selected components map
	ud.componentsSelected = make(map[component.Component]bool)
	return nil
}

func (ud *UMLDiagram) AddAttributeToGadget(section int, content string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.AddAttribute(section, content)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

func (ud *UMLDiagram) RemoveAttributeFromGadget(section int, index int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch g := c.(type) {
	case *component.Gadget:
		return g.RemoveAttribute(section, index)
	default:
		return duerror.NewInvalidArgumentError("selected component is not a gadget")
	}
}

func (ud *UMLDiagram) AddAttributeToAssociation(ratio float64, content string) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch ass := c.(type) {
	case *component.Association:
		return ass.AddAttribute(ratio, content)
	default:
		return duerror.NewInvalidArgumentError("selected component is not an association")
	}
}

func (ud *UMLDiagram) RemoveAttributeFromAssociation(index int) duerror.DUError {
	c, err := ud.getSelectedComponent()
	if err != nil {
		return err
	}

	switch ass := c.(type) {
	case *component.Association:
		return ass.RemoveAttribute(index)
	default:
		return duerror.NewInvalidArgumentError("selected component is not an association")
	}
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

func (ud *UMLDiagram) removeGadget(gad *component.Gadget) duerror.DUError {
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
	return ud.componentsContainer.Remove(gad)
}

func (ud *UMLDiagram) removeAssociation(a *component.Association) duerror.DUError {
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
	return ud.componentsContainer.Remove(a)
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
		newAtt, err := attribute.FromSavedAssAttributes(savedAtt)
		if err != nil {
			return err, index
		}
		if err = ass.AddAttribute(newAtt); err != nil {
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
func (ud *UMLDiagram) collectGadgets(res *utils.SavedFile) (map[*component.Gadget]int, duerror.DUError) {
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

func (ud *UMLDiagram) collectAssociations(dp map[*component.Gadget]int, res *utils.SavedFile) duerror.DUError {
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

func (ud *UMLDiagram) SaveToFile(filename string) (*utils.SavedFile, duerror.DUError) {
	if filename != ud.name {
		ud.name = filename
	}

	res := &utils.SavedFile{
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

func (ud *UMLDiagram) IfUnsavedChangesExist() bool {
	return ud.lastModified.After(ud.lastSave)
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
	ud.lastModified = time.Now()
	return ud.updateParentDraw()
}
