package umlproject

import (
	"testing"
	"time"

	"Dr.uml/backend/umldiagram"
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
	diagram1 := umldiagram.NewUMLDiagram("Diagram1", umldiagram.NewDiagramType("ClassDiagram"))
	diagram2 := umldiagram.NewUMLDiagram("Diagram2", umldiagram.NewDiagramType("SequenceDiagram"))

	project.diagrams[diagram1.GetId()] = diagram1
	project.diagrams[diagram2.GetId()] = diagram2
	project.openedDiagrams[diagram1.GetId()] = diagram1

	opened, allDiagrams := project.OpenProject()

	if len(opened) != 1 {
		t.Errorf("Expected 1 opened diagram, got %d", len(opened))
	}

	if opened[0].GetName() != "Diagram1" {
		t.Errorf("Expected opened diagram name Diagram1, got %s", opened[0].GetName())
	}

	if len(allDiagrams) != 2 {
		t.Errorf("Expected 2 diagrams in total, got %d", len(allDiagrams))
	}
}

func TestGetAvailableDiagrams(t *testing.T) {
	project := NewUMLProject("TestProject")
	diagram1 := umldiagram.NewUMLDiagram("Diagram1", "class")
	diagram2 := umldiagram.NewUMLDiagram("Diagram2", "sequence")

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
	diagram1 := umldiagram.NewUMLDiagram("Diagram1", "class")
	diagram2 := umldiagram.NewUMLDiagram("Diagram2", "sequence")

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
	diagram := umldiagram.NewUMLDiagram("Diagram1", "class")
	project.diagrams[diagram.GetId()] = diagram

	err := project.SelectDiagram(diagram.GetId())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if project.currentDiagram != diagram {
		t.Error("Expected currentDiagram to be set to selected diagram")
	}

	err = project.SelectDiagram(uuid.New())
	if err == nil {
		t.Error("Expected error for invalid diagram ID")
	}
	if err.Error() != "Diagram not found" {
		t.Errorf("Expected 'Diagram not found' error, got %s", err.Error())
	}
}

func TestAddGadget(t *testing.T) {
	project := NewUMLProject("TestProject")
	diagram := umldiagram.NewUMLDiagram("Diagram1", "class")
	project.diagrams[diagram.GetId()] = diagram

	gadgetType := "class"
	err := project.AddGadget(gadgetType, diagram.GetId())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if project.lastModified.Before(time.Now().Add(-time.Second)) {
		t.Error("Expected lastModified to be updated")
	}

	err = project.AddGadget(gadgetType, uuid.New())
	if err == nil {
		t.Error("Expected error for invalid diagram ID")
	}
	if err.Error() != "Diagram not found" {
		t.Errorf("Expected 'Diagram not found' error, got %s", err.Error())
	}
}

func TestAddNewDiagram(t *testing.T) {
	project := NewUMLProject("TestProject")
	diagramType := umldiagram.NewDiagramType("ClassDiagram")
	name := "NewDiagram"

	err := project.AddNewDiagram(diagramType, name)
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

	if project.currentDiagram.GetName() != name {
		t.Errorf("Expected diagram name %s, got %s", name, project.currentDiagram.GetName())
	}

	if project.lastModified.Before(time.Now().Add(-time.Second)) {
		t.Error("Expected lastModified to be updated")
	}
}
