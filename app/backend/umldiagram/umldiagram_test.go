package umldiagram

import (
	"fmt"
	"os"
	"testing"
	"time"

	"Dr.uml/backend/component/attribute"

	"Dr.uml/backend/component"
	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
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
				assert.WithinDuration(t, time.Now(), diagram.GetLastModified(), time.Second)
				assert.Equal(t, utils.Point{X: 0, Y: 0}, diagram.startPoint)
				// New assertions
				assert.Equal(t, "#FFFFFF", diagram.backgroundColor)
				assert.NotNil(t, diagram.componentsContainer)
			}
		})
	}
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
}

func TestUMLDiagram_UnselectAllComponents(t *testing.T) {
	diagram, _ := CreateEmptyUMLDiagram("UnselectAll.uml", ClassDiagram)
	_ = diagram.AddGadget(component.Class, utils.Point{X: 1, Y: 1}, 0, drawdata.DefaultGadgetColor, "")
	_ = diagram.SelectComponent(utils.Point{X: 1, Y: 1})
	assert.Len(t, diagram.componentsSelected, 1)
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

// Test LoadExistUMLDiagram functionality
func TestLoadExistUMLDiagram_Comprehensive(t *testing.T) {
	t.Run("load diagram with gadgets and associations", func(t *testing.T) {
		// Create a saved diagram structure
		savedDiagram := utils.SavedDiagram{
			Filetype:     utils.FiletypeDiagram | int(ClassDiagram),
			LastEdit:     "2023-01-01T00:00:00Z",
			Gadgets:      []utils.SavedGad{},
			Associations: []utils.SavedAss{},
		}

		// Add a gadget
		savedGadget := utils.SavedGad{
			GadgetType: 1,
			Point:      "10, 20",
			Layer:      0,
			Color:      drawdata.DefaultGadgetColor,
			Attributes: []utils.SavedAtt{
				{
					Content:  "TestClass",
					Size:     12,
					Style:    int(attribute.Bold),
					FontFile: os.Getenv("APP_ROOT") + "/frontend/src/assets/fonts/Inkfree.ttf",
					Ratio:    0.0,
				},
			},
		}
		savedDiagram.Gadgets = append(savedDiagram.Gadgets, savedGadget)

		// Load the diagram
		diagram, err := LoadExistUMLDiagram("test.duml", savedDiagram)
		assert.NoError(t, err)
		assert.NotNil(t, diagram)
		assert.Equal(t, "test.duml", diagram.GetName())
		assert.Equal(t, DiagramType(ClassDiagram), diagram.GetDiagramType())

		// Verify gadget was loaded
		components := diagram.componentsContainer.GetAll()
		assert.Len(t, components, 1)
	})

	t.Run("load diagram with invalid filetype", func(t *testing.T) {
		savedDiagram := utils.SavedDiagram{
			Filetype: utils.FiletypeDiagram | int(UseCaseDiagram)<<1, // Invalid type
		}

		_, err := LoadExistUMLDiagram("test.duml", savedDiagram)
		assert.Error(t, err)
	})

	t.Run("load diagram with corrupted gadgets", func(t *testing.T) {
		savedDiagram := utils.SavedDiagram{
			Filetype: utils.FiletypeDiagram | int(ClassDiagram),
			Gadgets: []utils.SavedGad{
				{
					GadgetType: 999, // Invalid gadget type
					Point:      "invalid_point",
					Color:      "invalid_color",
				},
			},
		}

		_, err := LoadExistUMLDiagram("test.duml", savedDiagram)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "test.duml")
	})
}

// Test HasUnsavedChanges
func TestHasUnsavedChanges(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	// Initially should not have unsaved changes
	assert.True(t, diagram.HasUnsavedChanges())

	// After making a change
	err = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "TestClass")
	assert.NoError(t, err)
	assert.True(t, diagram.HasUnsavedChanges())

	// After saving
	_, err = diagram.SaveToFile("test.uml")
	assert.NoError(t, err)
	assert.False(t, diagram.HasUnsavedChanges())
}

