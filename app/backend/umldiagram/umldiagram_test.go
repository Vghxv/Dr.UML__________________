// VIBE CODING

package umldiagram

import (
	"slices"
	"testing"
	"time"

	"Dr.uml/backend/component"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmptyDiagram(t *testing.T) {
	tests := []struct {
		name        string
		inputName   string
		diagramType DiagramType
		expectError bool
		errorMsg    string
	}{
		{
			name:        "ValidClassDiagram",
			inputName:   "test1.uml",
			diagramType: ClassDiagram,
			expectError: false,
		},
		{
			name:        "ValidUseCaseDiagram",
			inputName:   "test2.uml",
			diagramType: UseCaseDiagram,
			expectError: true,
			errorMsg:    "Invalid diagram type",
		},
		{
			name:        "ValidSequenceDiagram",
			inputName:   "test3.uml",
			diagramType: SequenceDiagram,
			expectError: true,
			errorMsg:    "Invalid diagram type",
		},
		{
			name:        "InvalidDiagramType",
			inputName:   "test4.uml",
			diagramType: DiagramType(8),
			expectError: true,
			errorMsg:    "Invalid diagram type",
		},
		{
			name:        "InvalidDiagramType2",
			inputName:   "test5.uml",
			diagramType: DiagramType(8787),
			expectError: true,
			errorMsg:    "Invalid diagram type",
		},
		{
			name:        "InvalidName",
			inputName:   "",
			diagramType: ClassDiagram,
			expectError: true,
			errorMsg:    "file path is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram, err := CreateEmptyUMLDiagram(tt.inputName, tt.diagramType)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
				assert.Nil(t, diagram)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, diagram)
				assert.Equal(t, tt.inputName, diagram.GetName())
				assert.Equal(t, tt.diagramType, diagram.diagramType)
				assert.WithinDuration(t, time.Now(), diagram.lastModified, time.Second)
				assert.Equal(t, utils.Point{X: 0, Y: 0}, diagram.startPoint)
				// New assertions
				assert.Equal(t, "#FFFFFF", diagram.backgroundColor)
				assert.NotNil(t, diagram.componentsContainer)
			}
		})
	}
}

func TestLoadExistUMLDiagram(t *testing.T) {
	diagram, err := LoadExistUMLDiagram("existing.uml")
	assert.NoError(t, err)
	assert.NotNil(t, diagram)
	assert.Equal(t, "existing.uml", diagram.GetName())
	assert.Equal(t, DiagramType(ClassDiagram), diagram.GetDiagramType())
}

func TestCheckDiagramType(t *testing.T) {
	tests := []struct {
		name        string
		diagramType DiagramType
		expected    bool
	}{
		{
			name:        "ClassDiagram",
			diagramType: ClassDiagram,
			expected:    true,
		},
		{
			name:        "UseCaseDiagram",
			diagramType: UseCaseDiagram,
			expected:    false,
		},
		{
			name:        "SequenceDiagram",
			diagramType: SequenceDiagram,
			expected:    false,
		},
		{
			name:        "InvalidDiagram",
			diagramType: DiagramType(8),
			expected:    false,
		},
		{
			name:        "ZeroValue",
			diagramType: DiagramType(0),
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			noError := validateDiagramType(tt.diagramType) == nil
			assert.Equal(t, tt.expected, noError)
		})
	}
}

func TestUMLDiagramGetters(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, diagram)

	// Test GetName
	assert.Equal(t, "TestDiagram", diagram.GetName())

	// Test GetDiagramType
	assert.Equal(t, DiagramType(ClassDiagram), diagram.GetDiagramType())

	// Test GetLastModified
	assert.WithinDuration(t, time.Now(), diagram.GetLastModified(), time.Second)

	// Test GetDrawData
	drawData := diagram.GetDrawData()
	assert.Equal(t, drawdata.Margin, drawData.Margin)
	assert.Equal(t, drawdata.LineWidth, drawData.LineWidth)
	assert.Equal(t, drawdata.DefaultDiagramColor, drawData.Color)

	err = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 20}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.NoError(t, err)
}
func TestUMLDiagram_AddGadget(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("AddGadgetTest.uml", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, diagram)
	err = diagram.AddGadget(component.Class, utils.Point{X: 5, Y: 5}, 0, drawdata.DefaultGadgetColor, "")
	assert.NoError(t, err)

	// Should have one gadget in container
	all := diagram.componentsContainer.GetAll()
	assert.Len(t, all, 1)
}

func TestUMLDiagram_AddGadget_InvalidType(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("AddGadgetInvalid.uml", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, diagram)
	// Use an invalid gadget type (assuming -1 is invalid)
	err = diagram.AddGadget(component.GadgetType(-1), utils.Point{X: 1, Y: 1}, 0, drawdata.DefaultGadgetColor, "")
	assert.Error(t, err)
}

