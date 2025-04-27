package Components

import (
	"Dr.uml/backend/Component"
)

type componentPool struct {
	list []*Component.Component
}

func (cp *componentPool) insert(c *Component.Component) {
	cp.list = append(cp.list, c)
}

func (cp *componentPool) remove(c *Component.Component) {
	for i, comp := range cp.list {
		if comp == c {
			cp.list = append(cp.list[:i], cp.list[i+1:]...)
			break
		}
	}
}

func (cp *componentPool) Len() int {
	return len(cp.list)
}