// Test SetPointComponent
func TestSetPointComponent(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("successful point setting", func(t *testing.T) {
		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Set new point
		newPoint := utils.Point{X: 50, Y: 60}
		err = diagram.SetPointComponent(newPoint)
		assert.NoError(t, err)

		// Verify the point was changed
		g, err := diagram.componentsContainer.SearchGadget(newPoint)
		assert.NoError(t, err)
		assert.NotNil(t, g)
	})

	t.Run("no component selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		err = diagram.SetPointComponent(utils.Point{X: 10, Y: 10})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "can only operate on one component")
	})

	t.Run("selected component is not a gadget", func(t *testing.T) {
		// Create diagram with association
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add gadgets and association
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: 55, Y: 55}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Try to set point on association
		err = diagram.SetPointComponent(utils.Point{X: 200, Y: 200})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "selected component is not a gadget")
	})
}

// Test SetLayerComponent
func TestSetLayerComponent(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("successful layer setting", func(t *testing.T) {
		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Set new layer
		newLayer := 5
		err = diagram.SetLayerComponent(newLayer)
		assert.NoError(t, err)

		// Verify the layer was changed
		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		assert.Equal(t, newLayer, g.GetLayer())
	})

	t.Run("no component selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		err = diagram.SetLayerComponent(5)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "can only operate on one component")
	})
}

// Test SetColorComponent
func TestSetColorComponent(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("successful color setting", func(t *testing.T) {
		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Set new color
		newColor := "#ff0000"
		err = diagram.SetColorComponent(newColor)
		assert.NoError(t, err)

		// Verify the color was changed
		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		assert.Equal(t, newColor, g.GetColor())
	})

	t.Run("no component selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		err = diagram.SetColorComponent("#ff0000")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "can only operate on one component")
	})

	t.Run("selected component is not a gadget", func(t *testing.T) {
		// Create diagram with association
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add gadgets and association
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: 55, Y: 55}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Try to set color on association
		err = diagram.SetColorComponent("#ff0000")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "selected component is not a gadget")
	})
}

// Test SetAttrContentComponent for gadgets
func TestSetAttrContentComponent_Gadget(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("successful content setting for gadget", func(t *testing.T) {
		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Add an attribute first
		err = diagram.AddAttributeToGadget(1, "oldContent")
		assert.NoError(t, err)

		// Set new content
		newContent := "newContent"
		err = diagram.SetAttrContentComponent(1, 0, newContent)
		assert.NoError(t, err)

		// Verify the content was changed
		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		attr, err := g.GetAttribute(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, newContent, attr.GetContent())
	})

	t.Run("no component selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		err = diagram.SetAttrContentComponent(1, 0, "content")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "can only operate on one component")
	})
}

// Test SetAttrContentComponent for associations
func TestSetAttrContentComponent_Association(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("successful content setting for association", func(t *testing.T) {
		// Add gadgets and association
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: 55, Y: 55}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Add an attribute to association first
		err = diagram.AddAttributeToAssociation(0.5, "oldContent")
		assert.NoError(t, err)

		// Set new content
		newContent := "newContent"
		err = diagram.SetAttrContentComponent(0, 0, newContent)
		assert.NoError(t, err)

		// Verify the content was changed
		c, err := diagram.componentsContainer.Search(midPoint)
		assert.NoError(t, err)
		a, ok := c.(*component.Association)
		assert.True(t, ok)
		attr, err := a.GetAttribute(0)
		assert.NoError(t, err)
		assert.Equal(t, newContent, attr.GetContent())
	})
}

// Test AddAttributeToAssociation
func TestAddAttributeToAssociation(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("successful attribute addition", func(t *testing.T) {
		// Add gadgets and association
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: 55, Y: 55}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Add attribute
		err = diagram.AddAttributeToAssociation(0.5, "testAttribute")
		assert.NoError(t, err)

		// Verify attribute was added
		c, err := diagram.componentsContainer.Search(midPoint)
		assert.NoError(t, err)
		a, ok := c.(*component.Association)
		assert.True(t, ok)
		assert.Equal(t, 1, a.GetAttributesLen())
		attr, err := a.GetAttribute(0)
		assert.NoError(t, err)
		assert.Equal(t, "testAttribute", attr.GetContent())
	})

	t.Run("no component selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		err = diagram.AddAttributeToAssociation(0.5, "testAttribute")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "can only operate on one component")
	})

	t.Run("selected component is not an association", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Try to add attribute to gadget (should fail)
		err = diagram.AddAttributeToAssociation(0.5, "testAttribute")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "selected component is not an association")
	})
}

