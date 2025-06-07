// VIBE CODING

package umldiagram

import (
	"Dr.uml/backend/component/attribute"
	"encoding/json"
	"os"
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

func TestLoadExistUMLDiagram(t *testing.T) {

}

func TestLoadGadgetAttributes(t *testing.T) {
	dia, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, dia)

	expectedContent := "test content"
	expectedSize := 12
	expectedStyle := attribute.Textstyle(attribute.Bold | attribute.Italic)
	expectedFontFile := os.Getenv("APP_ROOT") + "/assets/Inkfree.ttf"

	savedAttributeBase := utils.SavedAtt{
		Content:  expectedContent,
		Size:     expectedSize,
		Style:    int(expectedStyle),
		FontFile: expectedFontFile,
	}
	savedAttributes := make([]utils.SavedAtt, 3)
	for i := 0; i < 3; i++ {
		savedAttributes[i] = savedAttributeBase
		savedAttributes[i].Size += i
		savedAttributes[i].Ratio = 0.3 * float64(i)
	}

	gad, err := component.NewGadget(component.Class, utils.Point{}, 0, "someInvalidColorHex", "")
	assert.NoError(t, err)
	assert.NotNil(t, gad)

	err, _ = dia.loadGadgetAttributes(gad, savedAttributes) // Err-index is not important cuz we expect err is nil
	assert.NoError(t, err)

	loadedAttributes := gad.GetAttributes()

	for i := 0; i < 3; i++ {
		assert.Equal(t, 1, len(loadedAttributes[i]))
		assert.Equal(t, savedAttributes[i].Content, loadedAttributes[i][0].GetContent())
		assert.Equal(t, savedAttributes[i].Style, int(loadedAttributes[i][0].GetStyle()))
		assert.Equal(t, savedAttributes[i].FontFile, loadedAttributes[i][0].GetFontFile())
		assert.Equal(t, savedAttributes[i].Size, loadedAttributes[i][0].GetSize())
		assert.Equal(t, savedAttributes[i].FontFile, loadedAttributes[i][0].GetFontFile())
	}
}
func TestLoadGadgetAttributesButWithJsonStr(t *testing.T) {
	dia, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, dia)

	jsonStr := `{
		"Content": "test content",
		"Size": 12,
		"Style": 3,
		"FontFile": "` + os.Getenv("APP_ROOT") + `/assets/Inkfree.ttf"
	}`

	var savedAttributeBase utils.SavedAtt
	err = json.Unmarshal([]byte(jsonStr), &savedAttributeBase)
	assert.NoError(t, err)

	savedAttributes := make([]utils.SavedAtt, 3)
	for i := 0; i < 3; i++ {
		savedAttributes[i] = savedAttributeBase
		savedAttributes[i].Size += i
		savedAttributes[i].Ratio = 0.3 * float64(i)
	}

	gad, err := component.NewGadget(component.Class, utils.Point{}, 0, "someInvalidColorHex", "")
	assert.NoError(t, err)
	assert.NotNil(t, gad)

	err, _ = dia.loadGadgetAttributes(gad, savedAttributes) // Err-index is not important cuz we expect err is nil
	assert.NoError(t, err)

	loadedAttributes := gad.GetAttributes()

	for i := 0; i < 3; i++ {
		assert.Equal(t, 1, len(loadedAttributes[i]))
		assert.Equal(t, savedAttributes[i].Content, loadedAttributes[i][0].GetContent())
		assert.Equal(t, savedAttributes[i].Style, int(loadedAttributes[i][0].GetStyle()))
		assert.Equal(t, savedAttributes[i].FontFile, loadedAttributes[i][0].GetFontFile())
		assert.Equal(t, savedAttributes[i].Size, loadedAttributes[i][0].GetSize())
		assert.Equal(t, savedAttributes[i].FontFile, loadedAttributes[i][0].GetFontFile())
	}
}
func TestLoadGadgets(t *testing.T) {
	dia, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, dia)

	expectedContent := "test content"
	expectedSize := 12
	expectedStyle := attribute.Textstyle(attribute.Bold | attribute.Italic)
	expectedFontFile := os.Getenv("APP_ROOT") + "/assets/Inkfree.ttf"

	savedAttributeBase := utils.SavedAtt{
		Content:  expectedContent,
		Size:     expectedSize,
		Style:    int(expectedStyle),
		FontFile: expectedFontFile,
	}

	savedAttributes := make([]utils.SavedAtt, 3)
	for i := 0; i < 3; i++ {
		savedAttributes[i] = savedAttributeBase
		savedAttributes[i].Size += i
		savedAttributes[i].Ratio = 0.3 * float64(i)
	}

	savedGadgetBase := utils.SavedGad{
		GadgetType: 1,
		Point:      "0, 0",
		Color:      "InvalidColorHex",
		Attributes: savedAttributes,
	}
	savedGadgets := make([]utils.SavedGad, 69)

	for i := 0; i < len(savedGadgets); i++ {
		savedGadgets[i] = savedGadgetBase
		savedGadgets[i].Layer = i
	}

	dp, err := dia.loadGadgets(savedGadgets)
	assert.NoError(t, err)
	assert.NotNil(t, dp)
	assert.Equal(t, len(savedGadgets), len(dp))
}

