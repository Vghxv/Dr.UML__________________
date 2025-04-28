package components

import (
	"Dr.uml/backend/component"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type Components struct {
	compoentsContainer ComponentsContainer
	selectedComponents map[component.Component]bool
}

func (cs *Components) SelectComponent(point utils.Point) duerror.DUError {
	comp, err := cs.compoentsContainer.Search(point)
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
	comp, err := cs.compoentsContainer.Search(point)
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