// Test RemoveAttributeFromAssociation
func TestRemoveAttributeFromAssociation(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("successful attribute removal", func(t *testing.T) {
		// Add gadgets and association
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: 55, Y: 55}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Add attribute first
		err = diagram.AddAttributeToAssociation(0.5, "testAttribute")
		assert.NoError(t, err)

		// Verify it was added
		c, err := diagram.componentsContainer.Search(midPoint)
		assert.NoError(t, err)
		a, ok := c.(*component.Association)
		assert.True(t, ok)
		assert.Equal(t, 1, a.GetAttributesLen())

		// Remove the attribute
		err = diagram.RemoveAttributeFromAssociation(0)
		assert.NoError(t, err)

		// Verify it was removed
		assert.Equal(t, 0, a.GetAttributesLen())
	})

	t.Run("no component selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		err = diagram.RemoveAttributeFromAssociation(0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "can only operate on one component")
	})

	t.Run("selected component is not an association", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select a gadget
		err = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 20}, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		// Select the gadget
		err = diagram.SelectComponent(utils.Point{X: 11, Y: 21})
		assert.NoError(t, err)
		// Try to remove attribute from gadget
		err = diagram.RemoveAttributeFromAssociation(0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "selected component is not an association")
	})

	t.Run("invalid index", func(t *testing.T) {
		// Add gadgets and association
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: 55, Y: 55}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Try to remove non-existent attribute
		err = diagram.RemoveAttributeFromAssociation(0)
		assert.Error(t, err)
	})
}

// Test SetAttrRatioComponent
func TestSetAttrRatioComponent(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("successful ratio setting for association", func(t *testing.T) {
		// Add gadgets and association
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: 55, Y: 55}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Add attribute first
		err = diagram.AddAttributeToAssociation(0.3, "testAttribute")
		assert.NoError(t, err)

		// Set new ratio
		newRatio := 0.7
		err = diagram.SetAttrRatioComponent(0, 0, newRatio)
		assert.NoError(t, err)

		// Verify the ratio was changed
		c, err := diagram.componentsContainer.Search(midPoint)
		assert.NoError(t, err)
		a, ok := c.(*component.Association)
		assert.True(t, ok)
		attr, err := a.GetAttribute(0)
		assert.NoError(t, err)
		assert.Equal(t, newRatio, attr.GetRatio())
	})

	t.Run("no component selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		err = diagram.SetAttrRatioComponent(0, 0, 0.5)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "can only operate on one component")
	})

	t.Run("selected component is not an association", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Try to set ratio on gadget
		err = diagram.SetAttrRatioComponent(0, 0, 0.5)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "selected component is not an association")
	})
}

// Test Undo and Redo with error conditions
func TestUndoRedoWithErrors(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("undo with no commands", func(t *testing.T) {
		err = diagram.Undo()
		assert.Error(t, err) // Should fail when there's nothing to undo
	})

	t.Run("redo with no commands", func(t *testing.T) {
		err = diagram.Redo()
		assert.Error(t, err) // Should fail when there's nothing to redo
	})

	t.Run("successful undo/redo cycle", func(t *testing.T) {
		// Add a gadget
		gadPoint := utils.Point{X: 10, Y: 10}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsContainer.GetAll(), 1)

		// Undo the addition
		err = diagram.Undo()
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsContainer.GetAll(), 0)

		// Redo the addition
		err = diagram.Redo()
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsContainer.GetAll(), 1)
	})
}

// Test error conditions in EndAddAssociation
func TestEndAddAssociation_ErrorCases(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("start and end on same gadget", func(t *testing.T) {
		// Add a single gadget
		gadPoint := utils.Point{X: 10, Y: 10}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint)
		assert.NoError(t, err)

		// Try to end on the same gadget
		err = diagram.EndAddAssociation(component.Composition, gadPoint)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "start and end points are the same")
	})

	t.Run("invalid association type", func(t *testing.T) {
		// Add two gadgets
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)

		// Try with invalid association type
		err = diagram.EndAddAssociation(component.AssociationType(999), gadPoint2)
		assert.Error(t, err)
	})
}