func TestUMLDiagram_LoadAsses(t *testing.T) {
	dia, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, dia)

	savedGadgetBase := utils.SavedGad{
		GadgetType: 1,
		Point:      "0, 0",
		Color:      "InvalidColorHex",
	}
	savedGadgets := make([]utils.SavedGad, 70)

	for i := 0; i < len(savedGadgets); i++ {
		savedGadgets[i] = savedGadgetBase
		savedGadgets[i].Layer = i
	}
	dp, err := dia.loadGadgets(savedGadgets)
	assert.NoError(t, err)

	expectedAssType := component.AssociationType(1)
	expectedLayer := 0
	expectedStartPoint := "10, 10"
	expectedEndPoint := "20, 20"

	savedAssBase := utils.SavedAss{
		AssType:    int(expectedAssType),
		Layer:      expectedLayer,
		StartPoint: expectedStartPoint,
		EndPoint:   expectedEndPoint,
	}
	savedAsses := make([]utils.SavedAss, 69)
	for i := 0; i < len(savedAsses); i++ {
		savedAsses[i] = savedAssBase
		savedAsses[i].Parents = []int{i, i + 1} // Assuming each association connects
	}
	err = dia.LoadAsses(savedAsses, dp)
	assert.NoError(t, err)
	// Check if associations are loaded correctly
	components := dia.componentsContainer.GetAll()
	for _, comp := range components {
		switch comp.(type) {
		case *component.Association:
			assert.Equal(t, expectedAssType, comp.(*component.Association).GetAssType())
			assert.Equal(t, expectedLayer, comp.(*component.Association).GetLayer())
		default:
			continue
		}
	}
}

func TestUMLDiagram_loadAssAttributes(t *testing.T) {
	dia, err := CreateEmptyUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, dia)

	var parents = [2]*component.Gadget{nil, nil}
	for i := 0; i < 2; i++ {
		parents[i], err = component.NewGadget(component.Class, utils.Point{X: 0, Y: 0}, 0, "InvalidColorHex", "")
		assert.NoError(t, err)
	}

	ass, err := component.NewAssociation(parents, component.AssociationType(1), utils.Point{X: 0, Y: 0}, utils.Point{X: 1, Y: 1})
	assert.NoError(t, err)
	assert.NotNil(t, ass)

	// Prepare attributes

	expectedContent := "test"
	expectedStyle := int(attribute.Bold)
	expectedFontFile := os.Getenv("APP_ROOT") + "/assets/Inkfree.ttf"
	expectedRatio := 0.69

	savedAttBase := utils.SavedAtt{
		Content:  expectedContent,
		Style:    expectedStyle,
		FontFile: expectedFontFile,
		Ratio:    expectedRatio,
	}

	attributes := make([]utils.SavedAtt, 2)
	for i := 0; i < len(attributes); i++ {
		attributes[i] = savedAttBase
		attributes[i].Size = i + 1
	}

	// Should succeed
	errRet, idx := dia.loadAssAttributes(ass, attributes)
	assert.NoError(t, errRet)
	assert.Equal(t, 0, idx)

	atts, err := ass.GetAttributes()
	assert.NoError(t, err)

	for i, att := range atts {
		assert.Equal(t, expectedContent, att.GetContent())
		assert.Equal(t, expectedStyle, int(att.GetStyle()))
		assert.Equal(t, expectedRatio, att.GetRatio())
		assert.Equal(t, expectedFontFile, att.GetFontFile())
		assert.Equal(t, i+1, att.GetSize())
	}

	// Should fail with nil association
	errRet, idx = dia.loadAssAttributes(nil, attributes)
	assert.Error(t, errRet)
	assert.Equal(t, 0, idx)
}

func TestUMLDiagram_SaveToFile(t *testing.T) {
	diagram, err := CreateEmptyUMLDiagram("SaveToFileTest.uml", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, diagram)

	// Add two gadgets
	err = diagram.AddGadget(component.Class, utils.Point{X: 1, Y: 2}, 0, drawdata.DefaultGadgetColor, "Header1")
	assert.NoError(t, err)
	err = diagram.AddGadget(component.Class, utils.Point{X: 3, Y: 4}, 1, drawdata.DefaultGadgetColor, "Header2")
	assert.NoError(t, err)

	// Add an association between the two gadgets
	gadgets := diagram.componentsContainer.GetAll()
	assert.Len(t, gadgets, 2)
	gad1, ok1 := gadgets[0].(*component.Gadget)
	gad2, ok2 := gadgets[1].(*component.Gadget)
	assert.True(t, ok1)
	assert.True(t, ok2)
	ass, err := component.NewAssociation([2]*component.Gadget{gad1, gad2}, component.AssociationType(1), utils.Point{X: 1, Y: 2}, utils.Point{X: 3, Y: 4})
	assert.NoError(t, err)
	err = diagram.componentsContainer.Insert(ass)
	assert.NoError(t, err)
	diagram.associations[gad1] = [2][]*component.Association{{ass}, {}}
	diagram.associations[gad2] = [2][]*component.Association{{}, {ass}}

	// Save to file
	_, err = diagram.SaveToFile("SaveToFileTest.uml")
	assert.NoError(t, err)
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
