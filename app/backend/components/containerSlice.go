package components

import (
	"Dr.uml/backend/component"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror/duerror"
)

// implement ComponentsContainer using slice
type containerSlice struct {
	list []component.Component
}

func NewContainerSlice() componentsContainer{
	return &containerSlice{list: make([]component.Component, 0)}
}

func (cp *containerSlice) Insert(c component.Component) duerror.DUError {
	for _, existedComp := range cp.list {
		if existedComp == c {
			return duerror.NewInvalidArgumentError("Component already exists in pool")
		}
	}
	cp.list = append(cp.list, c)
	return nil
}

func (cp *containerSlice) Remove(c component.Component) duerror.DUError{
	for i, comp := range cp.list {
		if comp == c {
			cp.list = append(cp.list[:i], cp.list[i+1:]...)
			return nil
		}
	}
	return duerror.NewInvalidArgumentError("Component not found in pool")
}

func (cp *containerSlice) Search(p utils.Point) (component.Component, duerror.DUError) {
	var candidate component.Component
	for _, c := range cp.list {
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
		if err != nil {
			return nil, err
		}
		if cLayer > candidateLayer {
			candidate = c
		}
	}
	return candidate, nil
}

func (cp *containerSlice) Len() (int, duerror.DUError) {
	return len(cp.list), nil
}