func TestUMLDiagram_validatePoint(t *testing.T) {
	diagram, _ := CreateEmptyUMLDiagram("ValidatePoint.uml", ClassDiagram)
	assert.NoError(t, diagram.validatePoint(utils.Point{X: 0, Y: 0}))
	assert.NoError(t, diagram.validatePoint(utils.Point{X: 10, Y: 10}))
	assert.Error(t, diagram.validatePoint(utils.Point{X: -1, Y: 0}))
	assert.Error(t, diagram.validatePoint(utils.Point{X: 0, Y: -1}))
}

func TestUMLDiagram_StartAddAssociation_InvalidPoint(t *testing.T) {
	diagram, _ := CreateEmptyUMLDiagram("StartAddAssInvalid.uml", ClassDiagram)
	err := diagram.StartAddAssociation(utils.Point{X: -1, Y: 0})
	assert.Error(t, err)
}

func TestUMLDiagram_SelectAndUnselectComponent(t *testing.T) {
	diagram, _ := CreateEmptyUMLDiagram("SelectUnselect.uml", ClassDiagram)
	_ = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "")

	// Select
	err := diagram.SelectComponent(utils.Point{X: 10, Y: 10})
	assert.NoError(t, err)
	assert.Len(t, diagram.componentsSelected, 1)

	// Unselect
	err = diagram.UnselectComponent(utils.Point{X: 10, Y: 10})
	assert.NoError(t, err)
	assert.Len(t, diagram.componentsSelected, 0)
}

func TestUMLDiagram_UnselectAllComponents(t *testing.T) {
	diagram, _ := CreateEmptyUMLDiagram("UnselectAll.uml", ClassDiagram)
	_ = diagram.AddGadget(component.Class, utils.Point{X: 1, Y: 1}, 0, drawdata.DefaultGadgetColor, "")
	_ = diagram.SelectComponent(utils.Point{X: 1, Y: 1})
	assert.Len(t, diagram.componentsSelected, 1)
	_ = diagram.UnselectAllComponents()
	assert.Len(t, diagram.componentsSelected, 0)
}

// func TestUMLDiagram_RegisterNotifyDrawUpdate(t *testing.T) {
// 	diagram, _ := CreateEmptyUMLDiagram("NotifyDrawUpdate.uml", ClassDiagram)
// 	err := diagram.RegisterNotifyDrawUpdate(nil)
// 	assert.Error(t, err)

// 	called := false
// 	err = diagram.RegisterNotifyDrawUpdate(func() duerror.DUError {
// 		called = true
// 		return nil
// 	})
// 	assert.NoError(t, err)
// 	_ = diagram.updateDrawData()
// 	assert.True(t, called)
// }

func TestUMLDiagram_GetDrawData(t *testing.T) {
	diagram, _ := CreateEmptyUMLDiagram("GetDrawData.uml", ClassDiagram)
	data := diagram.GetDrawData()
	assert.Equal(t, drawdata.Margin, data.Margin)
	assert.Equal(t, drawdata.LineWidth, data.LineWidth)
	assert.Equal(t, drawdata.DefaultDiagramColor, data.Color)
}

func TestUMLDiagram_RemoveSelectedComponents(t *testing.T) {
	diagram, _ := CreateEmptyUMLDiagram("RemoveSelected.uml", ClassDiagram)
	_ = diagram.AddGadget(component.Class, utils.Point{X: 2, Y: 2}, 0, drawdata.DefaultGadgetColor, "")
	_ = diagram.SelectComponent(utils.Point{X: 2, Y: 2})
	assert.Len(t, diagram.componentsSelected, 1)
	err := diagram.RemoveSelectedComponents()
	assert.NoError(t, err)
	assert.Len(t, diagram.componentsSelected, 0)
	assert.Len(t, diagram.componentsContainer.GetAll(), 0)
}

func TestUMLDiagram_EndAddAssociation_NoGadget(t *testing.T) {
	diagram, _ := CreateEmptyUMLDiagram("EndAddAssNoGadget.uml", ClassDiagram)
	_ = diagram.StartAddAssociation(utils.Point{X: 1, Y: 1})
	err := diagram.EndAddAssociation(component.AssociationType(0), utils.Point{X: 2, Y: 2})
	assert.Error(t, err)
}

func TestUMLDiagram_EndAddAssociation_InvalidPoint(t *testing.T) {
	diagram, _ := CreateEmptyUMLDiagram("EndAddAssInvalidPoint.uml", ClassDiagram)
	_ = diagram.StartAddAssociation(utils.Point{X: 1, Y: 1})
	err := diagram.EndAddAssociation(component.AssociationType(0), utils.Point{X: -1, Y: 2})
	assert.Error(t, err)
}

