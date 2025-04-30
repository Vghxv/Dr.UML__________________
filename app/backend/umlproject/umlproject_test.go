package umlproject

import (
	"testing"
	"time"

	"github.com/pkg/errors"

	"Dr.uml/backend/component"
	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

func TestNewUMLProject(t *testing.T) {
	name := "TestProject"
	project := NewUMLProject(name)

	if project.GetName() != name {
		t.Errorf("Expected project name %s, got %s", name, project.GetName())
	}

	if project.lastModified.IsZero() {
		t.Error("Expected non-zero lastModified time")
	}

	if project.diagrams == nil {
		t.Error("Expected initialized diagrams map")
	}

	if project.openedDiagrams == nil {
		t.Error("Expected initialized openedDiagrams map")
	}
}

func TestOpenProject(t *testing.T) {
	project := NewUMLProject("TestProject")

	// Create diagrams with valid DiagramType
	diagram1, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram1: %v", err)
	}
	diagram2, err := umldiagram.NewUMLDiagram("Diagram2", umldiagram.SequenceDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram2: %v", err)
	}

	project.diagrams[diagram1.GetName()] = diagram1
	project.diagrams[diagram2.GetName()] = diagram2
	project.openedDiagrams[diagram1.GetName()] = diagram1

	activeDiagrams, availableDiagrams, duErr := project.OpenProject()

	// Check no error was returned
	if duErr != nil {
		t.Errorf("Expected no error, got %v", duErr)
	}

	// activeDiagrams should be empty initially because we're not setting any active diagrams
	// even though the slice is initialized with a capacity based on activeDiagrams
	if len(activeDiagrams) != 1 {
		t.Errorf("Expected 0 active diagrams, got %d", len(activeDiagrams))
	}

	// Check available diagrams (diagram names)
	if len(availableDiagrams) != 2 {
		t.Errorf("Expected 2 available diagrams, got %d", len(availableDiagrams))
	}

	// Check diagram names in availableDiagrams
	expectedNames := []string{"Diagram1", "Diagram2"}
	for _, name := range expectedNames {
		found := false
		for _, d := range availableDiagrams {
			if d == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected diagram name %s not found in availableDiagrams", name)
		}
	}

	// Verify activeDiagrams map was populated
	if len(project.activeDiagrams) != 1 {
		t.Errorf("Expected 1 active diagram in map, got %d", len(project.activeDiagrams))
	}

	// Check if the active diagram was created with the correct name
	if _, exists := project.activeDiagrams["Diagram1"]; !exists {
		t.Errorf("Expected 'Diagram1' to be in active diagrams map")
	}
}

func TestGetAvailableDiagrams(t *testing.T) {
	project := NewUMLProject("TestProject")

	diagram1, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram1: %v", err)
	}
	diagram2, err := umldiagram.NewUMLDiagram("Diagram2", umldiagram.SequenceDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram2: %v", err)
	}

	project.diagrams[diagram1.GetName()] = diagram1
	project.diagrams[diagram2.GetName()] = diagram2

	diagrams := project.GetAvailableDiagrams()

	if len(diagrams) != 2 {
		t.Errorf("Expected 2 diagrams, got %d", len(diagrams))
	}

	expected := []string{"Diagram1", "Diagram2"}
	for _, name := range expected {
		found := false
		for _, d := range diagrams {
			if d == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected diagram %s not found", name)
		}
	}
}

func TestGetLastOpenedDiagrams(t *testing.T) {
	project := NewUMLProject("TestProject")

	diagram1, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram1: %v", err)
	}
	diagram2, err := umldiagram.NewUMLDiagram("Diagram2", umldiagram.SequenceDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram2: %v", err)
	}

	project.openedDiagrams[diagram1.GetName()] = diagram1
	project.openedDiagrams[diagram2.GetName()] = diagram2

	opened := project.GetLastOpenedDiagrams()

	if len(opened) != 2 {
		t.Errorf("Expected 2 opened diagrams, got %d", len(opened))
	}

	expected := []string{"Diagram1", "Diagram2"}
	for _, name := range expected {
		found := false
		for _, d := range opened {
			if d == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected opened diagram %s not found", name)
		}
	}
}

func TestSelectDiagram(t *testing.T) {
	project := NewUMLProject("TestProject")

	diagram, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram: %v", err)
	}

	project.diagrams[diagram.GetName()] = diagram

	err = project.SelectDiagram(diagram.GetName())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if project.currentDiagram != diagram {
		t.Error("Expected currentDiagram to be set to selected diagram")
	}
}

func TestAddGadget(t *testing.T) {
	project := NewUMLProject("TestProject")

	diagram, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram: %v", err)
	}

	project.diagrams[diagram.GetName()] = diagram
	project.currentDiagram = diagram

	gadgetType := component.Class
	previousModified := project.lastModified
	time.Sleep(time.Millisecond)
	project.SelectDiagram(diagram.GetName())

	// TODO: assert GadgetType in drawdata
	err = project.AddGadget(gadgetType, utils.Point{X: 10, Y: 20})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !project.lastModified.After(previousModified) {
		t.Error("Expected lastModified to be updated")
	}
}

