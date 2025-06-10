package umlproject

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"Dr.uml/backend/component/attribute"
	"Dr.uml/backend/drawdata"

	"Dr.uml/backend/component"
	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewUMLProject(t *testing.T) {
	// valid name
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	assert.Equal(t, "TestProject", p.GetName())

	// invalid name
	_, err = CreateEmptyUMLProject("")
	assert.Error(t, err)
}

func TestGetName(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	name := p.GetName()
	assert.Equal(t, "TestProject", name)
}

func TestGetLastModified(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	lastModified := p.GetLastModified()
	assert.WithinDuration(t, time.Now(), lastModified, time.Second)
}

func TestGetCurrentDiagramName(t *testing.T) {
	// no selected diagram
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	name := p.GetCurrentDiagramName()
	assert.Equal(t, "", name)

	// selected diagram
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)
	name = p.GetCurrentDiagramName()
	assert.Equal(t, "TestDiagram", name)
}

func TestGetAvailableDiagramsNames(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// no diagrams
	names := p.GetAvailableDiagramsNames()
	assert.Empty(t, names)

	// one diagram
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram1")
	assert.NoError(t, err)
	names = p.GetAvailableDiagramsNames()
	assert.Equal(t, []string{"TestDiagram1"}, names)

	// two diagrams
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram2")
	assert.NoError(t, err)
	names = p.GetAvailableDiagramsNames()
	assert.ElementsMatch(t, []string{"TestDiagram1", "TestDiagram2"}, names)
}

func TestGetActiveDiagramsNames(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// no active diagrams
	names := p.GetActiveDiagramsNames()
	assert.Empty(t, names)

	// one active diagram
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram1")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram1")
	assert.NoError(t, err)
	names = p.GetActiveDiagramsNames()
	assert.Equal(t, []string{"TestDiagram1"}, names)

	// two active diagrams
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram2")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram2")
	assert.NoError(t, err)
	names = p.GetActiveDiagramsNames()
	assert.ElementsMatch(t, []string{"TestDiagram1", "TestDiagram2"}, names)
}

func TestSelectDiagram(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)

	// success, select an active diagram
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)
	name := p.GetCurrentDiagramName()
	assert.Equal(t, "TestDiagram", name)

	// success, select a non-active diagram

	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// diagram not exist
	err = p.SelectDiagram("NonExistentDiagram")
	assert.Error(t, err)

	// TODO: error when LoadExistUMLDiagram
}

func TestCreateEmptyUMLDiagram(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// create success
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	names := p.GetAvailableDiagramsNames()
	assert.Equal(t, []string{"TestDiagram"}, names)

	// duplicate name
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.Error(t, err)

	// invalid diagram type
	err = p.CreateEmptyUMLDiagram(umldiagram.DiagramType(100), "TestDiagram2")
	assert.Error(t, err)

	// invalid name
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "")
	assert.Error(t, err)
}

func TestCloseDiagram(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)

	// success, close a non-selected diagram
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	names := p.GetAvailableDiagramsNames()
	assert.Equal(t, []string{"TestDiagram"}, names)
	names = p.GetActiveDiagramsNames()
	assert.Empty(t, names)

	// success, close a selected diagram
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	name := p.GetCurrentDiagramName()
	assert.Equal(t, "", name)
	names = p.GetActiveDiagramsNames()
	assert.Empty(t, names)

	// close a non-active diagram
	err = p.CloseDiagram("TestDiagram")
	assert.Error(t, err)
}

func TestDeleteDiagram(t *testing.T) {
	// TODO
}

