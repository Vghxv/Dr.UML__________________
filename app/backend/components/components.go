package components

import (
	"Dr.uml/backend/component"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type Components struct {
	componentsContainer componentsContainer
	selectedComponents  map[component.Component]bool
	drawData            drawdata.Components
	updateParentDraw    func() duerror.DUError
}

func NewComponents() *Components {
	return &Components{
		componentsContainer: NewContainerMap(),
		selectedComponents:  make(map[component.Component]bool),
		drawData: drawdata.Components{
			Margin:    drawdata.Margin,
			LineWidth: drawdata.LineWidth,
		},
	}
}

func (cs *Components) SelectComponent(point utils.Point) duerror.DUError {
	comp, err := cs.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	if comp == nil {
		return nil
	}
	cs.selectedComponents[comp] = true
	return nil
}

func (cs *Components) UnselectComponent(point utils.Point) duerror.DUError {
	comp, err := cs.componentsContainer.Search(point)
	if err != nil {
		return err
	}
	if comp == nil {
		return nil
	}
	delete(cs.selectedComponents, comp)
	return nil
}

func (cs *Components) UnselectAllComponents() duerror.DUError {
	for comp := range cs.selectedComponents {
		delete(cs.selectedComponents, comp)
	}
	return nil
}

func (cs *Components) GetDrawData() (drawdata.Components, duerror.DUError) {
	return cs.drawData, nil
}

func (cs *Components) updateDrawData() duerror.DUError {
	gs := make([]drawdata.Gadget, 0, len(cs.selectedComponents))
	// TODO
	// as := make([]drawdata.Association, 0, len(cs.selectedComponents))
	for _, c := range cs.componentsContainer.GetAll() {
		cDrawData, err := c.GetDrawData()
		if err != nil {
			return err
		}
		if cDrawData == nil {
			continue
		}
		switch c.(type) {
		case *component.Gadget:
			gs = append(gs, cDrawData.(drawdata.Gadget))
		case *component.Association:
			continue
		}
	}
	cs.drawData.Gadgets = gs
	// cs.drawData.Associations = as
	if cs.updateParentDraw == nil {
		return nil
	}
	return cs.updateParentDraw()
}

func (cs *Components) AddGadget(gadgetType component.GadgetType, point utils.Point) duerror.DUError {
	gadget, err := component.NewGadget(gadgetType, point)
	if err != nil {
		return err
	}
	err = gadget.RegisterUpdateParentDraw(cs.updateDrawData)
	if err != nil {
		return err
	}
	err = cs.componentsContainer.Insert(gadget)
	if err != nil {
		return err
	}
	err = cs.updateDrawData()
	if err != nil {
		return err
	}
	return nil
}

func (cs *Components) RegisterUpdateParentDraw(update func() duerror.DUError) duerror.DUError {
	cs.updateParentDraw = update
	return nil
}
