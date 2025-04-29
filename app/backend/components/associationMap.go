package components

import (
	"slices"

	"Dr.uml/backend/component"
	"Dr.uml/backend/utils/duerror"
)

// implement AssociationGraph interface
type associationMap struct {
	assMap map[*component.Gadget](map[*component.Gadget][]*component.Association)
}

func NewAssociationMap() AssociationGraph {
	return &associationMap{
		assMap: make(map[*component.Gadget](map[*component.Gadget]([]*component.Association))),
	}
}

func (am *associationMap) FindStartEnd(st *component.Gadget, en *component.Gadget) ([]*component.Association, duerror.DUError) {
	if st==nil || en==nil {
		return nil, duerror.NewInvalidArgumentError("start or end gadget is nil")
	}
	if _, ok := am.assMap[st]; !ok {
		return nil, nil
	}
	return am.assMap[st][en], nil
}

func (am *associationMap) FindStart(st *component.Gadget) ([]*component.Association, duerror.DUError) {
	if st==nil {
		return nil, duerror.NewInvalidArgumentError("start gadget is nil")
	}
	if _, ok := am.assMap[st]; !ok {
		return nil, nil
	}
	as := make([]*component.Association, 0, len(am.assMap[st]))
	for _, v := range am.assMap[st] {
		as = append(as, v...)
	}
	return as, nil
}

func (am *associationMap) FindEnd(en *component.Gadget) ([]*component.Association, duerror.DUError) {
	if en==nil {
		return nil, duerror.NewInvalidArgumentError("end gadget is nil")
	}
	as := make([]*component.Association, 0)
	for _, endMap := range am.assMap {
		if _, ok := endMap[en]; !ok {
			continue
		}
		as = append(as, endMap[en]...)
	}
	return as, nil
}

func (am *associationMap) FindEither(g *component.Gadget) ([]*component.Association, duerror.DUError) {
	if g==nil {
		return nil, duerror.NewInvalidArgumentError("gadget is nil")
	}
	asStart, err := am.FindStart(g)
	if err != nil {
		return nil, err
	}
	asEnd, err := am.FindEnd(g)
	if err != nil {
		return nil, err
	}
	asSet := make(map[*component.Association]bool)
	asResult := make([]*component.Association, 0, len(asStart)+len(asEnd))
	for _, a := range asStart {
		asSet[a] = true
		asResult = append(asResult, a)
	}
	for _, a := range asEnd {
		if _, ok := asSet[a]; ok {
			continue
		}
		asSet[a] = true
		asResult = append(asResult, a)
	}
	return asResult, nil
}

func (am *associationMap) Update(a *component.Association, oldSt *component.Gadget, oldEn *component.Gadget) duerror.DUError {
	if a == nil {
		return duerror.NewInvalidArgumentError("association is nil")
	}
	if oldSt == nil || oldEn == nil {
		return duerror.NewInvalidArgumentError("old start or end gadget is nil")
	}
	if _, ok := am.assMap[oldSt]; !ok {
		return duerror.NewInvalidArgumentError("old start is not in association map")
	}
	if _, ok := am.assMap[oldSt][oldEn]; !ok {
		return duerror.NewInvalidArgumentError("old end is not in association map")
	}
	idx := slices.IndexFunc(am.assMap[oldSt][oldEn], func(aa *component.Association) bool { return aa == a })
	if idx == -1 {
		return duerror.NewInvalidArgumentError("association is not in association map")
	}
	am.assMap[oldSt][oldEn] = slices.Delete(am.assMap[oldSt][oldEn], idx, idx+1)
	return am.Insert(a)
}

func (am *associationMap) Insert(a *component.Association) duerror.DUError {
	if a == nil {
		return duerror.NewInvalidArgumentError("association is nil")
	}
	start, _ := a.GetParentStart()
	end, _ := a.GetParentEnd()
	if _, ok := am.assMap[start]; !ok {
		am.assMap[start] = make(map[*component.Gadget][]*component.Association)
	}
	if _, ok := am.assMap[start][end]; !ok {
		am.assMap[start][end] = []*component.Association{a}
	}else {
		am.assMap[start][end] = append(am.assMap[start][end], a)
	}
	return nil
}

func (am *associationMap) Remove(a *component.Association) duerror.DUError {
	if a == nil {
		return duerror.NewInvalidArgumentError("association is nil")
	}
	start, _ := a.GetParentStart()
	end, _ := a.GetParentEnd()
	if _, ok := am.assMap[start]; !ok {
		return nil
	}
	if _, ok := am.assMap[start][end]; !ok {
		return nil
	}
	idx := slices.IndexFunc(am.assMap[start][end], func(aa *component.Association) bool { return aa == a })
	if idx == -1 {
		return nil
	}
	am.assMap[start][end] = slices.Delete(am.assMap[start][end], idx, idx+1)
	if len(am.assMap[start][end]) == 0 {
		delete(am.assMap[start], end)
		if len(am.assMap[start]) == 0 {
			delete(am.assMap, start)
		}
	}
	return nil
}

// return a list of associations that are connected to gadget
func (am *associationMap) RemoveGadget(g *component.Gadget) ([]*component.Association, duerror.DUError) {
	if g == nil {
		return nil, duerror.NewInvalidArgumentError("gadget is nil")
	}
	as, err := am.FindEither(g)
	if err != nil {
		return nil, err
	}
	delete(am.assMap, g)
	for start, endMap := range am.assMap {
		if _, ok := endMap[g]; !ok {
			continue
		}
		delete(endMap, g)
		if len(endMap) == 0 {
			delete(am.assMap, start)
		}
	}
	return as, nil
}
