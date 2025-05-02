package umlproject

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Dr.uml/backend/component"
	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

func TestNewUMLProject(t *testing.T) {
	// Given
	name := "TestProject"

	// When
	project := NewUMLProject(name)

	// Then
	assert.Equal(t, name, project.GetName(), "Project name should match the provided name")
	assert.False(t, project.lastModified.IsZero(), "LastModified time should not be zero")
	assert.NotNil(t, project.diagrams, "Diagrams map should be initialized")
	assert.NotNil(t, project.openedDiagrams, "OpenedDiagrams map should be initialized")
}

func TestOpenProject(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")

	// Create diagrams with valid DiagramType
	diagram1, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	require.NoError(t, err, "Failed to create diagram1")

	diagram2, err := umldiagram.NewUMLDiagram("Diagram2", umldiagram.SequenceDiagram)
	require.NoError(t, err, "Failed to create diagram2")

	project.diagrams[diagram1.GetName()] = diagram1
	project.diagrams[diagram2.GetName()] = diagram2
	project.openedDiagrams[diagram1.GetName()] = diagram1

	// When
	activeDiagrams, availableDiagrams, duErr := project.OpenProject()

	// Then
	assert.NoError(t, duErr, "OpenProject should not return an error")
	assert.NotNil(t, activeDiagrams, "ActiveDiagrams should not be nil")
	assert.Len(t, activeDiagrams, 1, "There should be 1 active diagram")
	assert.Len(t, availableDiagrams, 2, "There should be 2 available diagrams")

	// Check diagram names in availableDiagrams
	expectedNames := []string{"Diagram1", "Diagram2"}
	for _, name := range expectedNames {
		assert.Contains(t, availableDiagrams, name, "AvailableDiagrams should contain %s", name)
	}

	// Verify activeDiagrams map was not populated
	assert.Len(t, project.activeDiagrams, 0, "Project's activeDiagrams map should be empty")
}

func TestGetAvailableDiagrams(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")

	diagram1, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	require.NoError(t, err, "Failed to create diagram1")

	diagram2, err := umldiagram.NewUMLDiagram("Diagram2", umldiagram.SequenceDiagram)
	require.NoError(t, err, "Failed to create diagram2")

	project.diagrams[diagram1.GetName()] = diagram1
	project.diagrams[diagram2.GetName()] = diagram2

	// When
	diagrams := project.GetAvailableDiagrams()

	// Then
	assert.Len(t, diagrams, 2, "Should return 2 available diagrams")

	expected := []string{"Diagram1", "Diagram2"}
	for _, name := range expected {
		assert.Contains(t, diagrams, name, "Available diagrams should contain %s", name)
	}
}

func TestGetLastOpenedDiagrams(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")

	diagram1, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	require.NoError(t, err, "Failed to create diagram1")

	diagram2, err := umldiagram.NewUMLDiagram("Diagram2", umldiagram.SequenceDiagram)
	require.NoError(t, err, "Failed to create diagram2")

	project.openedDiagrams[diagram1.GetName()] = diagram1
	project.openedDiagrams[diagram2.GetName()] = diagram2

	// When
	opened := project.GetLastOpenedDiagrams()

	// Then
	assert.Len(t, opened, 2, "Should return 2 last opened diagrams")

	expected := []string{"Diagram1", "Diagram2"}
	for _, name := range expected {
		assert.Contains(t, opened, name, "Last opened diagrams should contain %s", name)
	}
}

func TestSelectDiagram(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")

	diagram, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	require.NoError(t, err, "Failed to create diagram")

	project.diagrams[diagram.GetName()] = diagram

	// When
	err = project.SelectDiagram(diagram.GetName())

	// Then
	assert.NoError(t, err, "SelectDiagram should not return an error")
	assert.Equal(t, diagram, project.currentDiagram, "CurrentDiagram should be set to the selected diagram")
}

func TestAddGadget(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")

	diagram, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	require.NoError(t, err, "Failed to create diagram")

	project.diagrams[diagram.GetName()] = diagram
	project.currentDiagram = diagram

	gadgetType := component.Class
	previousModified := project.lastModified
	time.Sleep(time.Millisecond) // Ensure time difference
	err = project.SelectDiagram(diagram.GetName())
	assert.NoError(t, err, "SelectDiagram should not return an error")
	// When
	err = project.AddGadget(gadgetType, utils.Point{X: 10, Y: 20})

	assert.NoError(t, err, "AddGadget should not return an error")
	assert.True(t, project.lastModified.After(previousModified), "LastModified should be updated")
}

