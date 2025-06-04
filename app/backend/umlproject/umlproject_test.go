package umlproject

import (
	"Dr.uml/backend/drawdata"
	"os"
	"testing"
	"time"

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

	// success, select a active diagram
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)
	name := p.GetCurrentDiagramName()
	assert.Equal(t, "TestDiagram", name)

	// success, select a non-active diagram
	err = p.CloseDiagram("TestDiagram")
	assert.NoError(t, err)
	err = p.SelectDiagram("TestDiagram")
	assert.NoError(t, err)
	name = p.GetCurrentDiagramName()
	assert.Equal(t, "TestDiagram", name)

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

func TestLoadExistUMLProject(t *testing.T) {
	// Test the function (currently returns nil, nil)
	p, err := LoadExistUMLProject("TestProject")
	assert.Nil(t, p)
	assert.Nil(t, err)
}

func TestOpenDiagram(t *testing.T) {
	proj, err := CreateEmptyUMLProject("fuck")
	assert.NoError(t, err)
	root, ok := os.LookupEnv("APP_ROOT")
	assert.True(t, ok)
	err = proj.OpenDiagram(root + "/backend/example.json5")
	assert.NoError(t, err)
}
