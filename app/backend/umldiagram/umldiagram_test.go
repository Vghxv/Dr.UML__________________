package umldiagram

import (
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
			expectError: false,
		},
		{
			name:        "ValidSequenceDiagram",
			inputName:   "test3.uml",
			diagramType: SequenceDiagram,
			expectError: false,
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
				assert.Equal(t, utils.Color{R: 255, G: 255, B: 255}, diagram.backgroundColor)
				assert.NotNil(t, diagram.componentsContainer)
				assert.NotNil(t, diagram.componentsGraph)
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
			expected:    true,
		},
		{
			name:        "SequenceDiagram",
			diagramType: SequenceDiagram,
			expected:    true,
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

	err = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 20}, 0, 0x808080)
	assert.NoError(t, err)
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
}

func TestComponentSelection(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)

	// Add a gadget to select
	err = diagram.AddGadget(component.Class, utils.Point{X: 50, Y: 50}, 0, 0x808080)
	assert.NoError(t, err)

	// Initially no components are selected
	assert.Equal(t, 0, len(diagram.componentsSelected))

	// Mock for ComponentsContainer.Search since we can't directly test it
	// We'll replace it with a custom implementation for testing
	originalContainer := diagram.componentsContainer

	// Create a mock container that returns a component for a specific point
	mockGadget := &component.Gadget{}
	mockContainer := &mockContainer{
		mockComponent: mockGadget,
		pointToMatch:  utils.Point{X: 50, Y: 50},
	}

	diagram.componentsContainer = mockContainer

	// Test SelectComponent
	err = diagram.SelectComponent(utils.Point{X: 50, Y: 50})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(diagram.componentsSelected))
	assert.True(t, diagram.componentsSelected[mockGadget])

	// Test SelectComponent with no component found
	err = diagram.SelectComponent(utils.Point{X: 100, Y: 100})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(diagram.componentsSelected))

	// Test UnselectComponent
	err = diagram.UnselectComponent(utils.Point{X: 50, Y: 50})
	assert.NoError(t, err)
	assert.Equal(t, 0, len(diagram.componentsSelected))

	// Test UnselectAllComponents
	diagram.componentsSelected[mockGadget] = true
	assert.Equal(t, 1, len(diagram.componentsSelected))

	err = diagram.UnselectAllComponents()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(diagram.componentsSelected))

	// Restore original container
	diagram.componentsContainer = originalContainer
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
	err = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 20}, 0, 0x808080)
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

	// Try to add attribute with no selected components
	err = diagram.AddAttributeToGadget("attribute", 0)
	assert.Error(t, err)
	assert.Equal(t, "can only add attribute to one gadget", err.Error())

	// Add a gadget to the diagram
	err = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 20}, 0, 0x808080)
	assert.NoError(t, err)

	// Get the gadget from the container
	components := diagram.componentsContainer.GetAll()
	assert.Equal(t, 1, len(components))

	gadget, ok := components[0].(*component.Gadget)
	assert.True(t, ok)

	// Select the gadget
	diagram.componentsSelected[gadget] = true

	// Add attribute
	err = diagram.AddAttributeToGadget("NewAttribute", 0)
	assert.NoError(t, err)

	// Test with multiple selected components
	// First clear the selection
	err = diagram.UnselectAllComponents()
	assert.NoError(t, err)

	// Add a second gadget
	err = diagram.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, 0x808080)
	assert.NoError(t, err)

	// Get all gadgets
	components = diagram.componentsContainer.GetAll()
	assert.Equal(t, 2, len(components))

	// Select both gadgets
	for _, comp := range components {
		diagram.componentsSelected[comp] = true
	}

	// Try to add attribute with multiple gadgets selected
	err = diagram.AddAttributeToGadget("attribute", 0)
	assert.Error(t, err)
	assert.Equal(t, "can only add attribute to one gadget", err.Error())
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