func TestAddNewDiagram(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")

	diagramType := umldiagram.ClassDiagram
	name := "NewDiagram"
	previousModified := project.lastModified
	time.Sleep(time.Millisecond) // Ensure time difference

	// When
	err := project.AddNewDiagram(umldiagram.DiagramType(diagramType), name)

	// Then
	assert.NoError(t, err, "AddNewDiagram should not return an error")
	assert.Len(t, project.diagrams, 1, "Should have 1 diagram")
	assert.Len(t, project.openedDiagrams, 1, "Should have 1 opened diagram")
	assert.NotNil(t, project.currentDiagram, "CurrentDiagram should be set")

	if project.currentDiagram != nil {
		assert.Equal(t, name, project.currentDiagram.GetName(), "CurrentDiagram should have the provided name")
	}

	assert.True(t, project.lastModified.After(previousModified), "LastModified should be updated")

	// Test invalid diagram type
	err = project.AddNewDiagram(0, "InvalidDiagram")
	assert.Error(t, err, "Should return error for invalid diagram type")
	assert.Equal(t, duerror.NewInvalidArgumentError("Invalid diagram type").Error(), err.Error(),
		"Should return 'Invalid diagram type' error")
}

func TestCreateDiagram(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")
	path := "test/path/diagram"

	// When
	err := project.createDiagram(path)

	// Then
	assert.NoError(t, err, "createDiagram should not return an error from stub")
	assert.Len(t, project.diagrams, 1, "Should have 1 diagram")
	assert.Len(t, project.openedDiagrams, 1, "Should have 1 opened diagram")
	assert.NotNil(t, project.currentDiagram, "CurrentDiagram should be set")
	assert.False(t, project.lastModified.IsZero(), "LastModified should be updated")
}

func TestInvalidPathDiagram(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")
	invalidPaths := []string{
		"test/path/diagram<>",
		"test/path/diagram*",
		"test/path/diagram?",
		"test/path/diagram|",
		"test/path/diagram\"",
	}

	// Then
	for _, path := range invalidPaths {
		err := project.createDiagram(path)
		assert.Error(t, err, "Should return error for invalid path: %s", path)
		assert.Equal(t, duerror.NewInvalidArgumentError("Invalid diagram name").Error(), err.Error(),
			"Should return 'Invalid diagram name' error for path: %s", path)
	}
}

func TestSelectDiagram_DiagramNotFound(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")

	// When
	err := project.SelectDiagram("NonexistentDiagram")

	// Then
	assert.Error(t, err, "Should return error when selecting non-existent diagram")
	assert.Equal(t, duerror.NewInvalidArgumentError("Diagram not found").Error(), err.Error(),
		"Should return 'Diagram not found' error")
}

func TestAddNewDiagram_DuplicateName(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")
	name := "NewDiagram"

	// Add first diagram
	err := project.AddNewDiagram(umldiagram.ClassDiagram, name)
	require.NoError(t, err, "Should not return error on first diagram")

	// When - Try to add second diagram with same name
	err = project.AddNewDiagram(umldiagram.ClassDiagram, name)

	// Then
	assert.Error(t, err, "Should return error when adding diagram with duplicate name")
	assert.Equal(t, duerror.NewInvalidArgumentError("Diagram name already exists").Error(), err.Error(),
		"Should return 'Diagram name already exists' error")
}

func TestCreateDiagram_EmptyPath(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")

	// When
	err := project.createDiagram("")

	// Then
	assert.Error(t, err, "Should return error when creating diagram with empty path")
	assert.Equal(t, duerror.NewInvalidArgumentError("Invalid diagram name").Error(), err.Error(),
		"Should return 'Invalid diagram name' error")
}

func TestOpenProject_NoOpenedDiagrams(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")

	diagram, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	require.NoError(t, err, "Failed to create diagram")

	project.diagrams[diagram.GetName()] = diagram
	// Don't add to openedDiagrams

	// When
	activeDiagrams, availableDiagrams, err := project.OpenProject()

	// Then
	assert.NoError(t, err, "OpenProject should not return an error")
	assert.Len(t, activeDiagrams, 0, "Should have 0 active diagrams")
	assert.Len(t, availableDiagrams, 1, "Should have 1 available diagram")
	assert.Len(t, project.activeDiagrams, 0, "Project's activeDiagrams map should be empty")
}

func TestInvalidateCanvas_NoDiagram(t *testing.T) {
	// Given
	project := NewUMLProject("TestProject")
	// Don't select a diagram, so currentDiagram will be nil

	// When
	err := project.InvalidateCanvas()

	// Then
	assert.Error(t, err, "Should return error when no diagram is selected")
	assert.Equal(t, duerror.NewInvalidArgumentError("No current diagram selected").Error(), err.Error(),
		"Should return 'No current diagram selected' error")
}
