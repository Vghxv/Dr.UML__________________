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

func NewContainerMap() componentsContainer {
	return &containerMap{compMap: make(map[component.Component]bool)}
}

func (cp *containerMap) Insert(c component.Component) duerror.DUError {
	cp.compMap[c] = true
	return nil
}

func (cp *containerMap) Remove(c component.Component) duerror.DUError {
	_, ok := cp.compMap[c]
	if ok {
		delete(cp.compMap, c)
	}
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
		candidateLayer, err := candidate.GetLayer()
		if err != nil {
			return nil, err
		}
		cLayer, err := c.GetLayer()
		if  err != nil {
			return nil, err
		}
		if cLayer > candidateLayer {
			candidate = c
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
