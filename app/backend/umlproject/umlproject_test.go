package umlproject

import (
	"context"
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

func TestEndAddAssociationTODO(t *testing.T) {
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
	err = p.AddGadget(component.Class, utils.Point{X: 110, Y: 120}, 2, drawdata.DefaultGadgetColor, "header2")
	assert.NoError(t, err)
	err = p.StartAddAssociation(utils.Point{X: 11, Y: 21})
	assert.NoError(t, err)
	err = p.EndAddAssociation(component.Extension, utils.Point{X: 111, Y: 121})
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

func TestDumlEndToEnd(t *testing.T) {
	// 建立新專案與圖
	p, err := CreateEmptyUMLProject("DumlTestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "DumlDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("DumlDiagram")
	assert.NoError(t, err)

	// 新增 Gadget
	err = p.AddGadget(component.Class, utils.Point{X: 20, Y: 30}, 1, drawdata.DefaultGadgetColor, "header1")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 120, Y: 130}, 2, drawdata.DefaultGadgetColor, "header2")
	assert.NoError(t, err)

	// 新增 Association
	err = p.StartAddAssociation(utils.Point{X: 20, Y: 30})
	assert.NoError(t, err)
	err = p.EndAddAssociation(component.Extension, utils.Point{X: 120, Y: 130})
	assert.NoError(t, err)

	// 儲存專案
	tmpFile, err := os.CreateTemp("", "duml_end_to_end_*.json")
	assert.NoError(t, err)
	defer func(name string) {
		_ = os.Remove(name)
	}(tmpFile.Name())
	assert.NoError(t, tmpFile.Close())
	err = p.SaveProject(tmpFile.Name())
	assert.NoError(t, err)

	// 載入專案
	p2, err := CreateEmptyUMLProject("DumlTestProject2")
	assert.NoError(t, err)
	err = p2.LoadProject(tmpFile.Name())
	assert.NoError(t, err)
	err = p2.SelectDiagram("DumlDiagram")
	assert.NoError(t, err)

	// 驗證 Gadget 與 Association 內容
	data1 := p.GetDrawData()
	data2 := p2.GetDrawData()
	// 比對 Gadget
	for _, g1 := range data1.Gadgets {
		found := false
		for _, g2 := range data2.Gadgets {
			if g1.GadgetType == g2.GadgetType && g1.X == g2.X && g1.Y == g2.Y {
				found = true
				break
			}
		}
		assert.True(t, found, "Gadget %+v not found after reload", g1)
	}

	// 比對 Association
	for _, a1 := range data1.Associations {
		found := false
		for _, a2 := range data2.Associations {
			if a1.AssType == a2.AssType && a1.StartX == a2.StartX && a1.EndX == a2.EndX {
				found = true
				break
			}
		}
		assert.True(t, found, "Association %+v not found after reload", a1)
	}
}

// Test SetPointComponent method
func TestSetPointComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget and select it
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)

	// Test setting point
	newPoint := utils.Point{X: 50, Y: 60}
	err = p.SetPointComponent(newPoint)
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetPointComponent(newPoint)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetLayerComponent method
func TestSetLayerComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget and select it
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)

	// Test setting layer
	err = p.SetLayerComponent(5)
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetLayerComponent(3)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetColorComponent method
func TestSetColorComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget and select it
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)

	// Test setting color
	err = p.SetColorComponent("#FF0000")
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetColorComponent("#00FF00")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetAttrContentComponent method
func TestSetAttrContentComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget, select it, and add attribute
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
	err = p.AddAttributeToGadget(1, "test attribute")
	assert.NoError(t, err)

	// Test setting attribute content
	err = p.SetAttrContentComponent(1, 0, "modified attribute")
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetAttrContentComponent(1, 0, "another content")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetAttrSizeComponent method
func TestSetAttrSizeComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget, select it, and add attribute
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
	err = p.AddAttributeToGadget(1, "test attribute")
	assert.NoError(t, err)

	// Test setting attribute size
	err = p.SetAttrSizeComponent(1, 0, 16)
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetAttrSizeComponent(1, 0, 18)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetAttrStyleComponent method
func TestSetAttrStyleComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget, select it, and add attribute
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
	err = p.AddAttributeToGadget(1, "test attribute")
	assert.NoError(t, err)

	// Test setting attribute style
	err = p.SetAttrStyleComponent(1, 0, attribute.Bold)
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetAttrStyleComponent(1, 0, attribute.Italic)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetAttrFontComponent method
func TestSetAttrFontComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget, select it, and add attribute
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
	err = p.AddAttributeToGadget(1, "test attribute")
	assert.NoError(t, err)

	// Test setting attribute font
	err = p.SetAttrFontComponent(1, 0, "Arial")
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetAttrFontComponent(1, 0, "Times")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetAttrRatioComponent method
func TestSetAttrRatioComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add two gadgets to create an association
	gadPoint1 := utils.Point{X: 10, Y: 10}
	gadPoint2 := utils.Point{X: 100, Y: 100}
	err = p.AddGadget(component.Class, gadPoint1, 0, drawdata.DefaultGadgetColor, "test gadget1")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, gadPoint2, 0, drawdata.DefaultGadgetColor, "test gadget2")
	assert.NoError(t, err)

	// Create association
	err = p.StartAddAssociation(gadPoint1)
	assert.NoError(t, err)
	err = p.EndAddAssociation(component.Composition, gadPoint2)
	assert.NoError(t, err)

	// Select the association and add attribute
	midPoint := utils.Point{X: (gadPoint1.X + gadPoint2.X) / 2, Y: (gadPoint1.Y + gadPoint2.Y) / 2}
	err = p.SelectComponent(midPoint)
	assert.NoError(t, err)
	err = p.AddAttributeToAssociation(0.5, "test attribute")
	assert.NoError(t, err)

	// Test setting attribute ratio
	err = p.SetAttrRatioComponent(0, 0, 0.75)
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetAttrRatioComponent(0, 0, 0.5)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetParentStartComponent method
func TestSetParentStartComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add gadgets and create association
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "gadget1")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "gadget2")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 200, Y: 200}, 0, drawdata.DefaultGadgetColor, "gadget3")
	assert.NoError(t, err)

	err = p.StartAddAssociation(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
	err = p.EndAddAssociation(component.Composition, utils.Point{X: 105, Y: 105})
	assert.NoError(t, err)

	// Select the association and change its start parent
	err = p.SelectComponent(utils.Point{X: 55, Y: 55})
	assert.NoError(t, err)
	err = p.SetParentStartComponent(utils.Point{X: 205, Y: 205})
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetParentStartComponent(utils.Point{X: 15, Y: 15})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetParentEndComponent method
func TestSetParentEndComponent(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add gadgets and create association
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "gadget1")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "gadget2")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 200, Y: 200}, 0, drawdata.DefaultGadgetColor, "gadget3")
	assert.NoError(t, err)

	err = p.StartAddAssociation(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
	err = p.EndAddAssociation(component.Composition, utils.Point{X: 105, Y: 105})
	assert.NoError(t, err)

	// Select the association and change its end parent
	err = p.SelectComponent(utils.Point{X: 55, Y: 55})
	assert.NoError(t, err)
	err = p.SetParentEndComponent(utils.Point{X: 205, Y: 205})
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetParentEndComponent(utils.Point{X: 105, Y: 105})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test SetAssociationType method
func TestSetAssociationType(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add gadgets and create association
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "gadget1")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "gadget2")
	assert.NoError(t, err)

	err = p.StartAddAssociation(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
	err = p.EndAddAssociation(component.Composition, utils.Point{X: 105, Y: 105})
	assert.NoError(t, err)

	// Select the association and change its type
	err = p.SelectComponent(utils.Point{X: 55, Y: 55})
	assert.NoError(t, err)
	err = p.SetAssociationType(component.Extension)
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SetAssociationType(component.Dependency)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test UndoDiagramChange method
func TestUndoDiagramChange(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget to have something to undo
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)

	// Test undo
	err = p.UndoDiagramChange()
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.UndoDiagramChange()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test RedoDiagramChange method
func TestRedoDiagramChange(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget and undo to have something to redo
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)
	err = p.UndoDiagramChange()
	assert.NoError(t, err)

	// Test redo
	err = p.RedoDiagramChange()
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.RedoDiagramChange()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test RemoveAttributeFromGadget method
func TestRemoveAttributeFromGadget(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add a gadget, select it, and add attributes
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test gadget")
	assert.NoError(t, err)
	err = p.SelectComponent(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
	err = p.AddAttributeToGadget(1, "attribute1")
	assert.NoError(t, err)
	err = p.AddAttributeToGadget(1, "attribute2")
	assert.NoError(t, err)

	// Test removing attribute
	err = p.RemoveAttributeFromGadget(1, 0)
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.RemoveAttributeFromGadget(1, 0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test RemoveAttributeFromAssociation method
func TestRemoveAttributeFromAssociation(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add gadgets and create association
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "gadget1")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "gadget2")
	assert.NoError(t, err)

	err = p.StartAddAssociation(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
	err = p.EndAddAssociation(component.Composition, utils.Point{X: 105, Y: 105})
	assert.NoError(t, err)

	// Select association and add attributes
	err = p.SelectComponent(utils.Point{X: 55, Y: 55})
	assert.NoError(t, err)
	err = p.AddAttributeToAssociation(0.3, "attr1")
	assert.NoError(t, err)
	err = p.AddAttributeToAssociation(0.7, "attr2")
	assert.NoError(t, err)

	// Test removing attribute
	err = p.RemoveAttributeFromAssociation(0)
	assert.NoError(t, err)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.RemoveAttributeFromAssociation(0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test proper EndAddAssociation implementation
func TestEndAddAssociationImplementation(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Add gadgets
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "gadget1")
	assert.NoError(t, err)
	err = p.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "gadget2")
	assert.NoError(t, err)

	// Start association
	err = p.StartAddAssociation(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)

	// End association successfully
	err = p.EndAddAssociation(component.Composition, utils.Point{X: 105, Y: 105})
	assert.NoError(t, err)

	// Verify association was created
	drawData := p.GetDrawData()
	assert.Len(t, drawData.Associations, 1)
	assert.Equal(t, component.Composition, drawData.Associations[0].AssType)

	// Test with no diagram selected
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.EndAddAssociation(component.Extension, utils.Point{X: 50, Y: 50})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test DeleteDiagram method (currently TODO)
func TestDeleteDiagramTODO(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)

	// Currently this method just returns nil as it's not implemented
	err = p.DeleteDiagram("TestDiagram")
	assert.NoError(t, err)
}

// Test InvalidateCanvas method more thoroughly
func TestInvalidateCanvasDetailed(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Test without context - should not error but do nothing
	err = p.InvalidateCanvas()
	assert.NoError(t, err)

	// Test with runFrontend false - should return early
	p.runFrontend = false
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.InvalidateCanvas()
	assert.NoError(t, err)

	// Test with no current diagram selected
	p.runFrontend = true
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.InvalidateCanvas()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")
}

// Test Startup method
func TestStartupDetailed(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	ctx := context.Background()
	p.Startup(ctx)

	// Verify startup effects
	assert.Equal(t, ctx, p.ctx)
	assert.True(t, p.runFrontend)
	assert.Contains(t, p.GetAvailableDiagramsNames(), "new class diagram")
	assert.Contains(t, p.GetActiveDiagramsNames(), "new class diagram")
	assert.Equal(t, "new class diagram", p.GetCurrentDiagramName())
}

// Test OpenDiagram error cases
func TestOpenDiagramErrors(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Test with invalid file path
	err = p.OpenDiagram("")
	assert.Error(t, err)

	// Test with non-existent file
	err = p.OpenDiagram("non_existent_file.json5")
	assert.Error(t, err)

	// Test with invalid JSON content
	tmpFile, err := os.CreateTemp("", "invalid_*.json5")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Write invalid JSON
	tmpFile.WriteString("invalid json content")
	tmpFile.Close()

	err = p.OpenDiagram(tmpFile.Name())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to decode file")
}

// Test SaveDiagram error cases
func TestSaveDiagramErrors(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Test with no current diagram
	err = p.SaveDiagram("test.json")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No current diagram selected")

	// Create a diagram to test other error cases
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Test with invalid file path (directory that doesn't exist)
	err = p.SaveDiagram("/non_existent_directory/test.json")
	assert.Error(t, err)
}

// Test LoadProject error cases
func TestLoadProjectErrors(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Test with invalid file path
	err = p.LoadProject("")
	assert.Error(t, err)

	// Test with non-existent file
	err = p.LoadProject("non_existent_project.json5")
	assert.Error(t, err)

	// Test with invalid JSON content
	tmpFile, err := os.CreateTemp("", "invalid_project_*.json5")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Write invalid JSON
	tmpFile.WriteString("invalid json content")
	tmpFile.Close()

	err = p.LoadProject(tmpFile.Name())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to decode project file")
}

// Test SaveProject error cases
func TestSaveProjectErrors(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Test with invalid file path (when filename differs from project name)
	err = p.SaveProject("")
	assert.Error(t, err)

	// Test with directory that doesn't exist
	err = p.SaveProject("/non_existent_directory/project.json")
	assert.Error(t, err)
}

// Test CloseProject when no save is needed
func TestCloseProjectNoSave(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Set lastSave to be after lastModified (no save needed)
	p.lastSave = time.Now().Add(time.Hour)
	p.lastModified = time.Now()

	err = p.CloseProject()
	assert.NoError(t, err)
}

// Test various diagram types
func TestCreateDiagramDifferentTypes(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Test creating class diagram (supported)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "ClassDiagram")
	assert.NoError(t, err)

	// Test creating use case diagram (not supported - should fail)
	err = p.CreateEmptyUMLDiagram(umldiagram.UseCaseDiagram, "UseCaseDiagram")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid diagram type")

	// Verify only supported diagram exists
	diagrams := p.GetAvailableDiagramsNames()
	assert.Contains(t, diagrams, "ClassDiagram")
	assert.NotContains(t, diagrams, "UseCaseDiagram")
}

// Test file dialog methods without context
func TestFileDialogsWithoutContext(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Test OpenFileDialog without context
	_, err = p.OpenFileDialog()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "application context not available")

	// Test SaveFileDialog without context
	_, err = p.SaveFileDialog()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "application context not available")

	// Test SaveDiagramFileDialog without context
	_, err = p.SaveDiagramFileDialog()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "application context not available")
}

