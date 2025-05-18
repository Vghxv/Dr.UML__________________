package components

import (
	"maps"
	"slices"

	"Dr.uml/backend/component"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

// implement ComponentsContainer using map
type containerMap struct {
	compMap map[component.Component]bool
}

func NewContainerMap() Container {
	return &containerMap{compMap: make(map[component.Component]bool)}
}

func (cp *containerMap) Insert(c component.Component) duerror.DUError {
	if c == nil {
		return duerror.NewInvalidArgumentError("component is nil")
	}
	cp.compMap[c] = true
	return nil
}

func (cp *containerMap) Remove(c component.Component) duerror.DUError {
	delete(cp.compMap, c)
	return nil
}

func (cp *containerMap) Search(p utils.Point) (component.Component, duerror.DUError) {
	var candidate component.Component
	for c := range cp.compMap {
		cover, err := c.Cover(p)
		if err != nil {
			return nil, err
		}
		if !cover {
			continue
		}
		if candidate == nil {
			candidate = c
			continue
		}
		if c.GetLayer() > candidate.GetLayer() {
			candidate = c
		}
	}
	return candidate, nil
}

func (cp *containerMap) SearchGadget(p utils.Point) (*component.Gadget, duerror.DUError) {
	var candidate *component.Gadget
	for c := range cp.compMap {
		switch c := c.(type) {
		case *component.Gadget:
			cover, err := c.Cover(p)
			if err != nil {
				return nil, err
			}
			if !cover {
				continue
			}
			if candidate == nil {
				candidate = c
				continue
			}
			if c.GetLayer() < candidate.GetLayer() {
				candidate = c
			}
		default:
			continue
		}
	}
	return candidate, nil
}

func (cp *containerMap) GetAll() []component.Component {
	return slices.Collect(maps.Keys(cp.compMap))
}

func (cp *containerMap) Len() (int, duerror.DUError) {
	return len(cp.compMap), nil
}