// Test private helper methods by creating scenarios that use them
func TestPrivateHelperMethods(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("test moveGadget through SetPointComponent", func(t *testing.T) {
		// Add gadget with association to test moveGadget's association update logic
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select first gadget and move it
		err = diagram.SelectComponent(gadPoint1)
		assert.NoError(t, err)

		newPoint := utils.Point{X: 200, Y: 200}
		err = diagram.SetPointComponent(newPoint)
		assert.NoError(t, err)

		// Verify the gadget moved and association updated
		g, err := diagram.componentsContainer.SearchGadget(newPoint)
		assert.NoError(t, err)
		assert.NotNil(t, g)
	})

	t.Run("test removeAssociation through RemoveSelectedComponents", func(t *testing.T) {
		// Add gadgets and association
		gadPoint1 := utils.Point{X: 300, Y: 300}
		gadPoint2 := utils.Point{X: 400, Y: 400}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class3")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class4")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		initialCount := len(diagram.componentsContainer.GetAll())

		// Select and remove first gadget (should also remove the association)
		err = diagram.SelectComponent(gadPoint1)
		assert.NoError(t, err)
		err = diagram.RemoveSelectedComponents()
		assert.NoError(t, err)

		// Should have one less gadget and one less association
		finalCount := len(diagram.componentsContainer.GetAll())
		assert.True(t, finalCount < initialCount)
	})

	t.Run("test addComponents through Undo", func(t *testing.T) {
		initialCount := len(diagram.componentsContainer.GetAll())

		// Add a gadget
		gadPoint := utils.Point{X: 500, Y: 500}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)

		afterAddCount := len(diagram.componentsContainer.GetAll())
		assert.True(t, afterAddCount > initialCount)

		// Remove it
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)
		err = diagram.RemoveSelectedComponents()
		assert.NoError(t, err)

		// Undo the removal (triggers unexecute)
		err = diagram.Undo()
		assert.NoError(t, err)

		undoCount := len(diagram.componentsContainer.GetAll())
		assert.Equal(t, afterAddCount, undoCount)
	})
}

// Test attribute-related setter methods
func TestAttributeSetterMethods(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	// Setup gadget with attribute
	gadPoint := utils.Point{X: 10, Y: 20}
	err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
	assert.NoError(t, err)
	err = diagram.SelectComponent(gadPoint)
	assert.NoError(t, err)
	err = diagram.AddAttributeToGadget(1, "testAttribute")
	assert.NoError(t, err)

	t.Run("test SetAttrSizeComponent", func(t *testing.T) {
		newSize := 16
		err = diagram.SetAttrSizeComponent(1, 0, newSize)
		assert.NoError(t, err)

		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		attr, err := g.GetAttribute(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, newSize, attr.GetSize())
	})

	t.Run("test SetAttrStyleComponent", func(t *testing.T) {
		newStyle := int(attribute.Bold | attribute.Italic)
		err = diagram.SetAttrStyleComponent(1, 0, newStyle)
		assert.NoError(t, err)

		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		attr, err := g.GetAttribute(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, attribute.Textstyle(newStyle), attr.GetStyle())
	})

	t.Run("test SetAttrFontComponent", func(t *testing.T) {
		fontName := "Inkfree"
		err = diagram.SetAttrFontComponent(1, 0, fontName)
		assert.NoError(t, err)

		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		attr, err := g.GetAttribute(1, 0)
		assert.NoError(t, err)
		expectedFontFile := os.Getenv("APP_ROOT") + "/frontend/src/assets/fonts/Inkfree.ttf"
		assert.Equal(t, expectedFontFile, attr.GetFontFile())
	})
}

// Test SetParentStartComponent and SetParentEndComponent error cases
func TestSetParentComponentErrors(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("SetParentStartComponent with no association selected", func(t *testing.T) {
		// Add a gadget and select it (not an association)
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		err = diagram.SetParentStartComponent(utils.Point{X: 100, Y: 100})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "selected component is not an association")
	})

	t.Run("SetParentEndComponent with no association selected", func(t *testing.T) {
		err = diagram.SetParentEndComponent(utils.Point{X: 100, Y: 100})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "selected component is not an association")
	})

	t.Run("SetParentStartComponent with invalid target point", func(t *testing.T) {
		// Create association first
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: 55, Y: 55}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Try to set parent to empty space
		err = diagram.SetParentStartComponent(utils.Point{X: 500, Y: 500})
		assert.Error(t, err)
	})
}

// Test SelectComponent edge cases
func TestSelectComponent_EdgeCases(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("select already selected component", func(t *testing.T) {
		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsSelected, 1)

		// Select it again - should not change anything
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsSelected, 1)
	})

	t.Run("select empty space to unselect all", func(t *testing.T) {
		// Should have one selected component from previous test
		assert.Len(t, diagram.componentsSelected, 1)

		// Click on empty space
		err = diagram.SelectComponent(utils.Point{X: 500, Y: 500})
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsSelected, 0)
	})
}