func TestAddGadget(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// success
	err = p.AddGadget(component.Class, utils.Point{X: 0, Y: 0}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.NoError(t, err)

	// invalid gadget type
	err = p.AddGadget(component.GadgetType(3), utils.Point{X: 0, Y: 0}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.Error(t, err)

	// no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 0, Y: 0}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.Error(t, err)
}

func TestStartAddAssociation(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// success
	err = p.StartAddAssociation(utils.Point{X: 0, Y: 0})
	assert.NoError(t, err)

	// invalid point
	err = p.StartAddAssociation(utils.Point{X: -1, Y: 0})
	assert.Error(t, err)

	// no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.StartAddAssociation(utils.Point{X: 0, Y: 0})
	assert.Error(t, err)

}

func TestEndAddAssociation(t *testing.T) {
	// TODO
}

func TestRemoveSelectedComponents(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add and select a component
	err = p.AddGadget(component.Class, utils.Point{X: 0, Y: 0}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 5, Y: 5})
	assert.NoError(t, err)

	// Remove selected components
	err = p.RemoveSelectedComponents()
	assert.NoError(t, err)

	// No diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.RemoveSelectedComponents()
	assert.Error(t, err)
}

func TestGetDrawData(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 0, Y: 0}, 0, drawdata.DefaultGadgetColor, "sample header")

	assert.NoError(t, err)
	// success
	data := p.GetDrawData()
	assert.Len(t, data.Gadgets, 1)

	// no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	data = p.GetDrawData()
	assert.Empty(t, data)
}

func TestAddAttributeToGadget(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget
	err = p.AddGadget(component.Class, utils.Point{X: 0, Y: 0}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.NoError(t, err)

	// Select the gadget
	err = p.SelectComponent(utils.Point{X: 5, Y: 5})
	assert.NoError(t, err)

	// Add attribute to the gadget
	err = p.AddAttributeToGadget(1, "newAttribute")
	assert.NoError(t, err)

	// No diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.AddAttributeToGadget(1, "newAttribute")
	assert.Error(t, err)
}

func TestAddAttributeToAssociation(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)
	// Start adding an association
	err = p.AddGadget(component.Class, utils.Point{X: 0, Y: 0}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 200, Y: 200}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.NoError(t, err)
	// Start adding association
	err = p.StartAddAssociation(utils.Point{X: 1, Y: 1})
	assert.NoError(t, err)
	// End adding association
	err = p.EndAddAssociation(component.Composition, utils.Point{X: 201, Y: 201})
	assert.NoError(t, err)
	// Add attribute to the association
	err = p.SelectComponent(utils.Point{X: 50, Y: 50})
	assert.NoError(t, err)
	err = p.AddAttributeToAssociation(0.5, "newAttribute")
	assert.NoError(t, err)
	err = p.SetAttrContentComponent(0, 0, "modifiedAttribute")
	assert.NoError(t, err)
	assert.Equal(t, "modifiedAttribute", p.GetDrawData().Associations[0].Attributes[0].Content)
	err = p.SetAttrSizeComponent(0, 0, 20)
	assert.NoError(t, err)
	assert.Equal(t, 20, p.GetDrawData().Associations[0].Attributes[0].FontSize)
	err = p.SetAttrStyleComponent(0, 0, attribute.Italic)
	assert.NoError(t, err)
	assert.Equal(t, attribute.Italic, p.GetDrawData().Associations[0].Attributes[0].FontStyle)
	err = p.SetAttrFontComponent(0, 0, "Inkfree")
	assert.NoError(t, err)
	assert.Equal(t, "Inkfree", p.GetDrawData().Associations[0].Attributes[0].FontFile)
	err = p.SetAttrRatioComponent(0, 0, 0.5)
	assert.NoError(t, err)
	assert.Equal(t, 0.5, p.GetDrawData().Associations[0].Attributes[0].Ratio)

}

func TestSelectComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget
	err = p.AddGadget(component.Class, utils.Point{X: 0, Y: 0}, 0, drawdata.DefaultGadgetColor, "sample header")
	assert.NoError(t, err)

	// Select the component
	err = p.SelectComponent(utils.Point{X: 5, Y: 5})
	assert.NoError(t, err)

	// No diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 5, Y: 5})
	assert.Error(t, err)
}

func TestStartup(t *testing.T) {
	// TODO
}

func TestInvalidateCanvas(t *testing.T) {
	// TODO
}

func TestOpenDiagram(t *testing.T) {
	proj, err := CreateEmptyUMLProject("fuck")
	assert.NoError(t, err)
	root, ok := os.LookupEnv("APP_ROOT")
	assert.True(t, ok)
	err = proj.OpenDiagram(root + "/backend/example.json5")
	assert.NoError(t, err)
}

