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

func TestRemoveComponent(t *testing.T) {
	t.Run("remove single gadget", func(t *testing.T) {
		p, err := CreateEmptyUMLProject("TestProject")
		assert.NoError(t, err)
		err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
		assert.NoError(t, err)
		err = p.SelectDiagram("TestDiagram")
		assert.NoError(t, err)

		// Add a gadget
		err = p.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)

		// Verify the gadget exists
		data := p.GetDrawData()
		assert.Len(t, data.Gadgets, 1)
		assert.Equal(t, "TestClass", data.Gadgets[0].Attributes[0][0].Content)

		// Remove the gadget by clicking on it
		err = p.RemoveComponent(utils.Point{X: 100, Y: 100})
		assert.NoError(t, err)

		// Verify the gadget was removed
		data = p.GetDrawData()
		assert.Len(t, data.Gadgets, 0)
	})

	t.Run("remove gadget with associations", func(t *testing.T) {
		p, err := CreateEmptyUMLProject("TestProject")
		assert.NoError(t, err)
		err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
		assert.NoError(t, err)
		err = p.SelectDiagram("TestDiagram")
		assert.NoError(t, err)

		// Add two gadgets
		err = p.AddGadget(component.Class, utils.Point{X: 50, Y: 50}, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = p.AddGadget(component.Class, utils.Point{X: 200, Y: 200}, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		// Add an association between them
		err = p.StartAddAssociation(utils.Point{X: 50, Y: 50})
		assert.NoError(t, err)
		err = p.EndAddAssociation(component.Composition, utils.Point{X: 200, Y: 200})
		assert.NoError(t, err)

		// Verify both gadgets and the association exist
		data := p.GetDrawData()
		assert.Len(t, data.Gadgets, 2)
		assert.Len(t, data.Associations, 1)

		// Remove the first gadget - this should also remove the association
		err = p.RemoveComponent(utils.Point{X: 50, Y: 50})
		assert.NoError(t, err)

		// Verify only one gadget remains and no associations
		data = p.GetDrawData()
		assert.Len(t, data.Gadgets, 1)
		assert.Equal(t, "Class2", data.Gadgets[0].Attributes[0][0].Content)
		assert.Len(t, data.Associations, 0)
	})

	t.Run("remove association", func(t *testing.T) {
		p, err := CreateEmptyUMLProject("TestProject")
		assert.NoError(t, err)
		err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
		assert.NoError(t, err)
		err = p.SelectDiagram("TestDiagram")
		assert.NoError(t, err)

		// Add two gadgets
		err = p.AddGadget(component.Class, utils.Point{X: 50, Y: 50}, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = p.AddGadget(component.Class, utils.Point{X: 200, Y: 200}, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)

		// Add an association between them
		err = p.StartAddAssociation(utils.Point{X: 50, Y: 50})
		assert.NoError(t, err)
		err = p.EndAddAssociation(component.Composition, utils.Point{X: 200, Y: 200})
		assert.NoError(t, err)

		// Verify both gadgets and the association exist
		data := p.GetDrawData()
		assert.Len(t, data.Gadgets, 2)
		assert.Len(t, data.Associations, 1)

		// Remove the association by clicking on its midpoint
		midPoint := utils.Point{X: 125, Y: 125}
		err = p.RemoveComponent(midPoint)
		assert.NoError(t, err)

		// Verify both gadgets remain but association is removed
		data = p.GetDrawData()
		assert.Len(t, data.Gadgets, 2)
		assert.Len(t, data.Associations, 0)
	})

	t.Run("no component at point", func(t *testing.T) {
		p, err := CreateEmptyUMLProject("TestProject")
		assert.NoError(t, err)
		err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
		assert.NoError(t, err)
		err = p.SelectDiagram("TestDiagram")
		assert.NoError(t, err)

		// Add a gadget
		err = p.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)

		// Try to remove component at a point where there's nothing
		err = p.RemoveComponent(utils.Point{X: 500, Y: 500})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "No component found at the specified point")

		// Verify the original gadget is still there
		data := p.GetDrawData()
		assert.Len(t, data.Gadgets, 1)
	})

	t.Run("no diagram selected", func(t *testing.T) {
		p, err := CreateEmptyUMLProject("TestProject")
		assert.NoError(t, err)

		// Try to remove component without selecting a diagram
		err = p.RemoveComponent(utils.Point{X: 100, Y: 100})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "No current diagram selected")
	})

	t.Run("integration with undo/redo", func(t *testing.T) {
		p, err := CreateEmptyUMLProject("TestProject")
		assert.NoError(t, err)
		err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
		assert.NoError(t, err)
		err = p.SelectDiagram("TestDiagram")
		assert.NoError(t, err)

		// Add a gadget
		err = p.AddGadget(component.Class, utils.Point{X: 100, Y: 100}, 0, drawdata.DefaultGadgetColor, "TestClass")
		assert.NoError(t, err)

		// Verify the gadget exists
		data := p.GetDrawData()
		assert.Len(t, data.Gadgets, 1)

		// Remove the gadget
		err = p.RemoveComponent(utils.Point{X: 100, Y: 100})
		assert.NoError(t, err)

		// Verify the gadget was removed
		data = p.GetDrawData()
		assert.Len(t, data.Gadgets, 0)

		// Undo the removal
		err = p.UndoDiagramChange()
		assert.NoError(t, err)

		// Verify the gadget is back
		data = p.GetDrawData()
		assert.Len(t, data.Gadgets, 1)
		assert.Equal(t, "TestClass", data.Gadgets[0].Attributes[0][0].Content)

		// Redo the removal
		err = p.RedoDiagramChange()
		assert.NoError(t, err)

		// Verify the gadget was removed again
		data = p.GetDrawData()
		assert.Len(t, data.Gadgets, 0)
	})

	t.Run("remove multiple components sequentially", func(t *testing.T) {
		p, err := CreateEmptyUMLProject("TestProject")
		assert.NoError(t, err)
		err = p.CreateEmptyUMLDiagram(umldiagram.ClassDiagram, "TestDiagram")
		assert.NoError(t, err)
		err = p.SelectDiagram("TestDiagram")
		assert.NoError(t, err)

		// Add multiple gadgets
		err = p.AddGadget(component.Class, utils.Point{X: 50, Y: 50}, 0, drawdata.DefaultGadgetColor, "Class1")
		assert.NoError(t, err)
		err = p.AddGadget(component.Class, utils.Point{X: 150, Y: 150}, 0, drawdata.DefaultGadgetColor, "Class2")
		assert.NoError(t, err)
		err = p.AddGadget(component.Class, utils.Point{X: 250, Y: 250}, 0, drawdata.DefaultGadgetColor, "Class3")
		assert.NoError(t, err)

		// Verify all gadgets exist
		data := p.GetDrawData()
		assert.Len(t, data.Gadgets, 3)

		// Remove them one by one
		err = p.RemoveComponent(utils.Point{X: 50, Y: 50})
		assert.NoError(t, err)
		data = p.GetDrawData()
		assert.Len(t, data.Gadgets, 2)

		err = p.RemoveComponent(utils.Point{X: 150, Y: 150})
		assert.NoError(t, err)
		data = p.GetDrawData()
		assert.Len(t, data.Gadgets, 1)
		assert.Equal(t, "Class3", data.Gadgets[0].Attributes[0][0].Content)

		err = p.RemoveComponent(utils.Point{X: 250, Y: 250})
		assert.NoError(t, err)
		data = p.GetDrawData()
		assert.Len(t, data.Gadgets, 0)
	})
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