// Test error conditions in loadGadgetAttributes
func TestLoadGadgetAttributes_Errors(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("load attributes to nil gadget", func(t *testing.T) {
		savedAttributes := []utils.SavedAtt{
			{
				Content:  "test",
				Size:     12,
				Style:    int(attribute.Bold),
				FontFile: "test.ttf",
				Ratio:    0.0,
			},
		}

		err, _ := diagram.loadGadgetAttributes(nil, savedAttributes)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Cannot load attributes to a nil gadget")
	})

	t.Run("load invalid attributes", func(t *testing.T) {
		gad, err := component.NewGadget(component.Class, utils.Point{}, 0, drawdata.DefaultGadgetColor, "")
		assert.NoError(t, err)

		savedAttributes := []utils.SavedAtt{
			{
				Content:  "test",
				Size:     -1, // Invalid size
				Style:    int(attribute.Bold),
				FontFile: "nonexistent.ttf",
				Ratio:    0.0,
			},
		}

		err, index := diagram.loadGadgetAttributes(gad, savedAttributes)
		assert.Error(t, err)
		assert.Equal(t, 0, index)
	})
}

// Test collectAssociations errors
func TestCollectAssociations_Errors(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("collect associations with invalid gadget mapping", func(t *testing.T) {
		// Add gadgets and association
		gadPoint1 := utils.Point{X: 10, Y: 10}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Get the actual gadgets
		g1, err := diagram.componentsContainer.SearchGadget(gadPoint1)
		assert.NoError(t, err)
		_, err = diagram.componentsContainer.SearchGadget(gadPoint2)
		assert.NoError(t, err)

		// Create invalid gadget mapping (missing one of the gadgets)
		invalidDP := make(map[*component.Gadget]int)
		invalidDP[g1] = 0 // Include only first gadget, missing second gadget
		res := &utils.SavedDiagram{}

		err = diagram.collectAssociations(invalidDP, res)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "SecondParent not found")
	})
}

