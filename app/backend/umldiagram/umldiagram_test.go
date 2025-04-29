package umldiagram

import (
	"testing"
	"time"

	"Dr.uml/backend/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUMLDiagram(t *testing.T) {
	name := "TestDiagram"
	diagramType := NewDiagramType("ClassDiagram")

	diagram := NewUMLDiagram(name, diagramType)

	assert.NotNil(t, diagram, "Diagram should not be nil")
	assert.Equal(t, name, diagram.GetName(), "Diagram name should match input")
	assert.Equal(t, diagramType, diagram.diagramType, "Diagram type should match input")
	assert.NotEqual(t, uuid.Nil, diagram.GetId(), "Diagram ID should be a valid UUID")
	assert.WithinDuration(t, time.Now(), diagram.lastModified, time.Second, "Last modified should be recent")
	assert.Equal(t, utils.Point{X: 0, Y: 0}, diagram.startPoint, "Start point should be initialized to zero")
}

func TestGetId(t *testing.T) {
	diagram := NewUMLDiagram("TestDiagram", NewDiagramType("ClassDiagram"))
	id := diagram.GetId()

	assert.NotEqual(t, uuid.Nil, id, "GetId should return a valid UUID")
	assert.Equal(t, diagram.id, id, "GetId should return the diagram's ID")
}

func TestGetName(t *testing.T) {
	name := "TestDiagram"
	diagram := NewUMLDiagram(name, NewDiagramType("ClassDiagram"))

	assert.Equal(t, name, diagram.GetName(), "GetName should return the diagram's name")
}

func TestNewUMLDiagramWithPath(t *testing.T) {
	// Since NewUMLDiagramWithPath is a TODO, we'll test for nil return as per current implementation
	path := "test/path"
	diagram, err := NewUMLDiagramWithPath(path)

	assert.Nil(t, diagram, "Diagram should be nil as per current implementation")
	assert.Nil(t, err, "Error should be nil as per current implementation")
}

func TestAddGadget(t *testing.T) {
	diagram := NewUMLDiagram("TestDiagram", NewDiagramType("ClassDiagram"))
	err := diagram.AddGadget("TestGadget")

	// Since AddGadget is a TODO, we'll test for nil error as per current implementation
	assert.Nil(t, err, "Error should be nil as per current implementation")
}
