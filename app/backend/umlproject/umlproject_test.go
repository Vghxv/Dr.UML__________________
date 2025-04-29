package umlproject

import (
	"testing"

	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/utils/duerror"
	"github.com/google/uuid"
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

	project.diagrams[diagram1.GetId()] = diagram1
	project.diagrams[diagram2.GetId()] = diagram2
	project.openedDiagrams[diagram1.GetId()] = diagram1

	opened, allDiagrams, uuidList := project.OpenProject()

	if len(opened) != 1 {
		t.Errorf("Expected 1 opened diagram, got %d", len(opened))
	}

	if len(opened) > 0 && opened[0].GetName() != "Diagram1" {
		t.Errorf("Expected opened diagram name Diagram1, got %s", opened[0].GetName())
	}

	if len(allDiagrams) != 2 {
		t.Errorf("Expected 2 diagrams in total, got %d", len(allDiagrams))
	}

	if len(uuidList) != 1 {
		t.Errorf("Expected 1 UUID in the list, got %d", len(uuidList))
	}

	if len(uuidList) > 0 && uuidList[0] != diagram1.GetId() {
		t.Errorf("Expected UUID to match diagram1's ID")
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

	project.diagrams[diagram1.GetId()] = diagram1
	project.diagrams[diagram2.GetId()] = diagram2

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

	project.openedDiagrams[diagram1.GetId()] = diagram1
	project.openedDiagrams[diagram2.GetId()] = diagram2

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

	project.diagrams[diagram.GetId()] = diagram

	err = project.SelectDiagram(diagram.GetId())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if project.currentDiagram != diagram {
		t.Error("Expected currentDiagram to be set to selected diagram")
	}

	invalidID := uuid.New()
	err = project.SelectDiagram(invalidID)
	if err == nil {
		t.Error("Expected error for invalid diagram ID")
	}
	if err.Error() != duerror.NewInvalidArgumentError("Diagram not found").Error() {
		t.Errorf("Expected 'Diagram not found' error, got %s", err.Error())
	}
}

func TestAddGadget(t *testing.T) {
	project := NewUMLProject("TestProject")

	diagram, err := umldiagram.NewUMLDiagram("Diagram1", umldiagram.ClassDiagram)
	if err != nil {
		t.Fatalf("Failed to create diagram: %v", err)
	}

	project.diagrams[diagram.GetId()] = diagram

	gadgetType := "class"
	previousModified := project.lastModified
	err = project.AddGadget(gadgetType, diagram.GetId())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !project.lastModified.After(previousModified) {
		t.Error("Expected lastModified to be updated")
	}

	invalidID := uuid.New()
	err = project.AddGadget(gadgetType, invalidID)
	if err == nil {
		t.Error("Expected error for invalid diagram ID")
	}
	if err.Error() != duerror.NewInvalidArgumentError("Diagram not found").Error() {
		t.Errorf("Expected 'Diagram not found' error, got %s", err.Error())
	}
}

func TestAddNewDiagram(t *testing.T) {
	project := NewUMLProject("TestProject")

	diagramType := umldiagram.ClassDiagram
	name := "NewDiagram"
	previousModified := project.lastModified

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

	// Since NewUMLDiagramWithPath returns nil, nil, createDiagram adds nothing
	if len(project.diagrams) != 0 {
		t.Errorf("Expected 0 diagrams, got %d", len(project.diagrams))
	}
	if len(project.openedDiagrams) != 0 {
		t.Errorf("Expected 0 opened diagrams, got %d", len(project.openedDiagrams))
	}
	if project.currentDiagram != nil {
		t.Error("Expected currentDiagram to be nil")
	}
}