func TestSaveDiagram(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	root, ok := os.LookupEnv("APP_ROOT")
	assert.True(t, ok)
	err = p.OpenDiagram(root + "/backend/example.json5")
	assert.NoError(t, err)

	diagramName := "umlproject_save_test_.json"

	// Save to a temp file
	tmpFile, err := os.Create(diagramName)
	assert.NoError(t, err)
	// defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	err = p.SaveDiagram(tmpFile.Name())
	assert.NoError(t, err)

	// No diagram selected
	err = p.CloseDiagram(diagramName)
	assert.NoError(t, err)
	err = p.SaveDiagram(tmpFile.Name())
	assert.Error(t, err)
}

func TestLoadProject(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	filename := os.Getenv("APP_ROOT") + "/backend/example_proj.json5"
	err = p.LoadProject(filename)
	assert.NoError(t, err)
	assert.Equal(t, filename, p.GetName())
	assert.Contains(t, p.GetAvailableDiagramsNames(), "path/to/diagram1.json5")
}

func TestSaveProject(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Create and select a diagram
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Save to a temp file
	tmpFile, err := os.CreateTemp("", "umlproject_save_test_*.json")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err)
	}(tmpFile.Name())
	assert.NoError(t, tmpFile.Close())

	// Save project
	err = p.SaveProject(tmpFile.Name())
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err)
	}(os.Getenv("APP_ROOT") + "/backend/umlproject/TestDiagram")

	// Try saving with an invalid file path
	err = p.SaveProject("")
	assert.Error(t, err)
}

func TestProjectSaveLoadContentEquality(t *testing.T) {
	// Create a new project and diagram
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget for some content
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 20}, 1, drawdata.DefaultGadgetColor, "header")
	assert.NoError(t, err)

	// Save project to a temp file
	tmpFile, err := os.CreateTemp("", "umlproject_content_test_*.json")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err)
	}(tmpFile.Name())

	err = tmpFile.Close()
	assert.NoError(t, err)

	err = p.SaveProject(tmpFile.Name())
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err)
	}(os.Getenv("APP_ROOT") + "/backend/umlproject/TestDiagram")

	// Read saved file content
	savedContent, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)

	// Load project from file
	p2, err := CreateEmptyUMLProject("AnotherProject")
	assert.NoError(t, err)
	err = p2.LoadProject(tmpFile.Name())
	assert.NoError(t, err)

	// Save loaded project to another temp file
	tmpFile2, err := os.Create("umlproject_content_test2.json")
	assert.NoError(t, err)
	// defer os.Remove(tmpFile2.Name())
	err = tmpFile2.Close()
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err)
	}(tmpFile2.Name())

	err = p2.SaveProject(tmpFile2.Name())
	assert.NoError(t, err)

	// Read loaded file content
	loadedContent, err := os.ReadFile(tmpFile2.Name())
	assert.NoError(t, err)

	// Assert the content is identical
	assert.Equal(t, string(savedContent), string(loadedContent))
}

func TestCloseProject(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Create and select a diagram to modify the project
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Save the project to set lastSave
	tmpFile, err := os.CreateTemp("", "umlproject_close_test_*.json")
	assert.NoError(t, err)
	defer func(name string) {
		_ = os.Remove(name)
	}(tmpFile.Name())
	assert.NoError(t, tmpFile.Close())

	err = p.SaveProject(tmpFile.Name())
	assert.NoError(t, err)
	err = os.Remove(tmpFile.Name())
	assert.NoError(t, err)

	time.Sleep(10 * time.Millisecond) // Ensure lastModified is later than lastSave
	// Modify the project
	err = p.AddGadget(component.Class, utils.Point{X: 1, Y: 1}, 0, drawdata.DefaultGadgetColor, "header")
	assert.NoError(t, err)

	// Call CloseProject, should trigger save
	err = p.CloseProject()
	assert.FileExists(t, tmpFile.Name())
	content, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)
	expectedContent := `{
		"diagrams": [
			"TestDiagram"
		]
	}`
	var expectedData, actualData map[string]interface{}
	assert.NoError(t, json.Unmarshal([]byte(expectedContent), &expectedData), "Failed to unmarshal expected JSON")
	assert.NoError(t, json.Unmarshal(content, &actualData), "Failed to unmarshal actual JSON")
	assert.Equal(t, expectedData, actualData, "Project content should match after close")
	assert.NoError(t, err)
}