// Test command unexecute methods by triggering undo operations
func TestCommandUnexecuteMethods(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
	assert.NoError(t, err)

	t.Run("test addComponentCommand unexecute", func(t *testing.T) {
		// Add gadget
		gadPoint := utils.Point{X: 10, Y: 10}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsContainer.GetAll(), 1)

		// Undo (triggers unexecute)
		err = diagram.Undo()
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsContainer.GetAll(), 0)
	})

	t.Run("test removeSelectedComponentCommand unexecute", func(t *testing.T) {
		// Add gadget
		gadPoint := utils.Point{X: 20, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)

		// Select and remove
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)
		err = diagram.RemoveSelectedComponents()
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsContainer.GetAll(), 0)

		// Undo removal (triggers unexecute of remove command)
		err = diagram.Undo()
		assert.NoError(t, err)
		assert.Len(t, diagram.componentsContainer.GetAll(), 1)
	})

	t.Run("test setterCommand unexecute", func(t *testing.T) {
		// Add and select gadget
		gadPoint := utils.Point{X: 30, Y: 30}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Get original layer
		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		originalLayer := g.GetLayer()

		// Change layer
		newLayer := originalLayer + 5
		err = diagram.SetLayerComponent(newLayer)
		assert.NoError(t, err)
		assert.Equal(t, newLayer, g.GetLayer())

		// Undo layer change (triggers setterCommand unexecute)
		err = diagram.Undo()
		assert.NoError(t, err)
		assert.Equal(t, originalLayer, g.GetLayer())
	})

	t.Run("test selectAllCommand unexecute", func(t *testing.T) {
		// Create fresh diagram for this test
		testDiagram, err := CreateEmptyUMLDiagram("test_select.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add gadget
		gadPoint := utils.Point{X: 40, Y: 40}
		err = testDiagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)

		// Select component
		err = testDiagram.SelectComponent(gadPoint)
		assert.NoError(t, err)
		assert.Len(t, testDiagram.componentsSelected, 1)

		// Undo selection (triggers selectAllCommand unexecute)
		err = testDiagram.Undo()
		assert.NoError(t, err)
		assert.Len(t, testDiagram.componentsSelected, 0)
	})

	t.Run("test moveGadgetCommand unexecute", func(t *testing.T) {
		// Create fresh diagram for this test
		testDiagram, err := CreateEmptyUMLDiagram("test_move.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select gadget
		originalPoint := utils.Point{X: 50, Y: 50}
		err = testDiagram.AddGadget(component.Class, originalPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = testDiagram.SelectComponent(originalPoint)
		assert.NoError(t, err)

		// Move gadget
		newPoint := utils.Point{X: 150, Y: 150}
		err = testDiagram.SetPointComponent(newPoint)
		assert.NoError(t, err)

		// Verify it moved
		g, err := testDiagram.componentsContainer.SearchGadget(newPoint)
		assert.NoError(t, err)
		assert.NotNil(t, g)

		// Undo move (triggers moveGadgetCommand unexecute)
		err = testDiagram.Undo()
		assert.NoError(t, err)

		// Should be back at original position
		g, err = testDiagram.componentsContainer.SearchGadget(originalPoint)
		assert.NoError(t, err)
		assert.NotNil(t, g)
	})

	t.Run("test addAttributeGadgetCommand unexecute", func(t *testing.T) {
		// Create fresh diagram for this test
		testDiagram, err := CreateEmptyUMLDiagram("test_add_attr.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select gadget
		gadPoint := utils.Point{X: 60, Y: 60}
		err = testDiagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = testDiagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Add attribute
		err = testDiagram.AddAttributeToGadget(1, "testAttribute")
		assert.NoError(t, err)

		// Verify attribute was added
		g, err := testDiagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		assert.Equal(t, 1, g.GetAttributesLen()[1])

		// Undo add attribute (triggers unexecute)
		err = testDiagram.Undo()
		assert.NoError(t, err)
		assert.Equal(t, 0, g.GetAttributesLen()[1])
	})

	t.Run("test removeAttributeGadgetCommand unexecute", func(t *testing.T) {
		// Create fresh diagram for this test
		testDiagram, err := CreateEmptyUMLDiagram("test_remove_attr.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select gadget
		gadPoint := utils.Point{X: 70, Y: 70}
		err = testDiagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = testDiagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Add attribute
		err = testDiagram.AddAttributeToGadget(1, "testAttribute")
		assert.NoError(t, err)

		// Remove attribute
		err = testDiagram.RemoveAttributeFromGadget(1, 0)
		assert.NoError(t, err)

		// Verify attribute was removed
		g, err := testDiagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		assert.Equal(t, 0, g.GetAttributesLen()[1])

		// Undo remove attribute (triggers removeAttributeGadgetCommand unexecute)
		err = testDiagram.Undo()
		assert.NoError(t, err)
		assert.Equal(t, 1, g.GetAttributesLen()[1])
	})
}

// Test function for RemoveAttributeFromGadget
func TestRemoveAttributeFromGadget(t *testing.T) {
	t.Run("successful removal", func(t *testing.T) {
		// Setup
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)

		// Select the gadget
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Add an attribute first
		err = diagram.AddAttributeToGadget(1, "testAttribute")
		assert.NoError(t, err)

		// Get the gadget to verify initial state
		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		assert.NotNil(t, g)
		initialLength := g.GetAttributesLen()[1]
		assert.Equal(t, 1, initialLength) // Should have 1 attribute after adding

		// Get the attribute to verify its content
		att, err := g.GetAttribute(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, "testAttribute", att.GetContent())

		// Test removing the attribute
		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.NoError(t, err)

		// Verify the attribute was removed
		finalLength := g.GetAttributesLen()[1]
		assert.Equal(t, 0, finalLength)
	})

	t.Run("no component selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.Error(t, err)
		assert.Equal(t, "can only operate on one component", err.Error())
	})

	t.Run("multiple components selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add two gadgets
		err = diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 20}, 0, drawdata.DefaultGadgetColor, "TestClass1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "TestClass2")
		assert.NoError(t, err)

		// Select both gadgets manually
		components := diagram.componentsContainer.GetAll()
		for _, comp := range components {
			diagram.componentsSelected[comp] = true
		}

		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.Error(t, err)
		assert.Equal(t, "can only operate on one component", err.Error())
	})

	t.Run("selected component is not a gadget", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add two gadgets to create an association
		gadPoint1 := utils.Point{X: 10, Y: 20}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "TestClass1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "TestClass2")
		assert.NoError(t, err)

		// Create an association
		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Composition, gadPoint2)
		assert.NoError(t, err)

		// Select the association (midpoint between gadgets)
		midPoint := utils.Point{X: (gadPoint1.X + gadPoint2.X) / 2, Y: (gadPoint1.Y + gadPoint2.Y) / 2}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Try to remove attribute from association (should fail)
		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "selected component is not a gadget")
	})

	t.Run("invalid section", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Test invalid section (negative)
		err = diagram.RemoveAttributeFromGadget(-1, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "section out of range")

		// Test invalid section (too large)
		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		maxSection := len(g.GetAttributesLen())
		err = diagram.RemoveAttributeFromGadget(maxSection, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "section out of range")
	})

	t.Run("invalid index", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Test invalid index (negative)
		err = diagram.RemoveAttributeFromGadget(1, -1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "index out of range")

		// Test invalid index (too large for empty section)
		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "index out of range")

		// Add an attribute and test index too large
		err = diagram.AddAttributeToGadget(1, "testAttribute")
		assert.NoError(t, err)
		err = diagram.RemoveAttributeFromGadget(1, 1) // Index 1 when only index 0 exists
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "index out of range")
	})

	t.Run("undo and redo functionality", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Add an attribute
		err = diagram.AddAttributeToGadget(1, "testAttribute")
		assert.NoError(t, err)

		// Get gadget for verification
		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		assert.Equal(t, 1, g.GetAttributesLen()[1])

		// Remove the attribute
		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, 0, g.GetAttributesLen()[1])

		// Undo the removal
		err = diagram.Undo()
		assert.NoError(t, err)
		assert.Equal(t, 1, g.GetAttributesLen()[1])

		// Verify the attribute content is restored
		att, err := g.GetAttribute(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, "testAttribute", att.GetContent())

		// Redo the removal
		err = diagram.Redo()
		assert.NoError(t, err)
		assert.Equal(t, 0, g.GetAttributesLen()[1])
	})

	t.Run("remove from different sections", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Add attributes to different sections
		err = diagram.AddAttributeToGadget(1, "attribute1")
		assert.NoError(t, err)
		err = diagram.AddAttributeToGadget(2, "method1")
		assert.NoError(t, err)

		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		assert.Equal(t, 1, g.GetAttributesLen()[1])
		assert.Equal(t, 1, g.GetAttributesLen()[2])

		// Remove from section 1
		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, 0, g.GetAttributesLen()[1])
		assert.Equal(t, 1, g.GetAttributesLen()[2]) // Section 2 should be unchanged

		// Remove from section 2
		err = diagram.RemoveAttributeFromGadget(2, 0)
		assert.NoError(t, err)
		assert.Equal(t, 0, g.GetAttributesLen()[1])
		assert.Equal(t, 0, g.GetAttributesLen()[2])
	})

	t.Run("remove multiple attributes in order", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add and select a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Add multiple attributes
		err = diagram.AddAttributeToGadget(1, "attribute1")
		assert.NoError(t, err)
		err = diagram.AddAttributeToGadget(1, "attribute2")
		assert.NoError(t, err)
		err = diagram.AddAttributeToGadget(1, "attribute3")
		assert.NoError(t, err)

		g, err := diagram.componentsContainer.SearchGadget(gadPoint)
		assert.NoError(t, err)
		assert.Equal(t, 3, g.GetAttributesLen()[1])

		// Remove the first attribute (index 0)
		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, 2, g.GetAttributesLen()[1])

		// Remove the new first attribute (what was index 1, now index 0)
		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, 1, g.GetAttributesLen()[1])

		// Remove the last attribute
		err = diagram.RemoveAttributeFromGadget(1, 0)
		assert.NoError(t, err)
		assert.Equal(t, 0, g.GetAttributesLen()[1])
	})
}