func TestAddNewDiagram(t *testing.T) {
	project := NewUMLProject("TestProject")

	diagramType := umldiagram.ClassDiagram
	name := "NewDiagram"
	previousModified := project.lastModified
	time.Sleep(time.Millisecond)
	err := project.AddNewDiagram(umldiagram.DiagramType(diagramType), name)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(project.diagrams) != 1 {
		t.Errorf("Expected 1 diagram, got %d", len(project.diagrams))
	}

	if len(project.openedDiagrams) != 1 {
		t.Errorf("Expected 1 opened diagram, got %d", len(project.openedDiagrams))
	}

	if project.currentDiagram == nil {
		t.Error("Expected currentDiagram to be set")
	}

	if project.currentDiagram != nil && project.currentDiagram.GetName() != name {
		t.Errorf("Expected diagram name %s, got %s", name, project.currentDiagram.GetName())
	}

	if !project.lastModified.After(previousModified) {
		t.Error("Expected lastModified to be updated")
	}

	// Test invalid diagram type
	err = project.AddNewDiagram(0, "InvalidDiagram")
	if err == nil {
		t.Error("Expected error for invalid diagram type")
	}
	if err.Error() != duerror.NewInvalidArgumentError("Invalid diagram type").Error() {
		if errors.Is(err, duerror.NewInvalidArgumentError("Invalid diagram type")) {
			t.Errorf("Expected 'Invalid diagram type' error, got %s", err.Error())
		}
		t.Errorf("Expected 'Invalid diagram type' error, got %s", err.Error())
	}
}

func TestCreateDiagram(t *testing.T) {
	project := NewUMLProject("TestProject")

	// Note: NewUMLDiagramWithPath is a stub, so we can't test actual path loading
	// Test with a dummy path, expecting the stub to return nil, nil
	path := "test/path/diagram"
	err := project.createDiagram(path)
	if err != nil {
		t.Errorf("Expected no error from stub, got %v", err)
	}

	// Check if diagram was added to project
	if len(project.diagrams) != 1 {
		t.Errorf("Expected 1 diagram, got %d", len(project.diagrams))
	}

	// Check if diagram was opened
	if len(project.openedDiagrams) != 1 {
		t.Errorf("Expected 1 opened diagram, got %d", len(project.openedDiagrams))
	}

	// Check if current diagram was set
	if project.currentDiagram == nil {
		t.Error("Expected currentDiagram to be set")
	}

	// Check if lastModified was updated
	if project.lastModified.IsZero() {
		t.Error("Expected lastModified to be updated")
	}
}

func TestInvalidPathDiagram(t *testing.T) {
	project := NewUMLProject("TestProject")

	// Test invalid paths
	invalidPaths := []string{
		"test/path/diagram<>",
		"test/path/diagram*",
		"test/path/diagram?",
		"test/path/diagram|",
		"test/path/diagram\"",
	}

	for _, path := range invalidPaths {
		err := project.createDiagram(path)
		if err == nil {
			t.Errorf("Expected error for invalid path %s, got nil", path)
		}
		if err.Error() != duerror.NewInvalidArgumentError("Invalid diagram name").Error() {
			t.Errorf("Expected 'Invalid diagram name' error for path %s, got %s", path, err.Error())
		}
	}
}

func TestSelectDiagram_DiagramNotFound(t *testing.T) {
	project := NewUMLProject("TestProject")

	err := project.SelectDiagram("NonexistentDiagram")
	if err == nil {
		t.Error("Expected error when selecting non-existent diagram")
	}
	if err.Error() != duerror.NewInvalidArgumentError("Diagram not found").Error() {
		t.Errorf("Expected 'Diagram not found' error, got %s", err.Error())
	}
}

func TestAddNewDiagram_DuplicateName(t *testing.T) {
	project := NewUMLProject("TestProject")

	// Add first diagram
	name := "NewDiagram"
	err := project.AddNewDiagram(umldiagram.ClassDiagram, name)
	if err != nil {
		t.Errorf("Expected no error on first diagram, got %v", err)
	}

	// Try to add second diagram with same name
	err = project.AddNewDiagram(umldiagram.ClassDiagram, name)
	if err == nil {
		t.Error("Expected error when adding diagram with duplicate name")
	}
	if err.Error() != duerror.NewInvalidArgumentError("Diagram name already exists").Error() {
		t.Errorf("Expected 'Diagram name already exists' error, got %s", err.Error())
	}
}

func TestCreateDiagram_EmptyPath(t *testing.T) {
	project := NewUMLProject("TestProject")

	err := project.createDiagram("")
	if err == nil {
		t.Error("Expected error when creating diagram with empty path")
	}
	if err.Error() != duerror.NewInvalidArgumentError("Invalid diagram name").Error() {
		t.Errorf("Expected 'Invalid diagram name' error, got %s", err.Error())
	}
}

func TestOpenProject_NoOpenedDiagrams(t *testing.T) {
	project := NewUMLProject("TestProject")

	diagram, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram: %v", err)
	}

	project.diagrams[diagram.GetName()] = diagram
	// Don't add to openedDiagrams

	activeDiagrams, availableDiagrams, err := project.OpenProject()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(activeDiagrams) != 0 {
		t.Errorf("Expected 0 active diagrams, got %d", len(activeDiagrams))
	}
	if len(availableDiagrams) != 1 {
		t.Errorf("Expected 1 available diagram, got %d", len(availableDiagrams))
	}
	if len(project.activeDiagrams) != 0 {
		t.Errorf("Expected 0 active diagrams in map, got %d", len(project.activeDiagrams))
	}
}