func TestUMLDiagram_LoadExistUMLDiagram(t *testing.T) {
	// TODO
}

func TestValidatePoint(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)

	// Valid point
	err = diagram.validatePoint(utils.Point{X: 10, Y: 20})
	assert.NoError(t, err)

	// Negative X
	err = diagram.validatePoint(utils.Point{X: -5, Y: 20})
	assert.Error(t, err)
	assert.Equal(t, "point coordinates must be non-negative", err.Error())

	// Negative Y
	err = diagram.validatePoint(utils.Point{X: 10, Y: -5})
	assert.Error(t, err)
	assert.Equal(t, "point coordinates must be non-negative", err.Error())

	// Both negative
	err = diagram.validatePoint(utils.Point{X: -10, Y: -10})
	assert.Error(t, err)
	assert.Equal(t, "point coordinates must be non-negative", err.Error())
}

func TestAddGadget(t *testing.T) {
	// TODO
}

func TestRemoveGadget(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)

	// This is a placeholder test since removeGadget is currently a no-op
	gad := &component.Gadget{}
	err = diagram.removeGadget(gad)
	assert.NoError(t, err)
}

func TestAssociationMethods(t *testing.T) {
	// TODO

	diagram, _ := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)

	gadgetPoint0 := utils.Point{X: 0, Y: 0}
	gadgetPoint1 := utils.Point{X: 0, Y: 200}
	gadgetPoint2 := utils.Point{X: 200, Y: 0}
	gadgetPoint3 := utils.Point{X: 200, Y: 200}
	diagram.AddGadget(component.Class, gadgetPoint0, 0, drawdata.DefaultGadgetColor, "")
	diagram.AddGadget(component.Class, gadgetPoint1, 0, drawdata.DefaultGadgetColor, "")
	diagram.AddGadget(component.Class, gadgetPoint2, 0, drawdata.DefaultGadgetColor, "")
	diagram.AddGadget(component.Class, gadgetPoint3, 0, drawdata.DefaultGadgetColor, "")
	gad0, _ := diagram.componentsContainer.SearchGadget(gadgetPoint0)
	gad1, _ := diagram.componentsContainer.SearchGadget(gadgetPoint1)
	gad2, _ := diagram.componentsContainer.SearchGadget(gadgetPoint2)
	gad3, _ := diagram.componentsContainer.SearchGadget(gadgetPoint3)

	t.Run("StartAddAssociation", func(t *testing.T) {
		err := diagram.StartAddAssociation(gadgetPoint0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	var a *component.Association
	t.Run("EndAddAssociation", func(t *testing.T) {
		err := diagram.EndAddAssociation(component.Extension, gadgetPoint1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		for _, c := range diagram.componentsContainer.GetAll() {
			switch tmp := c.(type) {
			case *component.Association:
				a = tmp
			}
			if a != nil {
				break
			}
		}
		if a == nil {
			t.Errorf("add association fail")
		}
		if a.GetParentStart() != gad0 {
			t.Errorf("incorrect start parent")
		}
		if a.GetParentEnd() != gad1 {
			t.Errorf("incorrect end parent")
		}

		stList := diagram.associations[gad0][0]
		if !slices.Contains(stList, a) {
			t.Errorf("not in stList of diagram.associations")
		}
		enList := diagram.associations[gad1][1]
		if !slices.Contains(enList, a) {
			t.Errorf("not in enList of diagram.associations")
		}
	})

	t.Run("SetParentStartComponent", func(t *testing.T) {
		if a == nil {
			t.Errorf("add association fail")
		}

		diagram.UnselectAllComponents()
		dd := a.GetDrawData().(drawdata.Association)
		midPoint := utils.Point{
			X: (dd.StartX + dd.EndX) / 2,
			Y: (dd.StartY + dd.EndY) / 2,
		}
		diagram.SelectComponent(midPoint)
		c, _ := diagram.getSelectedComponent()
		if a != c.(*component.Association) {
			t.Errorf("can not select added association")
		}

		err := diagram.SetParentStartComponent(gadgetPoint2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if a.GetParentStart() != gad2 {
			t.Errorf("incorrect start parent")
		}
		stListOld := diagram.associations[gad0][0]
		stListNew := diagram.associations[gad2][0]
		if slices.Contains(stListOld, a) {
			t.Errorf("updated association in old stList")
		}
		if !slices.Contains(stListNew, a) {
			t.Errorf("updated association not in new stList")
		}
	})

	t.Run("SetParentEndComponent", func(t *testing.T) {
		if a == nil {
			t.Errorf("add association fail")
		}

		diagram.UnselectAllComponents()
		dd := a.GetDrawData().(drawdata.Association)
		midPoint := utils.Point{
			X: (dd.StartX + dd.EndX) / 2,
			Y: (dd.StartY + dd.EndY) / 2,
		}
		diagram.SelectComponent(midPoint)
		c, _ := diagram.getSelectedComponent()
		if a != c.(*component.Association) {
			t.Errorf("can not select added association")
		}

		err := diagram.SetParentEndComponent(gadgetPoint3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if a.GetParentEnd() != gad3 {
			t.Errorf("incorrect end parent")
		}
		enListOld := diagram.associations[gad1][1]
		enListNew := diagram.associations[gad3][1]
		if slices.Contains(enListOld, a) {
			t.Errorf("updated association in old enList")
		}
		if !slices.Contains(enListNew, a) {
			t.Errorf("updated association not in new enList")
		}
	})
}

func TestRegisterUpdateParentDraw(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)

	// Try to register nil function
	err = diagram.RegisterUpdateParentDraw(nil)
	assert.Error(t, err)
	assert.Equal(t, "update function cannot be nil", err.Error())

	// Register valid function
	updateCalled := false
	updateFunc := func() duerror.DUError {
		updateCalled = true
		return nil
	}

	err = diagram.RegisterUpdateParentDraw(updateFunc)
	assert.NoError(t, err)

	// Test that the function gets called through updateDrawData
	err = diagram.updateDrawData()
	assert.NoError(t, err)
	assert.True(t, updateCalled)
}

func TestUpdateDrawData(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)

	// Add a gadget
	err = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 20}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.NoError(t, err)

	// Check that drawData contains the gadget
	assert.Equal(t, 1, len(diagram.drawData.Gadgets))

	// Manual update of drawData
	err = diagram.updateDrawData()
	assert.NoError(t, err)

	// Test with a registered update function
	updateCalled := false
	updateFunc := func() duerror.DUError {
		updateCalled = true
		return nil
	}

	err = diagram.RegisterUpdateParentDraw(updateFunc)
	assert.NoError(t, err)

	err = diagram.updateDrawData()
	assert.NoError(t, err)
	assert.True(t, updateCalled)
}

func TestAddAttributeToGadget(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)

	// Try to add an attribute with no selected components
	err = diagram.AddAttributeToGadget(0, "attribute")
	assert.Error(t, err)
	assert.Equal(t, "can only operate on one component", err.Error())

	// Add a gadget to the diagram
	err = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 20}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.NoError(t, err)

	// Get the gadget from the container
	components := diagram.componentsContainer.GetAll()
	assert.Equal(t, 1, len(components))

	gadget, ok := components[0].(*component.Gadget)
	assert.True(t, ok)

	// Select the gadget
	diagram.componentsSelected[gadget] = true

	// Add attribute
	err = diagram.AddAttributeToGadget(0, "NewAttribute")
	assert.NoError(t, err)

	// Test with multiple selected components
	// First clear the selection
	err = diagram.UnselectAllComponents()
	assert.NoError(t, err)

	// Add a second gadget
	err = diagram.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "sample header 2")
	assert.NoError(t, err)

	// Get all gadgets
	components = diagram.componentsContainer.GetAll()
	assert.Equal(t, 2, len(components))

	// Select both gadgets
	for _, comp := range components {
		diagram.componentsSelected[comp] = true
	}

	// Try to add attribute with multiple gadgets selected
	err = diagram.AddAttributeToGadget(0, "attribute")
	assert.Error(t, err)
	assert.Equal(t, "can only operate on one component", err.Error())
}

func TestAddAttributeToAssociation(t *testing.T) {
	// TODO
}

func TestRemoveAttributeFromAssociation(t *testing.T) {
	// TODO
}

func TestAttributeSettersGadget(t *testing.T) {
	// TODO
}

func TestAttributeSettersAssociation(t *testing.T) {
	// TODO
}

// Mock container for testing selection methods
type mockContainer struct {
	mockComponent component.Component
	pointToMatch  utils.Point
}

func (m *mockContainer) Insert(c component.Component) duerror.DUError {
	return nil
}

func (m *mockContainer) Remove(c component.Component) duerror.DUError {
	return nil
}

func (m *mockContainer) Search(p utils.Point) (component.Component, duerror.DUError) {
	if p.X == m.pointToMatch.X && p.Y == m.pointToMatch.Y {
		return m.mockComponent, nil
	}
	return nil, nil
}

func (m *mockContainer) GetAll() []component.Component {
	if m.mockComponent != nil {
		return []component.Component{m.mockComponent}
	}
	return []component.Component{}
}

func (m *mockContainer) Len() (int, duerror.DUError) {
	if m.mockComponent != nil {
		return 1, nil
	}
	return 0, nil
}