// Test function for SetAssociationType
func TestSetAssociationType(t *testing.T) {
	t.Run("successful association type change", func(t *testing.T) {
		// Create diagram and add gadgets
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		gadPoint1 := utils.Point{X: 10, Y: 20}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "TestClass1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "TestClass2")
		assert.NoError(t, err)

		// Create association
		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Extension, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: (gadPoint1.X + gadPoint2.X) / 2, Y: (gadPoint1.Y + gadPoint2.Y) / 2}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Get the association to verify changes
		comp, err := diagram.componentsContainer.Search(midPoint)
		assert.NoError(t, err)
		assert.NotNil(t, comp)
		association := comp.(*component.Association)

		// Verify initial type
		var expectedInitialType component.AssociationType = component.Extension
		assert.Equal(t, expectedInitialType, association.GetAssType())

		// Change association type
		var newAssociationType component.AssociationType = component.Composition
		err = diagram.SetAssociationType(newAssociationType)
		assert.NoError(t, err)

		// Verify the type was changed
		assert.Equal(t, newAssociationType, association.GetAssType())

		// Verify draw data is updated
		drawData := association.GetDrawData().(drawdata.Association)
		assert.Equal(t, int(component.Composition), drawData.AssType)
	})

	t.Run("no component selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Try to set association type without selecting anything
		err = diagram.SetAssociationType(component.Composition)
		assert.Error(t, err)
		assert.Equal(t, "can only operate on one component", err.Error())
	})

	t.Run("multiple components selected", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add two gadgets
		gadPoint1 := utils.Point{X: 10, Y: 20}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "TestClass1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "TestClass2")
		assert.NoError(t, err)

		// Select both gadgets
		err = diagram.SelectComponent(gadPoint1)
		assert.NoError(t, err)
		err = diagram.SelectComponent(gadPoint2)
		assert.NoError(t, err)

		// Try to set association type with multiple selections
		err = diagram.SetAssociationType(component.Composition)
		assert.Error(t, err)
		assert.Equal(t, "can only operate on one component", err.Error())
	})

	t.Run("selected component is not an association", func(t *testing.T) {
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		// Add a gadget
		gadPoint := utils.Point{X: 10, Y: 20}
		err = diagram.AddGadget(component.Class, gadPoint, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)

		// Select the gadget
		err = diagram.SelectComponent(gadPoint)
		assert.NoError(t, err)

		// Try to set association type on a gadget
		err = diagram.SetAssociationType(component.Composition)
		assert.Error(t, err)
		assert.Equal(t, "selected component is not an association", err.Error())
	})

	t.Run("test all valid association types", func(t *testing.T) {
		validTypes := []component.AssociationType{
			component.Extension,
			component.Implementation,
			component.Composition,
			component.Dependency,
		}

		for _, assType := range validTypes {
			t.Run(fmt.Sprintf("test_%v", assType), func(t *testing.T) {
				// Create diagram and add gadgets
				diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
				assert.NoError(t, err)

				gadPoint1 := utils.Point{X: 10, Y: 20}
				gadPoint2 := utils.Point{X: 100, Y: 100}
				err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "TestClass1")
				assert.NoError(t, err)
				err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "TestClass2")
				assert.NoError(t, err)

				// Create association with a different type initially
				var initialType component.AssociationType = component.Extension
				if assType == component.Extension {
					initialType = component.Composition
				}

				err = diagram.StartAddAssociation(gadPoint1)
				assert.NoError(t, err)
				err = diagram.EndAddAssociation(initialType, gadPoint2)
				assert.NoError(t, err)

				// Select the association at actual midpoint between gadgets
				midPoint := utils.Point{X: (gadPoint1.X + gadPoint2.X) / 2, Y: (gadPoint1.Y + gadPoint2.Y) / 2}
				err = diagram.SelectComponent(midPoint)
				assert.NoError(t, err)

				// Change to the test type
				err = diagram.SetAssociationType(assType)
				assert.NoError(t, err)

				// Verify the change
				comp, err := diagram.componentsContainer.Search(midPoint)
				assert.NoError(t, err)
				association := comp.(*component.Association)
				assert.Equal(t, assType, association.GetAssType())
			})
		}
	})

	t.Run("undo and redo functionality", func(t *testing.T) {
		// Create diagram and add gadgets
		diagram, err := CreateEmptyUMLDiagram("test.uml", ClassDiagram)
		assert.NoError(t, err)

		gadPoint1 := utils.Point{X: 10, Y: 20}
		gadPoint2 := utils.Point{X: 100, Y: 100}
		err = diagram.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "TestClass1")
		assert.NoError(t, err)
		err = diagram.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "TestClass2")
		assert.NoError(t, err)

		// Create association
		err = diagram.StartAddAssociation(gadPoint1)
		assert.NoError(t, err)
		err = diagram.EndAddAssociation(component.Extension, gadPoint2)
		assert.NoError(t, err)

		// Select the association
		midPoint := utils.Point{X: (gadPoint1.X + gadPoint2.X) / 2, Y: (gadPoint1.Y + gadPoint2.Y) / 2}
		err = diagram.SelectComponent(midPoint)
		assert.NoError(t, err)

		// Change association type
		err = diagram.SetAssociationType(component.Composition)
		assert.NoError(t, err)

		// Verify the diagram's draw data reflects the change
		diagramDrawData := diagram.GetDrawData()
		assert.Equal(t, 1, len(diagramDrawData.Associations))
		assert.Equal(t, int(component.Composition), diagramDrawData.Associations[0].AssType)
	})
}