// Test edge cases for component selection and modification
func TestComponentSelectionEdgeCases(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	// Test selecting component when no components exist - should not error but not select anything
	err = p.SelectComponent(utils.Point{X: 50, Y: 50})
	assert.NoError(t, err) // No error but nothing selected

	// Add a component and test selecting outside its bounds - should not error
	err = p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test")
	assert.NoError(t, err)

	err = p.SelectComponent(utils.Point{X: 500, Y: 500})
	assert.NoError(t, err) // No error but nothing selected

	// Test that we can successfully select the component at its correct location
	err = p.SelectComponent(utils.Point{X: 15, Y: 15})
	assert.NoError(t, err)
}

// Test multiple diagram management
func TestMultipleDiagramManagement(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)

	// Create multiple diagrams
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "Diagram1")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "Diagram2")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "Diagram3")
	assert.NoError(t, err)

	// Select different diagrams
	err = p.SelectDiagram("Diagram1")
	assert.NoError(t, err)
	assert.Equal(t, "Diagram1", p.GetCurrentDiagramName())

	err = p.SelectDiagram("Diagram2")
	assert.NoError(t, err)
	assert.Equal(t, "Diagram2", p.GetCurrentDiagramName())

	// Close one diagram
	err = p.CloseDiagram("Diagram1")
	assert.NoError(t, err)

	activeDiagrams := p.GetActiveDiagramsNames()
	assert.NotContains(t, activeDiagrams, "Diagram1")
	assert.Contains(t, activeDiagrams, "Diagram2")
	assert.Contains(t, activeDiagrams, "Diagram3")

	// Available diagrams should still contain all
	availableDiagrams := p.GetAvailableDiagramsNames()
	assert.Contains(t, availableDiagrams, "Diagram1")
	assert.Contains(t, availableDiagrams, "Diagram2")
	assert.Contains(t, availableDiagrams, "Diagram3")
}

// Test lastModified updates
func TestLastModifiedUpdates(t *testing.T) {
	p, err := CreateEmptyUMLProject("TestProject")
	assert.NoError(t, err)
	err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)

	initialTime := p.GetLastModified()

	// Sleep to ensure time difference
	time.Sleep(10 * time.Millisecond)

	// Operations that should update lastModified
	operations := []func() error{
		func() error {
			return p.AddGadget(component.Class, utils.Point{X: 10, Y: 10}, 0, drawdata.DefaultGadgetColor, "test")
		},
		func() error { return p.SelectComponent(utils.Point{X: 15, Y: 15}) },
		func() error { return p.SetLayerComponent(1) },
		func() error { return p.SetColorComponent("#FF0000") },
	}

	for i, operation := range operations {
		time.Sleep(10 * time.Millisecond) // Ensure time difference
		err := operation()
		assert.NoError(t, err, "Operation %d failed", i)

		newTime := p.GetLastModified()
		assert.True(t, newTime.After(initialTime), "LastModified should be updated after operation %d", i)
		initialTime = newTime
	}
}
